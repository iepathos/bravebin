package main

import (
	"fmt"
	core "github.com/iepathos/bravebin/core"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing playbook, pass yaml file as first argument to bravebin")
		os.Exit(2)
	}
	filename := string(os.Args[1])
	parser := core.BraveParser{
		filename,
		[]core.Instruction{},
	}
	instructions := parser.Parse()
	goImports, goInstructions := parser.ParseInstructions(instructions)

	gofile := core.GenerateGoMainPackage(goImports, goInstructions)
	gobin := core.BuildGofile(gofile)
	fmt.Println(gobin)
}
