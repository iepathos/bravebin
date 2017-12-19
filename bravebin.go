package main

import (
	"fmt"
	core "github.com/iepathos/bravebin/core"
	"strings"
)

func main() {
	// imports, instructions := core.DebugMsgInstruction("make devops better and braver")
	// imports2, instructions2 := core.DebugMsgInstruction("and another instruction!")
	// gofile := core.GenerateGoMainPackage(append(imports, imports2...), append(instructions, instructions2...))
	// gobin := core.BuildGofile(gofile)
	// fmt.Println(gobin)

	parser := core.BraveParser{
		"/home/cha0s/go/src/github.com/iepathos/bravebin/samples/simple_debug.yml",
		[]core.Instruction{},
	}
	instructions := parser.Parse()
	fmt.Printf("%v", instructions)

	goImports := []string{}
	goInstructions := []string{}

	for _, instruction := range instructions {
		if instruction.Module == "debug" {
			for _, arg := range instruction.Args {
				if strings.Contains(arg, "msg") {
					// fmt.Println(arg)
					msgIdx := strings.Index(arg, "msg")
					msg := ""
					if string(arg[msgIdx+4]) == "\"" {
						// ok, strip the quotes
						msg = strings.Replace(string(arg[msgIdx+4:]), "\"", "", 2)
					} else {
						msg = string(arg[msgIdx+4:])
					}
					imports, instructions := core.DebugMsgInstruction(msg)
					goImports = append(goImports, imports...)
					goInstructions = append(goInstructions, instructions...)

				}
			}
		}
	}
	// fmt.Printf("%v", goInstructions)
	gofile := core.GenerateGoMainPackage(goImports, goInstructions)
	gobin := core.BuildGofile(gofile)
	fmt.Println(gobin)
}
