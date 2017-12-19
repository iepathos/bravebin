package bravebin

import (
	"fmt"
)

type DebugModule struct {
	Msg string
}

func DebugMsgInstruction(msg string) ([]string, []string) {
	imports := []string{
		"\"fmt\"",
	}
	instructions := []string{
		fmt.Sprintf("fmt.Printf(\"%%v\", \"%v\\n\")", msg),
	}
	return imports, instructions
}
