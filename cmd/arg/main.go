package main

import (
	"fmt"
	"minecraftmapper/internal/args"
)

func main() {
	args := args.BuildArgsFromPrompt()
	fmt.Println(args)
}
