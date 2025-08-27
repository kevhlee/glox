package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kevhlee/glox/pkg/lox"
)

func main() {
	globals := lox.NewEnvironment()

	if len(os.Args) == 1 {
		runREPL(globals)
	} else {
		runFile(globals, os.Args[1])
	}
}

func runREPL(globals *lox.Environment) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">> ")
		if !reader.Scan() {
			break
		}
		lox.RunSource(globals, reader.Text())
	}
}

func runFile(globals *lox.Environment, filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read file '%s'", filename)
		os.Exit(74)
	}
	os.Exit(lox.RunSource(globals, string(data)))
}
