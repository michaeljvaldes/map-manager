package mapdeploy

import (
	"fmt"
	"log"
	"mapmanager/internal/zip"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

const apiUrlPattern string = "https://api.netlify.com/api/v1/sites/%s/deploys"

func DeployMapSite(siteDir string, siteId string, deployToken string) {
	log.Println("Deploying site to site id: " + siteId)
	zipFileName := "site.zip"
	zipFilePath := filepath.Join(filepath.Dir(siteDir), zipFileName)
	zip.ZipDirectory(siteDir, zipFilePath)
	deployZipToSite(zipFilePath, siteId, deployToken)
	log.Println("Finished deploying site to site id: " + siteId)
}

func deployZipToSite(zipFilePath string, siteId string, deployToken string) {
	zipFile, err := os.Open(zipFilePath)
	if err != nil {
		handleError(err, "Error opening site zip")
	}
	defer zipFile.Close()

	url := fmt.Sprintf(apiUrlPattern, siteId)

	request, err := http.NewRequest(http.MethodPost, url, zipFile)
	if err != nil {
		handleError(err, "Error creating deploy request")
	}

	request.Header.Add("Authorization", "Bearer "+deployToken)
	request.Header.Add("Content-Type", "application/zip")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		handleError(err, "Error executing deployment post request")
	}
	if response.StatusCode != http.StatusOK {
		handleError(nil, "Error executing deployment post request; Response Code = "+strconv.Itoa(response.StatusCode))
	}
}

func handleError(err error, context string) {
	deployErr := DeployError{Err: err, Context: context}
	log.Fatal(deployErr.ErrorMessage())
}
