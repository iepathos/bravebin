package bravebin

import (
	"os"
)

// func DebugMsgInstruction(msg string) ([]string, []string) {
// 	imports := []string{
// 		"\"fmt\"",
// 	}
// 	instructions := []string{
// 		fmt.Sprintf("fmt.Printf(\"%%v\", \"%v\\n\")", msg),
// 	}
// 	return imports, instructions
// }

type DnfInstruction struct {
	Packages []string
}

func (i DnfInstruction) install() {
	// dnf install bravebin filthymongrels
	i.Packages
}
