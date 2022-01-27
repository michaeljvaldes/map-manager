package main

import (
	"log"
	"mapmanager/internal/config"
	"mapmanager/internal/mapcycle"
	"os"
	"path/filepath"
)

func main() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal("Cannot get path of current directory", err)
	}
	configFilePath := filepath.Join(filepath.Dir(exePath), config.DefaultConfigName)
	configuration := config.BuildConfigFromFile(configFilePath)
	mapcycle.ExecuteMapCycleOnSchedule(configuration)
}
