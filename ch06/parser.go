package main

import (
	"bufio"
	"strings"
)

const (
	CommandA int = 1
	CommandC int = 2
	CommandL int = 3
)

type Parser struct {
	scanner *bufio.Scanner
	currentCommand string
}

type ParserInterface interface {
	HasMoreCommands() bool
	Advance()
	CommandType() int
	Symbol() string
	Dest() string
	Comp() string
	Jump() string
}

func NewParser(scanner *bufio.Scanner) *Parser {
	return &Parser{scanner, ""}
}

func (p *Parser) HasMoreCommands() bool {
	return p.scanner.Scan()
}

func (p *Parser) Advance() {
	text := p.scanner.Text()
	if len(text) == 0 {
		return
	}
	commentIndex :=strings.Index(text, "//")
	if commentIndex != -1 {
		text = text[:commentIndex]
	}
	p.currentCommand = strings.TrimSpace(text)
}

func (p *Parser) CommandType() int {
	if len(p.currentCommand) == 0 {
		return 0
	}else if p.currentCommand[0] == '@' {
		return CommandA
	} else if p.currentCommand[0] == '(' {
		return CommandL
	} else {
		return CommandC
	}
}

func (p *Parser) Symbol() string {
	if p.currentCommand[0] == '@' {
		return p.currentCommand[1:]
	} else if p.currentCommand[0] == '(' {
		return p.currentCommand[1:len(p.currentCommand)-1]
	}
	return ""
}

func (p *Parser) Dest() string {
	return helper(p.currentCommand)[0]
}

func (p *Parser) Comp() string {
	return helper(p.currentCommand)[1]
}

func (p *Parser) Jump() string {
	return helper(p.currentCommand)[2]
}

// helper returns [dest, comp, jump]
// dest=comp;jump
func helper(s string) []string {
	indexOfSemicolon := strings.Index(s, ";")
	indexOfEquation := strings.Index(s, "=")
	if indexOfSemicolon == -1 {
		return []string{s[:indexOfEquation], s[indexOfEquation+1:], ""}
	} else if indexOfEquation == -1 {
		return []string{"", s[:indexOfSemicolon], s[indexOfSemicolon+1:]}
	} else {
		return []string{s[:indexOfEquation], s[indexOfEquation+1:indexOfSemicolon], s[indexOfSemicolon+1:]}
	}
}