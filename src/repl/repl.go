package repl

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"monkey/src/evaluator"
	"monkey/src/lexer"
	"monkey/src/object"
	"monkey/src/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		fmt.Println(program.String())
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
		}

		evaluated := evaluator.Eval(program, env, &bytes.Buffer{})
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

const MONKEY_FACE = `MONKEY`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
