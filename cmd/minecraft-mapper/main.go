package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	baseFolder := "../../assets/test1/"
	zipFile := "../../assets/site.zip"
	zipFolder(baseFolder, zipFile)
	deployZipToSite(zipFile)
}

func zipFolder(baseFolder string, zipFile string) {

	// Get a Buffer to Write To
	outFile, err := os.Create(zipFile)
	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()

	// Create a new zip archive.
	w := zip.NewWriter(outFile)

	// Add some files to the archive.
	addFiles(w, baseFolder, "")

	if err != nil {
		fmt.Println(err)
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		fmt.Println(err)
	}

}

func addFiles(w *zip.Writer, basePath, baseInZip string) {
	// Open the Directory
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		fmt.Println(basePath + file.Name())
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(basePath + file.Name())
			if err != nil {
				fmt.Println(err)
			}

			// Add some files to the archive.
			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				fmt.Println(err)
			}
			_, err = f.Write(dat)
			if err != nil {
				fmt.Println(err)
			}
		} else if file.IsDir() {

			// Recurse
			newBase := basePath + file.Name() + "/"
			fmt.Println("Recursing and Adding SubDir: " + file.Name())
			fmt.Println("Recursing and Adding SubDir: " + newBase)

			addFiles(w, newBase, baseInZip+file.Name()+"/")
		}
	}
}

func deployZipToSite(zipFilePath string) {
	// curl -X POST -H "Authorization: Bearer -BWqwl7FipqgTcJmzKl-GbDqwNIcFXAR853qg1itMVw" -H "Content-Type: application/zip" -d '{"files": "assets\test2.zip"}' https://api.netlify.com/api/v1/sites/23498d3f-a255-4471-980f-fe15896ef693

	zipFile, err := os.Open(zipFilePath)
	if err != nil {
		fmt.Println(err)
	}

	request, err := http.NewRequest(http.MethodPost, "https://api.netlify.com/api/v1/sites/23498d3f-a255-4471-980f-fe15896ef693/deploys", zipFile)
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Add("Authorization", "Bearer -BWqwl7FipqgTcJmzKl-GbDqwNIcFXAR853qg1itMVw")
	request.Header.Add("Content-Type", "application/zip")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Response:", response)

}
