package main

import (
	"fmt"
	"net/http"
	"os"
)

func DeployMapSite(zipFile string, siteId string, deployToken string) {
	deployZipToSite(zipFile, siteId, deployToken)
}

func deployZipToSite(zipFilePath string, siteId string, deployToken string) {
	// curl -X POST -H "Authorization: Bearer -BWqwl7FipqgTcJmzKl-GbDqwNIcFXAR853qg1itMVw" -H "Content-Type: application/zip" -d '{"files": "assets\test2.zip"}' https://api.netlify.com/api/v1/sites/23498d3f-a255-4471-980f-fe15896ef693

	zipFile, err := os.Open(zipFilePath)
	if err != nil {
		fmt.Println(err)
	}

	url := fmt.Sprintf("https://api.netlify.com/api/v1/sites/%s/deploys", siteId)

	request, err := http.NewRequest(http.MethodPost, url, zipFile)
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Add("Authorization", "Bearer "+deployToken)
	request.Header.Add("Content-Type", "application/zip")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Response:", response)

}
