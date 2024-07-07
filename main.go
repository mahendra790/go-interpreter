package main

import (
	"monkey/src/server"
)

func main() {

	server.RunServer()
	// js.Global().Set("execute", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	// 	if len(args) == 0 {
	// 		return "Hello, World!"
	// 	}
	// 	code := args[0].String()

	// 	l := lexer.New(code)
	// 	p := parser.New(l)

	// 	program := p.ParseProgram()

	// 	if len(p.Errors()) > 0 {
	// 		return strings.Join(p.Errors(), ", ")
	// 	}

	// 	buffer := bytes.Buffer{}

	// 	evaluated := evaluator.Eval(program, object.NewEnvironment(), &buffer)

	// 	return buffer.String() + evaluated.Inspect()
	// }))

	// select {}
}
