package main

import (
	"fmt"
	"monkey/src/evaluator"
	"monkey/src/lexer"
	"monkey/src/object"
	"monkey/src/parser"
	"monkey/src/server"
	"os"
)

func main() {
	server.Start()
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
	evaluated := evaluator.Eval(program, env)

	fmt.Println(evaluated.Inspect())

	// user, err := user.Current()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Hello %s! This is the Monnkey Programming Language!\n", user.Username)
	// fmt.Printf("Feel free to type in commands\n")
	// repl.Start(os.Stdin, os.Stdout)
}
