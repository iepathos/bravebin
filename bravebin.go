package main

import (
	"fmt"
	core "github.com/iepathos/bravebin/core"
)

func main() {
	imports, instructions := core.CreateEchoInstruction("make devops better and braver")
	imports2, instructions2 := core.CreateEchoInstruction("and another instruction!")
	gofile := core.GenerateGoMainPackage(append(imports, imports2...), append(instructions, instructions2...))
	gobin := core.BuildGofile(gofile)
	fmt.Println(gobin)
}
