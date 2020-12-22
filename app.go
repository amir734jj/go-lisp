package main

import (
	"fmt"
	. "go-lexer/src"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("code.txt")
	if err != nil {
		fmt.Print(err)
	}

	code := string(b)

	tokens := NewLexer().
		Add(Token{Name: "R_PRN", Pattern: "^[)]$"}).
		Add(Token{Name: "L_PRN", Pattern: "^[(]$"}).
		Add(Token{Name: "DEFUN", Pattern: "^defun$"}).
		Add(Token{Name: "COND", Pattern: "^if$"}).
		Add(Token{Name: "ID", Pattern: "^[a-zA-Z]+[\\w]*$"}).
		Add(Token{Name: "OP", Pattern: "^[-+*/]$"}).
		Add(Token{Name: "OP", Pattern: "^(<|<=|==|>|>=)$"}).
		Add(Token{Name: "NUM", Pattern: "^[0-9]+$"}).
		Add(Token{Name: "WS", Pattern: "^[\\s+]$", Ignore: true}).
		Add(Token{Name: "STRING", Pattern: "^\".*?\"$"}).
		Build()(code)

	for _, token := range tokens {
		fmt.Print(token.Value)
		fmt.Print(" ")
	}

	println()

	ast := Parse(tokens)
	println(VisitToString(ast))
}
