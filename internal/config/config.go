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
	Period      int
	StartTime   time.Time
}

type YmlConfig struct {
	UnminedPath string
	WorldPath   string
	SiteId      string
	DeployToken string
	Period      int
	StartTime   string
}
