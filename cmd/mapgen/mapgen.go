package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func generateMap(worldPath string, outputPath string, dimension Dimension, night bool) {
	command := buildCommand(worldPath, outputPath, dimension, night)
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

func buildCommand(worldPath string, outputPath string, dimension Dimension, night bool) *exec.Cmd {
	path := "C:/dev/go/minecraft-mapper/third_party/unmined/unmined-cli.exe"
	command := exec.Command(path)
	command.SysProcAttr = &syscall.SysProcAttr{}
	command.SysProcAttr.CmdLine = `C://dev//go//minecraft-mapper//third_party//unmined//unmined-cli.exe web render --world="C://Users//micha//AppData//Roaming//.minecraft//saves//World of Duane" --output="C://dev/go//minecraft-mapper//assets//test5//" --dimension=overworld --night=false --log-level=debug`
	return command
}
