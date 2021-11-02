package cmd

import (
	"blue/lexer"
	"blue/parser"
	"blue/token"
	"flag"
	"fmt"
	"os"
)

const VERSION = "v0.0.1"

func readAll(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("readAll: " + err.Error())
	}
	return string(data)
}

func Run(args []string) {
	lFlag := flag.String("l", "", "Enter a file to be lexed and printed to the screen")
	sFlag := flag.String("s", "", "Enter a file to be lexed and print illegal token spans")
	aFlag := flag.String("a", "", "Enter a file to be parsed and the ast printed to the screen")
	vFlag := flag.Bool("v", false, "Prints the version of blue to the screen")
	flag.Parse()
	if lFlag != nil && *lFlag != "" {
		lexFile(*lFlag)
	}
	if sFlag != nil && *sFlag != "" {
		spanFile(*sFlag)
	}
	if vFlag != nil && *vFlag {
		fmt.Println(VERSION)
	}
	if aFlag != nil && *aFlag != "" {
		parseFile(*aFlag)
	}
}

func lexFile(filename string) {
	l := lexer.New(readAll(filename), filename)
	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Println(tok.String())
	}
}

func spanFile(filename string) {
	input := readAll(filename)
	l := lexer.New(input, filename)
	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		if tok.Type == token.ILLEGAL {
			msg := "lexer error encountered"
			msgToPrint := l.GetSpanPrintable(tok.Span, msg)
			fmt.Print(msgToPrint)
		}
	}
}

func parseFile(filename string) {
	input := readAll(filename)
	l := lexer.New(input, filename)
	p := parser.New(l)
	ast := p.ParseProgram()
	fmt.Println(ast.Display())
}
