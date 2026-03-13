package lexer

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `HORA DO SHOW
    MONSTRO a = 10;
    FRANGO s = "birl";
    CE QUER VER ESSA PORRA?(a);
BIRL`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{START_PROG, "HORA DO SHOW"},
		{TYPE_INT, "MONSTRO"},
		{IDENT, "a"},
		{ASSIGN, "="},
		{INT, "10"},
		{SEMICOLON, ";"},
		{TYPE_CHAR, "FRANGO"},
		{IDENT, "s"},
		{ASSIGN, "="},
		{STRING, "birl"},
		{SEMICOLON, ";"},
		{PRINT, "CE QUER VER ESSA PORRA?"},
		{LPAREN, "("},
		{IDENT, "a"},
		{RPAREN, ")"},
		{SEMICOLON, ";"},
		{END_PROG, "BIRL"},
		{EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
