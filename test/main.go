package main

import (
	"minecraftmapper/internal/config"
	"minecraftmapper/internal/mapcycle"
	"time"
)

func main() {
	configuration := getTestConfig()
	mapcycle.ExecuteMapCycleOnSchedule(configuration)

}

func getTestConfig() config.Config {
	testConfigPath := "config/sample_config.yml"
	configuration := config.BuildConfigFromFile(testConfigPath)
	configuration.Period = time.Minute
	configuration.StartTime = time.Now().Add(time.Minute)
	return configuration
}
