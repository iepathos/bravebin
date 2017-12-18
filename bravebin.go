package main

import (
	"fmt"
	core "github.com/iepathos/bravebin/core"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func GenerateGoMainPackage(imports []string, instructions []string) string {
	// package main
	// import (
	// {{ imports }}
	// )
	// func main() {
	// {{program}}
	// }
	imports = removeDuplicates(imports)
	content := []string{
		"package main",
		"import (",
		strings.Join(imports, ",\n"),
		")",
		"func main() {",
		strings.Join(instructions, "\n"),
		"}",
	}
	tmpfile, err := ioutil.TempFile("/tmp", "bravebin")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Write([]byte(strings.Join(content, "\n"))); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
	newname := tmpfile.Name() + ".go"
	os.Rename(tmpfile.Name(), newname)
	return newname
}

func BuildGofile(gofile string) string {
	// go build {{ gofile }}
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	os.Chdir(filepath.Dir(gofile))
	cmd := exec.Command("sh", "-c", fmt.Sprintf("go build %s", gofile))
	_, err = cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	os.Chdir(cwd)
	return gofile[0 : len(gofile)-3]
}

func main() {
	imports, instructions := core.CreateEchoInstruction("make devops better and braver")
	imports2, instructions2 := core.CreateEchoInstruction("and another instruction!")
	gofile := GenerateGoMainPackage(append(imports, imports2...), append(instructions, instructions2...))
	gobin := BuildGofile(gofile)
	fmt.Println(gobin)
}
