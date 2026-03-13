package evaluator

import (
	"bytes"
	"birl-go/lexer"
	"birl-go/object"
	"birl-go/parser"
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
