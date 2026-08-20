package evaluator

import (
	"birl-go/lexer"
	"birl-go/object"
	"birl-go/parser"
	"bytes"
	"strings"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"HORA DO SHOW MONSTRO a = 5; BORA CUMPADE a; BIRL", 5},
		{"HORA DO SHOW MONSTRO a = 10; BORA CUMPADE a * 2; BIRL", 20},
		{"HORA DO SHOW MONSTRO a = 5; MONSTRO b = 5; BORA CUMPADE a + b; BIRL", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestFunctionCallRejectsWrongArgumentCount(t *testing.T) {
	function := `HORA DO SHOW
	OH O HOMEM AI PO MONSTRO SOMA(MONSTRO a, MONSTRO b)
		BORA CUMPADE a + b;
	BIRL`
	tests := []struct {
		name     string
		call     string
		expected string
	}{
		{name: "argumento ausente", call: "AJUDA O MALUCO TA DOENTE SOMA(10);", expected: "ESPERAVA 2 ARGUMENTO(S), MAS RECEBEU 1"},
		{name: "argumento excedente", call: "AJUDA O MALUCO TA DOENTE SOMA(10, 20, 30);", expected: "ESPERAVA 2 ARGUMENTO(S), MAS RECEBEU 3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evaluated := testEval(function + "\nMONSTRO resultado = " + tt.call + "\nBIRL")
			err, ok := evaluated.(*object.Error)
			if !ok {
				t.Fatalf("object is not Error. got=%T (%+v)", evaluated, evaluated)
			}
			if !strings.Contains(err.Message, tt.expected) {
				t.Errorf("error has wrong message. got=%q, want to contain %q", err.Message, tt.expected)
			}
		})
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	eval := New(&bytes.Buffer{})

	return eval.Eval(program, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}
	return true
}
