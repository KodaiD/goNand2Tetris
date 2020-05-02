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
	text = text[:strings.Index(text, "//")]
	p.currentCommand = strings.TrimSpace(text)
}

func (p *Parser) CommandType() int {
	if p.currentCommand[0] == '@' {
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
	return p.currentCommand[10:13]
}

func (p *Parser) Comp() string {
	return p.currentCommand[3:10]
}

func (p *Parser) Jump() string {
	return p.currentCommand[13:]
}