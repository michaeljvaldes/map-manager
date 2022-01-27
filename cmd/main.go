package main

import (
	"log"
	"mapmanager/internal/config"
	"mapmanager/internal/mapcycle"
	"os"
	"path/filepath"
)

/* TODO
create structs for inputs
accept terminal inputs
accept config inputs
validate inputs
- reorganize code into packages
- combine all functions into one main.go (this one)
- create 4 maps
- combine 4 maps into one html site
- create temp files, destroy
- deal with paths responsibly
deal with path to unmined-cli in mapgen
add args to choose between gen, deploy, both, and both at a frequency
add helpful comments and handle errors
*/

func main() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal("Cannot get path of current directory", err)
	}
	configFilePath := filepath.Join(filepath.Dir(exePath), config.DefaultConfigName)
	configuration := config.BuildConfigFromFile(configFilePath)
	mapcycle.ExecuteMapCycleOnSchedule(configuration)
}
