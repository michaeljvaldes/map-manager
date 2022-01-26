package main

import (
	"minecraftmapper/internal/config"
	"minecraftmapper/internal/mapcycle"
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

// curl -H "Authorization: Bearer -BWqwl7FipqgTcJmzKl-GbDqwNIcFXAR853qg1itMVw" https://api.netlify.com/api/v1/sites/23498d3f-a255-4471-980f-fe15896ef693/files
func main() {
	configuration := config.BuildConfigFromFile("C:/dev/go/minecraft-mapper/test/config/sample_config.yml")
	mapcycle.ExecuteMapCycleOnSchedule(configuration)
}
