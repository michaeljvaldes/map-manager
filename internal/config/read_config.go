package config

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

func BuildConfigFromFile(configFilePath string) Config {
	// Set the file name of the configurations file
	viper.SetConfigName(filepath.Base(configFilePath))

	// Set the path to look for the configurations file
	viper.AddConfigPath(filepath.Dir(configFilePath))

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var ymlConfig YmlConfig

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("period", 24)

	err := viper.Unmarshal(&ymlConfig)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	sanitizedConfig := SanitizeConfig(ymlConfig)
	return sanitizedConfig
}
