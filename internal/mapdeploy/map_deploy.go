package mapdeploy

import (
	"log"
	"mapmanager/internal/netlify"
	"mapmanager/internal/zip"
	"path/filepath"
)

func DeployMapSite(siteDir string, siteId string, deployToken string) {
	log.Println("Deploying site to site id: " + siteId)
	zipFileName := "site.zip"
	zipFilePath := filepath.Join(filepath.Dir(siteDir), zipFileName)
	zip.ZipDirectory(siteDir, zipFilePath)
	err := netlify.PostDeploy(siteId, deployToken, zipFilePath)
	if err != nil {
		handleError(err, "error posting deploy")
	}
	log.Println("Finished deploying site to site id: " + siteId)
}

func handleError(err error, context string) {
	deployErr := DeployError{Err: err, Context: context}
	log.Fatal(deployErr.ErrorMessage())
}
