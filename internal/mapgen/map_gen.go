package mapgen

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

func GenerateMaps(unminedPath string, worldPath string, siteDir string) {
	for _, mapConfig := range getAllMapConfigs() {
		generateMap(unminedPath, worldPath, siteDir, mapConfig)
	}
}

func generateMap(unminedPath string, worldPath string, siteDir string, mapConfig MapConfig) {
	mapDir := filepath.Join(siteDir, mapConfig.Name)
	command := buildCommand(unminedPath, worldPath, mapDir, mapConfig)
	fmt.Printf("Command: %q\n", command.String())
	executeCommand(command)
}

func executeCommand(command *exec.Cmd) {
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		fmt.Println("Error: " + err.Error())
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
