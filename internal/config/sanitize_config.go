package config

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func SanitizeConfig(ymlConfig YmlConfig) Config {
	up := sanitizeUnminedPath(ymlConfig.UnminedPath)
	wp := sanitizeWorldPath(ymlConfig.WorldPath)
	si, dt := sanitizeSiteIdAndDeployToken(ymlConfig.SiteId, ymlConfig.DeployToken)
	p := sanitizePeriod(ymlConfig.Period)
	st := sanitizeStartTime(ymlConfig.StartTime)
	return Config{
		UnminedPath: up,
		WorldPath:   wp,
		SiteId:      si,
		DeployToken: dt,
		Period:      p,
		StartTime:   st,
	}
}

func sanitizeUnminedPath(filePath string) string {
	errMsg := "not a valid path for unmined-cli.exe"
	sanitized := filepath.Clean(filepath.FromSlash(filePath))
	if info, err := os.Stat(sanitized); err != nil {
		handleError(err, errMsg)
	} else if !strings.Contains(info.Name(), "exe") || info.IsDir() {
		handleError(err, errMsg)
	}
	return sanitized
}

func sanitizeWorldPath(dirPath string) string {
	errMsg := "not a valid path for minecraft world directory"
	sanitized := filepath.Clean(filepath.FromSlash(dirPath))
	if info, err := os.Stat(sanitized); err != nil || !info.IsDir() {
		handleError(err, errMsg)
	}
	return sanitized
}

func sanitizeSiteIdAndDeployToken(siteId string, deployToken string) (string, string) {
	sanitizedSiteId := strings.Trim(siteId, " ")
	sanitizedDeployToken := strings.Trim(deployToken, " ")
	// test with get request
	return sanitizedSiteId, sanitizedDeployToken
}

func sanitizePeriod(period string) time.Duration {
	errMsg := "invalid period for map generation/deployment"
	sanitized, err := parseHoursMinutes(period)
	if err != nil {
		handleError(err, errMsg)
	}
	return sanitized
}

func sanitizeStartTime(startTime string) time.Time {
	if startTime == "now" {
		return time.Now().Add(time.Minute)
	}

	errMsg := "invalid start time for map generation/deployment"
	dur, err := parseHoursMinutes(startTime)
	if err != nil {
		handleError(err, errMsg)
	}
	now := time.Now()
	startToday := time.Date(now.Year(), now.Month(), now.Day(), 0, int(dur.Minutes()), 0, 0, now.Location())
	tomorrow := now.Add(time.Hour * 24)
	startTomorrow := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, int(dur.Minutes()), 0, 0, tomorrow.Location())
	if startToday.After(now) {
		return startToday
	} else {
		return startTomorrow
	}
}

func parseHoursMinutes(input string) (time.Duration, error) {
	split := strings.Split(input, ":")
	if len(split) == 2 {
		hours, err := strconv.Atoi(split[0])
		if err == nil {
			minutes, err := strconv.Atoi(split[1])
			if err == nil {
				return time.Duration(time.Hour*time.Duration(hours) + time.Minute*time.Duration(minutes)), nil
			}
		}
	}
	return time.Duration(0), errors.New("invalid time for format (HH:mm)")
}

func handleError(err error, context string) {
	prepErr := ConfigError{Err: err, Context: context}
	log.Fatal(prepErr.ErrorMessage())
}
