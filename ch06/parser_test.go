package main

import (
	"bufio"
	"strings"
	"testing"
)

type expectData struct {
	command int
	symbol  string
	dest    string
	comp    string
	jump    string
}

var asmCode = `// comment line
	@i
	M=1 // i=1
	@sum
	M=0 // sum=0
(LOOP)
	@i
	AMD=D-A;JMP
	0;JMP`

var expects = []expectData{
	{0, "", "", "", ""},
	{CommandA, "i", "", "", ""},
	{CommandC, "", "M", "1", ""},
	{CommandA, "sum", "", "", ""},
	{CommandC, "", "M", "0", ""},
	{CommandL, "LOOP", "", "", ""},
	{CommandA, "i", "", "", ""},
	{CommandC, "", "AMD", "D-A", "JMP"},
	{CommandC, "", "", "0", "JMP"},
}

func TestParser(t *testing.T) {
	code := strings.NewReader(asmCode)
	scanner := bufio.NewScanner(code)
	parser := NewParser(scanner)

	for line := range expects {
		if !parser.HasMoreCommands() {
			t.Fatalf("line %v is not found", line)
		}

		parser.Advance()

		currentCommand := parser.CommandType()
		if currentCommand != expects[line].command {
			t.Fatalf("Command %v is unexpected. line: %v", currentCommand, line)
		}

		switch currentCommand {
		case CommandA, CommandL:
			if parser.Symbol() != expects[line].symbol {
				t.Fatalf("Symbol %v is unexpected. line: %v", parser.Symbol(), line)
			}
		case CommandC:
			if parser.Dest() != expects[line].dest {
				t.Fatalf("Dest %v is wrong. line: %v", parser.Dest(), line)
			}
			if parser.Comp() != expects[line].comp {
				t.Fatalf("Comp %v is wrong. line: %v", parser.Comp(), line)
			}
			if parser.Jump() != expects[line].jump {
				t.Fatalf("Jump %v is wrong. line: %v", parser.Jump(), line)
			}
		}
	}
}
