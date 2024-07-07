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
	buffer := &bytes.Buffer{}
	evaluated := evaluator.Eval(program, env, buffer)

	fmt.Println(buffer.String())
	fmt.Println(evaluated.Inspect())
}
