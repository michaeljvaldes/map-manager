package mapcycle

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

func ExecuteMapCycleOnSchedule(configuration config.Config) {
	log.Printf("Starting first map cycle at " + configuration.StartTime.String())
	gocron.Every(uint64(configuration.Period.Minutes())).Minutes().From(&configuration.StartTime).Do(executeMapCycle, configuration.UnminedPath, configuration.WorldPath, configuration.SiteId, configuration.DeployToken, configuration.Period)
	<-gocron.Start()
}

func executeMapCycle(unminedPath, worldPath, siteId, deployToken string, period time.Duration) {
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
