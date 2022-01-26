package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"minecraftmapper/internal/config"
	"minecraftmapper/internal/mapdeploy"
	"minecraftmapper/internal/mapgen"
	"minecraftmapper/internal/mapprep"
	"os"
	"path/filepath"
	"time"

	"github.com/jasonlvhit/gocron"
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

	var configuration config.Config
	useTestData := false
	if useTestData {
		configuration = getTestData()
	} else {
		configuration = config.BuildConfigFromFile("C:/dev/go/minecraft-mapper/test/config/sample_config.yml")
	}

	genPrepAndDeployOnSchedule(configuration)
}

func genPrepAndDeployOnSchedule(configuration config.Config) {
	log.Printf("Starting first map cycle at " + configuration.StartTime.String())
	gocron.Every(uint64(configuration.Period.Minutes())).Minutes().From(&configuration.StartTime).Do(genPrepAndDeploy, configuration.UnminedPath, configuration.WorldPath, configuration.SiteId, configuration.DeployToken, configuration.Period)
	<-gocron.Start()
}

func genPrepAndDeploy(unminedPath, worldPath, siteId, deployToken string, period time.Duration) {
	startTime := time.Now()
	log.Println("Restarting map cycle: current time " + startTime.String())

	tempDir := createTempDir()
	defer os.RemoveAll(tempDir)
	siteDir := filepath.Join(tempDir, "site")

	mapgen.GenerateMaps(unminedPath, worldPath, siteDir)
	mapprep.PrepareMaps(siteDir)
	mapdeploy.DeployMapSite(siteDir, siteId, deployToken)

	log.Println("Map cycle complete: current time " + time.Now().String())
	log.Println("Beginning next map cycle at approximately " + startTime.Add(period).String())
}

func createTempDir() string {
	tempDir, err := ioutil.TempDir("", "map_temp_dir")
	if err != nil {
		fmt.Printf("Could not create temp dir")
	}
	tempDir = filepath.Clean(tempDir)
	return tempDir
}

func getTestData() config.Config {
	return config.Config{
		UnminedPath: "C:/dev/go/minecraft-mapper/third_party/unmined/unmined-cli.exe",
		WorldPath:   filepath.Clean(filepath.FromSlash("C:/dev/go/minecraft-mapper/test/sample_map/World_of_Duane/")),
		SiteId:      "23498d3f-a255-4471-980f-fe15896ef693",
		DeployToken: "-BWqwl7FipqgTcJmzKl-GbDqwNIcFXAR853qg1itMVw",
		Period:      time.Minute,
		StartTime:   time.Now().Add(time.Minute),
	}
}
