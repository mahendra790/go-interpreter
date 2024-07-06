package main

import (
	"bytes"
	"fmt"
	"monkey/src/evaluator"
	"monkey/src/lexer"
	"monkey/src/object"
	"monkey/src/parser"
	"os"
)

func main() {
	content, err := os.ReadFile("p.mky")
	if err != nil {
		fmt.Printf("error")
		return
	}

	str := string(content)

	l := lexer.New(str)
	p := parser.New(l)

	program := p.ParseProgram()

	env := object.NewEnvironment()
	evaluated := evaluator.Eval(program, env, &bytes.Buffer{})

	fmt.Println(evaluated.Inspect())
}
