package main

import (
	"bicep-doc/parser"
	"bicep-doc/printer"
	"fmt"
	"os"
)

func printUsage() {
	fmt.Println("usage: bicep-doc path/to/file.bicep")
	fmt.Println(`
This utility is used to generate input and output markdown documentation for a bicep file.
Run this program by passing a path to a bicep file on the command line. The program will then
write markdown tables for inputs and outputs to stdout.
	`)
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	p := parser.NewParser(os.Args[1])
	p.ProcessFile()

	printer.PrintMarkdown(p)
}
