package main

import (
	"minecraftmapper/pkg/mapdeploy"
)

func main() {
	zipFile := "C://dev//go//minecraft-mapper//test//site.zip"
	siteId := "23498d3f-a255-4471-980f-fe15896ef693"
	deployToken := "-BWqwl7FipqgTcJmzKl-GbDqwNIcFXAR853qg1itMVw"
	mapdeploy.DeployMapSite(zipFile, siteId, deployToken)
}
