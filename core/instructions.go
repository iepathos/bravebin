package bravebin

import (
	"fmt"
)

func DebugMsgInstruction(msg string) ([]string, []string) {
	imports := []string{
		"\"fmt\"",
	}
	instructions := []string{
		fmt.Sprintf("fmt.Printf(\"%%v\", \"%v\\n\")", msg),
	}
	return imports, instructions
}
