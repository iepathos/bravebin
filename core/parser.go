package bravebin

import (
	// "encoding/json"
	// "bufio"
	// "fmt"
	"io/ioutil"
	// "os"
	"strings"
)

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

type Instruction struct {
	Module string
	Args   []string
}

type BraveParser struct {
	YmlPath      string
	Instructions []Instruction
}

func isTopLevelInstruction(line string) bool {
	if strings.HasPrefix(string(line), "- ") {
		return true
	}
	return false
}

func (bp BraveParser) Read() []byte {
	data, err := ioutil.ReadFile(bp.YmlPath)
	if err != nil {
		panic(err)
	}
	return data
}

func (bp BraveParser) Parse() []Instruction {
	data := string(bp.Read())
	instructions := []Instruction{}
	args := []string{}
	for _, line := range strings.Split(data, "\n") {
		// ignore blank lines
		if len(strings.TrimSpace(line)) > 0 {
			if isTopLevelInstruction(line) {
				lsplit := strings.Split(line, ":")
				instruction := lsplit[0][2:]

				if len(lsplit) > 0 {
					afterColon := strings.TrimSpace(strings.Join(lsplit[1:], " "))
					args = []string{afterColon}
				}
				i := Instruction{instruction, args}
				instructions = append(instructions, i)
			} else {
				// must be argument for previous instruction
				if len(instructions) > 0 {
					lastIdx := len(instructions) - 1
					instructions[lastIdx].Args = append(instructions[lastIdx].Args, strings.TrimSpace(line))
				}
			}
		}
	}

	bp.Instructions = instructions
	return instructions
}

func (bp BraveParser) ParseInstructions(instructions []Instruction) ([]string, []string) {
	goImports := []string{}
	goInstructions := []string{}

	firstShellCmd := true
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
						msg = strings.TrimSpace(string(arg[msgIdx+4:]))
					}
					imports, instructions := DebugMsgInstruction(msg)
					goImports = append(goImports, imports...)
					goInstructions = append(goInstructions, instructions...)
				}
			}
		} else if instruction.Module == "shell" {
			sudo := false
			for _, arg := range instruction.Args {
				sudo = false
				if strings.Contains(arg, "sudo") {
					if strings.Contains(arg, "=") {
						// this is broken for multiple args on one line need smarter parsing for this case

						if strings.TrimSpace(strings.Split(arg, "=")[1]) == "yes" || strings.TrimSpace(strings.Split(arg, "=")[1]) == "true" {
							sudo = true
						}
					} else if strings.Contains(arg, ":") {
						if strings.TrimSpace(strings.Split(arg, ":")[1]) == "yes" || strings.TrimSpace(strings.Split(arg, ":")[1]) == "true" {
							sudo = true
						}
					}
				}
			}
			imports, instructions := ShellInstruction(instruction.Args[0], sudo, firstShellCmd)
			goImports = append(goImports, imports...)
			goInstructions = append(goInstructions, instructions...)
			if firstShellCmd {
				firstShellCmd = false
			}
		}
	}
	return goImports, goInstructions
}
