package main

import (
	"embed"
	"fmt"
)

//go:embed example.txt
var content string

//go:embed basics
var basicsFolder embed.FS

func main() {

	// We can blank import a package for "side effects" aka, asking for the package's init() only

	fmt.Println("Embedded content:", content)

	content, err := basicsFolder.ReadFile("basics/panic.go")
	if err != nil {
		panic(err)
	}
	fmt.Println("File content:", string(content))
}
