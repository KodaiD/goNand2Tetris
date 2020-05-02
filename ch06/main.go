package main

import (
	"bufio"
	"os"
)

func main() {
	asmFile := os.Args[1]
	fp, err := os.Open(asmFile)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fp)
	parser := NewParser(scanner)

	for parser.HasMoreCommands() {
		parser.Advance()
		com := parser.CommandType()
		if com == CommandC {

		} else {

		}
	}
}