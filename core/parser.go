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

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
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
	invalidInstructionStarts := []string{
		"-",
		"",
	}
	invalidTopInstructionStarts := append(invalidInstructionStarts, " ")
	if StringInSlice(string(line[0]), invalidTopInstructionStarts) {
		return false
	}
	return true
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
				instruction := lsplit[0]

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
	// fmt.Printf("%v", instructions)
	return instructions
}
