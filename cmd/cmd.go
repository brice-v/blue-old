package cmd

import (
	"blue/lexer"
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
