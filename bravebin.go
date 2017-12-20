package main

import (
	"fmt"
	core "github.com/iepathos/bravebin/core"
	"os"
)

func main() {
	// imports, instructions := core.DebugMsgInstruction("make devops better and braver")
	// imports2, instructions2 := core.DebugMsgInstruction("and another instruction!")
	// gofile := core.GenerateGoMainPackage(append(imports, imports2...), append(instructions, instructions2...))
	// gobin := core.BuildGofile(gofile)
	// fmt.Println(gobin)
	if len(os.Args) < 2 {
		fmt.Println("Missing playbook, pass yaml file as first argument to bravebin")
		os.Exit(2)
	}
	filename := string(os.Args[1])
	parser := core.BraveParser{
		filename, //"/home/cha0s/go/src/github.com/iepathos/bravebin/samples/simple_debug.yml",
		[]core.Instruction{},
	}
	instructions := parser.Parse()
	goImports, goInstructions := parser.ParseInstructions(instructions)

	gofile := core.GenerateGoMainPackage(goImports, goInstructions)
	gobin := core.BuildGofile(gofile)
	fmt.Println(gobin)
}
