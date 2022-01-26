package main

import (
	"fmt"
	"minecraftmapper/internal/args"
)

func main() {
	args := args.BuildArgsFromPromt()
	fmt.Println(args)
}
