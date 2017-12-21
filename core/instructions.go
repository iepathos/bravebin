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

func ShellInstruction(shellCode string, sudo bool, first bool) ([]string, []string) {
	shellCode = EscapeShellCode(shellCode)
	imports := []string{
		"\"os/exec\"",
		"\"fmt\"",
	}
	instructions := []string{}
	cmdStr := ""
	if first {
		if sudo {
			cmdStr = fmt.Sprintf("cmd := exec.Command(\"/bin/sh\", \"-c\", \"sudo %v\")", shellCode)
		} else {
			cmdStr = fmt.Sprintf("cmd := exec.Command(\"/bin/sh\", \"-c\", \"%v\")", shellCode)
		}
		instructions = []string{
			cmdStr,
			fmt.Sprintf("out, err := cmd.Output()"),
			fmt.Sprintf("if err != nil {"),
			fmt.Sprintf("fmt.Println(err)"),
			fmt.Sprintf("}"),
			fmt.Sprintf("fmt.Printf(\"%%s\", out)"),
		}
	} else {
		if sudo {
			cmdStr = fmt.Sprintf("cmd = exec.Command(\"/bin/sh\", \"-c\", \"sudo %v\")", shellCode)
		} else {
			cmdStr = fmt.Sprintf("cmd = exec.Command(\"/bin/sh\", \"-c\", \"%v\")", shellCode)
		}
		instructions = []string{
			cmdStr,
			fmt.Sprintf("out, err = cmd.Output()"),
			fmt.Sprintf("if err != nil {"),
			fmt.Sprintf("fmt.Println(err)"),
			fmt.Sprintf("}"),
			fmt.Sprintf("fmt.Printf(\"%%s\", out)"),
		}
	}
	return imports, instructions
}
