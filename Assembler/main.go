package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	asmFile, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer asmFile.Close()

	binFile, err := os.Create("Prog.hack")
	if err != nil {
		panic(err)
	}
	defer binFile.Close()

	scanner := bufio.NewScanner(asmFile)
	writer := bufio.NewWriter(binFile)
	parser := NewParser(scanner)
	symbolTable := NewST()

	// 1st step
	romAddr := 0
	for parser.HasMoreCommands() {
		parser.Advance()
		switch parser.CommandType() {
		case CommandA, CommandC:
			romAddr += 1
		case CommandL:
			symbolTable.AddEntry(parser.Symbol(), romAddr)
		}
	}

	// 2nd step
	ramAddr := 0x0010
	var binary string
	asmFile.Seek(0, 0)
	scanner = bufio.NewScanner(asmFile)
	parser = NewParser(scanner)
	for parser.HasMoreCommands() {
		parser.Advance()
		switch parser.CommandType() {
		case CommandA:
			symbol := parser.Symbol()
			if addr, err := strconv.Atoi(symbol); err != nil {
				if symbolTable.Contains(symbol) {
					addr = symbolTable.GetAddress(symbol)
					binary = "0" + strconv.FormatInt(int64(addr), 2) + "\n"
				} else {
					symbolTable.AddEntry(symbol, ramAddr)
					binary = "0" + strconv.FormatInt(int64(ramAddr), 2) + "\n"
					ramAddr++
				}
			} else {
				binary = "0" + strconv.FormatInt(int64(addr), 2) + "\n"
			}
			for len(binary) != 17 {
				binary = "0" + binary
			}
			_, err = writer.Write([]byte(binary))
			if err != nil {
				panic(err)
			}
		case CommandC:
			comp, err := ConvertComp(parser.Comp())
			if err != nil {
				panic(err)
			}
			dest, err := ConvertDest(parser.Dest())
			if err != nil {
				panic(err)
			}
			jump, err := ConvertJump(parser.Jump())
			if err != nil {
				panic(err)
			}
			binary = "111" + comp + dest + jump + "\n"
			_, err = writer.Write([]byte(binary))
			if err != nil {
				panic(err)
			}
		case CommandL:
		}
		writer.Flush()
	}
}