package main

import (
	"fmt"
	"os"
	"io"

	"birl-go/evaluator"
	"birl-go/lexer"
	"birl-go/object"
	"birl-go/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("ERRO: CADÊ O ARQUIVO, CUMPADE?")
		fmt.Println("Uso: birl <arquivo.birl>")
		os.Exit(1)
	}

	filename := os.Args[1]
	code, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("ERRO: NÃO CONSEGUI LER ESSA PORRA (%s)\n", filename)
		os.Exit(1)
	}

	runBIRL(string(code), os.Stdout)
}

func runBIRL(input string, out io.Writer) {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		printParserErrors(out, p.Errors())
		return
	}

	env := object.NewEnvironment()
	eval := evaluator.New(out)
	
	result := eval.Eval(program, env)

	if result != nil && result.Type() == object.ERROR_OBJ {
		fmt.Fprintf(out, "ERRO DE EXECUÇÃO: %s\n", result.Inspect())
	}
}

func printParserErrors(out io.Writer, errors []string) {
	fmt.Fprintf(out, "ERRO DE COMPILAÇÃO, MONSTRO!\n")
	for _, msg := range errors {
		fmt.Fprintf(out, "\t%s\n", msg)
	}
}
