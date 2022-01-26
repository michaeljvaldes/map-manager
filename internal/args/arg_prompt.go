package args

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
)

func BuildArgsFromPrompt() Arguments {
	return Arguments{
		UnminedPath: getUnminedPath(),
		WorldPath:   getWorldPath(),
		SiteId:      getSiteId(),
		DeployToken: getDeployToken(),
		Period:      getPeriod(),
		StartTime:   getStartTime(),
	}
}

func getUnminedPath() string {
	errMsg := "not a valid path for unmined-cli.exe"
	validate := func(input string) error {
		if info, err := os.Stat(input); err != nil {
			return errors.New(errMsg)
		} else if !strings.Contains(info.Name(), "exe") || info.IsDir() {
			return errors.New(errMsg)
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Enter the full path of your installed unmined-cli.exe",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		handleError(err, errMsg)
	}
	return result
}

func getWorldPath() string {
	errMsg := "not a valid path for minecraft world directory"
	validate := func(input string) error {
		if info, err := os.Stat(input); err != nil || !info.IsDir() {
			return errors.New(errMsg)
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Enter the full path of your minecraft world directory",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		handleError(err, errMsg)
	}
	return result
}

func getSiteId() string {
	errMsg := "site ID must be more than 0 characters"
	validate := func(input string) error {
		if len(input) == 0 {
			return errors.New(errMsg)
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Enter the site ID of your Netlify site",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		handleError(err, errMsg)
	}
	return result
}

func getDeployToken() string {
	errMsg := "invalid site ID or deploy token"
	validate := func(input string) error {
		if len(input) == 0 {
			return errors.New(errMsg)
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Enter deploy token of your Netlify user",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		handleError(err, errMsg)
	}
	return result
}

func getPeriod() time.Duration {
	errMsg := "invalid frequency for map generation"

	validate := func(input string) error {
		_, err := parseHoursMinutes(input)
		if err != nil {
			return errors.New(errMsg)
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:     "Enter the frequency with which you want to deploy (HH:mm)",
		Validate:  validate,
		Default:   "24:00",
		AllowEdit: true,
	}

	result, err := prompt.Run()

	if err != nil {
		handleError(err, errMsg)
	}
	duration, _ := parseHoursMinutes(result)
	return duration
}

func getStartTime() time.Time {
	errMsg := "invalid start time for map generation"

	validate := func(input string) error {
		_, err := parseHoursMinutes(input)
		if err != nil {
			return errors.New(errMsg)
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:     "Enter the time you want to start (HH:mm)",
		Validate:  validate,
		Default:   "24:00",
		AllowEdit: true,
	}

	result, err := prompt.Run()

	if err != nil {
		handleError(err, errMsg)
	}
	dur, _ := parseHoursMinutes(result)
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
	prepErr := ArgError{Err: err, Context: context}
	log.Fatal(prepErr.ErrorMessage())
}
