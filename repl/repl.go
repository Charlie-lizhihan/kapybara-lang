package repl

import (
	"bufio"
	"fmt"
	"io"
	"kapybara-lang/evaluator"
	"kapybara-lang/lexer"
	"kapybara-lang/object"
	"kapybara-lang/parser"
)

const PROMPT = `>>`
const KAPYBARA = `／* ｀ ｴ´）www`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserError(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserError(out io.Writer, errors []string) {
	io.WriteString(out, KAPYBARA)
	io.WriteString(out, "\nWoops! Kapybara can't handle your code!\n")
	io.WriteString(out, "parser errors:\n")

	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
