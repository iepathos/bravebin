package bravebin

import (
	"fmt"
	"strings"
)

func EscapeShellCode(shellCode string) string {
	shellCode = strings.Replace(shellCode, "\"", "\\\"", -1)
	return shellCode
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

func ShellInstruction(shellCode string, first bool) ([]string, []string) {
	// escape double quotes in shellCode
	shellCode = EscapeShellCode(shellCode)
	imports := []string{
		"\"os/exec\"",
		"\"fmt\"",
	}
	instructions := []string{}
	if first {
		instructions = []string{
			fmt.Sprintf("cmd := exec.Command(\"/bin/sh\", \"-c\", \"%v\")", shellCode),
			fmt.Sprintf("out, _ := cmd.Output()"),
			fmt.Sprintf("fmt.Printf(\"%%s\", out)"),
		}
	} else {
		instructions = []string{
			fmt.Sprintf("cmd = exec.Command(\"/bin/sh\", \"-c\", \"%v\")", shellCode),
			fmt.Sprintf("out, _ = cmd.Output()"),
			fmt.Sprintf("fmt.Printf(\"%%s\", out)"),
		}
	}
	return imports, instructions
}
