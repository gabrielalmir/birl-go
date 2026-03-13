package lexer

import (
	"strings"
)

type Lexer struct {
	input        string
	position     int  // posição atual na entrada (aponta para o caractere atual)
	readPosition int  // posição de leitura atual (aponta para após o caractere atual)
	ch           byte // caractere atual sendo examinado
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	// Tenta casar frases BIRL primeiro
	if l.isLetter(l.ch) {
		literal := l.readBIRLPhrase()
		if tokenType, ok := keywords[literal]; ok {
			tok.Type = tokenType
			tok.Literal = literal
			return tok
		}
		// Se não for frase BIRL, volta para o início do identificador e lê apenas o identificador
		// Na verdade, readBIRLPhrase já leu o "identificador" se não for frase.
		// Mas BIRL frases podem ter espaços. 
		// Vamos simplificar: se não é palavra-chave, tratamos como IDENT.
		tok.Literal = literal
		tok.Type = LookupIdent(literal)
		return tok
	}

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(PLUS, l.ch)
	case '-':
		tok = newToken(MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(BANG, l.ch)
		}
	case '/':
		tok = newToken(SLASH, l.ch)
	case '*':
		tok = newToken(ASTERISK, l.ch)
	case '<':
		tok = newToken(LT, l.ch)
	case '>':
		tok = newToken(GT, l.ch)
	case ';':
		tok = newToken(SEMICOLON, l.ch)
	case ',':
		tok = newToken(COMMA, l.ch)
	case '(':
		tok = newToken(LPAREN, l.ch)
	case ')':
		tok = newToken(RPAREN, l.ch)
	case '{':
		tok = newToken(LBRACE, l.ch)
	case '}':
		tok = newToken(RBRACE, l.ch)
	case '"':
		tok.Type = STRING
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			if strings.Contains(tok.Literal, ".") {
				tok.Type = FLOAT
			} else {
				tok.Type = INT
			}
			return tok
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readBIRLPhrase() string {
	position := l.position
	
	// Tentamos ler o máximo possível que possa ser uma palavra-chave BIRL
	// Palavras-chave BIRL podem conter espaços e símbolos como ? ou Ã
	
	// Uma estratégia melhor: se começar com uma das palavras-chave, 
	// verificamos se o prefixo bate com alguma keyword.
	
	remaining := l.input[position:]
	for phrase := range keywords {
		if strings.HasPrefix(strings.ToUpper(remaining), strings.ToUpper(phrase)) {
			// Consome os caracteres da frase
			for i := 0; i < len(phrase); i++ {
				l.readChar()
			}
			return phrase
		}
	}

	// Se não for frase, lê como identificador normal (até espaço ou pontuação)
	for l.isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch > 127 // Suporte a acentos
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9' || ch == '.'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}
