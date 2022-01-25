package main

import (
	_ "embed"
	"io"
	"os"
)

/*
take directory of map folders
turn into one html site with 4 links
zip
return zip file path
*/

//go:embed "assets/index_template.html"
var indexTemplate string

func PrepareMaps(baseDirectory string) string {
	copyIndexTemplate(baseDirectory)
	zipFile := "../../assets/site.zip"
	ZipFolder(baseDirectory, zipFile)
	return ""
}

func copyIndexTemplate(baseDirectory string) (int, error) {
	newIndexPath := baseDirectory + "/index.html"
	destination, err := os.Create(newIndexPath)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.WriteString(destination, indexTemplate)
	return nBytes, err
}
