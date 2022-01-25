package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

func GenerateMaps(worldPath string, outputPath string) {
	for _, mapConfig := range getAllMapConfigs() {
		generateMap(worldPath, outputPath, mapConfig)
	}
}

func generateMap(worldPath string, outputPath string, mapConfig MapConfig) {
	command := buildCommand(worldPath, outputPath, mapConfig)
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

func buildCommand(worldPath string, outputPath string, mapConfig MapConfig) *exec.Cmd {
	path := "C:/dev/go/minecraft-mapper/third_party/unmined/unmined-cli.exe"
	worldArg := fmt.Sprintf(`--world="%s"`, worldPath)
	outputArg := fmt.Sprintf(`--output="%s"`, outputPath+"//"+mapConfig.Name+"//")
	dimensionArg := fmt.Sprintf(`--dimension=%s`, mapConfig.Dimension.toString())
	nightArg := fmt.Sprintf(`--night=%s`, strconv.FormatBool(mapConfig.Night))

	commandString := fmt.Sprintf(`%s web render %s %s %s %s`, path, worldArg, outputArg, dimensionArg, nightArg)
	command := exec.Command(path)
	command.SysProcAttr = &syscall.SysProcAttr{}
	command.SysProcAttr.CmdLine = commandString
	return command
}

func DeleteMapsDirectory(mapDirectory string) {
	os.RemoveAll(mapDirectory)
}
