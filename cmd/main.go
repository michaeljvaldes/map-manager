package main

import (
	"fmt"
	"io/ioutil"
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

	unminedPath := "C:/dev/go/minecraft-mapper/third_party/unmined/unmined-cli.exe"
	worldPath := filepath.Clean(filepath.FromSlash("C:/dev/go/minecraft-mapper/World_of_Duane/"))
	tempDir := createTempDir()
	defer os.RemoveAll(tempDir)
	siteDir := filepath.Join(tempDir, "site")
	siteId := "23498d3f-a255-4471-980f-fe15896ef693"
	deployToken := "-BWqwl7FipqgTcJmzKl-GbDqwNIcFXAR853qg1itMVw"
	time := 1

	gocron.Every(uint64(time)).Minutes().Do(genPrepAndDeploy, unminedPath, worldPath, siteDir, siteId, deployToken)
	<-gocron.Start()
	// args := arguments{worldPath: worldPath, siteId: siteId, deployToken: deployToken}
	// valid, errs := args.Valid()
	// if !valid {
	// 	for _, err := range errs {
	// 		fmt.Errorf(err.Error())
	// 	}
	// } else {

	// }

}

func genPrepAndDeploy(unminedPath, worldPath, siteDir, siteId, deployToken string) {
	fmt.Println(time.Now())
	mapgen.GenerateMaps(unminedPath, worldPath, siteDir)
	mapprep.PrepareMaps(siteDir)
	mapdeploy.DeployMapSite(siteDir, siteId, deployToken)
}

func createTempDir() string {
	tempDir, err := ioutil.TempDir("", "map_temp_dir")
	if err != nil {
		fmt.Printf("Could not create temp dir")
	}
	tempDir = filepath.Clean(tempDir)
	return tempDir
}
