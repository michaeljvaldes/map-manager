package mapdeploy

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

const apiUrlPattern string = "https://api.netlify.com/api/v1/sites/%s/deploys"

func DeployMapSite(siteDir string, siteId string, deployToken string) {
	zipFileName := "site.zip"
	zipFilePath := filepath.Join(filepath.Dir(siteDir), zipFileName)
	ZipDirectory(siteDir, zipFilePath)
	// zipFilePath = "C:/dev/go/minecraft-mapper/test/site.zip"
	deployZipToSite(zipFilePath, siteId, deployToken)
}

func deployZipToSite(zipFilePath string, siteId string, deployToken string) {
	zipFile, err := os.Open(zipFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer zipFile.Close()

	url := fmt.Sprintf(apiUrlPattern, siteId)

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

/*
curl -X POST https://api.netlify.com/api/v1/sites/23498d3f-a255-4471-980f-fe15896ef693/deploys -H "Authorization: Bearer -BWqwl7FipqgTcJmzKl-GbDqwNIcFXAR853qg1itMVw" -H "Content-Type: application/zip" --data-binary @site.zip

fc.exe /b "C:\dev\go\minecraft-mapper\test" "C:\Users\micha\AppData\Local\Temp\map_temp_dir2936816216\site.zip"
*/
