package lexer

import (
	"sort"
	"strings"
)

type Lexer struct {
	input        string
	position     int  // posição atual na entrada (aponta para o caractere atual)
	readPosition int  // posição de leitura atual (aponta para após o caractere atual)
	ch           byte // caractere atual sendo examinado
	line         int  // Linha atual
}

func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1}
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
		line := l.line
		literal := l.readBIRLPhrase()
		if tokenType, ok := keywords[literal]; ok {
			tok.Type = tokenType
			tok.Literal = literal
			tok.Line = line
			return tok
		}
		tok.Literal = literal
		tok.Type = LookupIdent(literal)
		tok.Line = line
		return tok
	}

	tok.Line = l.line

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: EQ, Literal: string(ch) + string(l.ch), Line: tok.Line}
		} else {
			tok = newToken(ASSIGN, l.ch, tok.Line)
		}
	case '+':
		tok = newToken(PLUS, l.ch, tok.Line)
	case '-':
		tok = newToken(MINUS, l.ch, tok.Line)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: NOT_EQ, Literal: string(ch) + string(l.ch), Line: tok.Line}
		} else {
			tok = newToken(BANG, l.ch, tok.Line)
		}
	case '/':
		tok = newToken(SLASH, l.ch, tok.Line)
	case '*':
		tok = newToken(ASTERISK, l.ch, tok.Line)
	case '<':
		tok = newToken(LT, l.ch, tok.Line)
	case '>':
		tok = newToken(GT, l.ch, tok.Line)
	case ';':
		tok = newToken(SEMICOLON, l.ch, tok.Line)
	case ',':
		tok = newToken(COMMA, l.ch, tok.Line)
	case '(':
		tok = newToken(LPAREN, l.ch, tok.Line)
	case ')':
		tok = newToken(RPAREN, l.ch, tok.Line)
	case '{':
		tok = newToken(LBRACE, l.ch, tok.Line)
	case '}':
		tok = newToken(RBRACE, l.ch, tok.Line)
	case ']':
		tok = newToken(RBRACKET, l.ch, tok.Line)
	case '[':
		tok = newToken(LBRACKET, l.ch, tok.Line)
	case '&':
		if l.peekChar() == '&' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: AND, Literal: string(ch) + string(l.ch), Line: tok.Line}
		} else {
			tok = newToken(AMPERSAND, l.ch, tok.Line)
		}
	case '|':
		if l.peekChar() == '|' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: OR, Literal: string(ch) + string(l.ch), Line: tok.Line}
		} else {
			tok = newToken(ILLEGAL, l.ch, tok.Line)
		}
	case ':':
		tok = newToken(COLON, l.ch, tok.Line)
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
			tok = newToken(ILLEGAL, l.ch, tok.Line)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for {
		if l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
			l.readChar()
		} else if l.ch == '\n' {
			l.line++
			l.readChar()
		} else if l.ch == '/' && l.peekChar() == '/' {
			l.skipSingleLineComment()
		} else if l.ch == '/' && l.peekChar() == '*' {
			l.skipMultiLineComment()
		} else {
			break
		}
	}
}

func (l *Lexer) skipSingleLineComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
}

func (l *Lexer) skipMultiLineComment() {
	l.readChar() // Pula o '/'
	l.readChar() // Pula o '*'
	for l.ch != 0 {
		if l.ch == '\n' {
			l.line++
		}
		if l.ch == '*' && l.peekChar() == '/' {
			l.readChar() // Pula o '*'
			l.readChar() // Pula o '/'
			break
		}
		l.readChar()
	}
}

func (l *Lexer) readBIRLPhrase() string {
	position := l.position
	remaining := l.input[position:]

	// Pegamos todas as keywords e ordenamos pelo tamanho (decrescente)
	// para garantir que frases longas como "BORA DIVIDIR O PESO"
	// sejam testadas antes de frases curtas como "BIRL".
	keys := make([]string, 0, len(keywords))
	for k := range keywords {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return len(keys[i]) > len(keys[j])
	})

	for _, phrase := range keys {
		if strings.HasPrefix(strings.ToUpper(remaining), strings.ToUpper(phrase)) {
			// Consome os caracteres da frase
			for i := 0; i < len(phrase); i++ {
				l.readChar()
			}
			return phrase
		}
	}

	// Se não for frase, lê como identificador normal
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

func newToken(tokenType TokenType, ch byte, line int) Token {
	return Token{Type: tokenType, Literal: string(ch), Line: line}
}
