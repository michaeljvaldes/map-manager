package mapgen

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

func GenerateMaps(unminedPath string, worldPath string, siteDir string) {
	log.Println("Generating maps for world: " + worldPath)
	for _, mapConfig := range getAllMapConfigs() {
		generateMap(unminedPath, worldPath, siteDir, mapConfig)
	}
	log.Println("Finished generating maps for world: " + worldPath)

}

func generateMap(unminedPath string, worldPath string, siteDir string, mapConfig MapConfig) {
	log.Println("Generating map for dimension: " + mapConfig.Name)
	mapDir := filepath.Join(siteDir, mapConfig.Name)
	command := buildCommand(unminedPath, worldPath, mapDir, mapConfig)
	executeCommand(command)
}

func executeCommand(command *exec.Cmd) {
	err := command.Run()
	if err != nil {
		genErr := GenError{Err: err, Context: "Error executing unmined-cli command"}
		log.Fatal(genErr.ErrorMessage())
	}
}

func buildCommand(unminedPath string, worldPath string, mapDir string, mapConfig MapConfig) *exec.Cmd {
	args := buildArgs(worldPath, mapDir, mapConfig)
	commandString := unminedPath
	commandString += " " + strings.Join(args, " ")

	command := exec.Command(unminedPath)
	command.SysProcAttr = &syscall.SysProcAttr{}
	command.SysProcAttr.CmdLine = commandString
	return command
}

func buildArgs(worldPath string, mapDir string, mapConfig MapConfig) []string {
	webArg := "web"
	renderArg := "render"
	worldArg := buildStringArg("world", worldPath)
	outputArg := buildStringArg("output", mapDir)
	dimensionArg := buildStringArg("dimension", mapConfig.Dimension.toString())
	nightArg := buildBoolArg("night", mapConfig.Night)
	return []string{webArg, renderArg, worldArg, outputArg, dimensionArg, nightArg}
}

func buildStringArg(key, value string) string {
	return fmt.Sprintf(`--%s="%s"`, key, value)
}

func buildBoolArg(key string, value bool) string {
	valueStr := strconv.FormatBool(value)
	return fmt.Sprintf(`--%s=%s`, key, valueStr)
}
