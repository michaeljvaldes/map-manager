package config

import (
	"time"
)

const DefaultConfigName = "config.yml"

type Config struct {
	UnminedPath string
	WorldPath   string
	SiteId      string
	DeployToken string
	Period      time.Duration
	StartTime   time.Time
}

type YmlConfig struct {
	UnminedPath string
	WorldPath   string
	SiteId      string
	DeployToken string
	Period      string
	StartTime   string
}
