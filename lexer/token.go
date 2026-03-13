package lexer

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identificadores e Literais
	IDENT  = "IDENT"  // soma, a, b
	INT    = "INT"    // 13, 37
	FLOAT  = "FLOAT"  // 42.0
	STRING = "STRING" // "frango"

	// Operadores
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="
	AMPERSAND = "&"
	AND      = "&&"
	OR       = "||"

	// Delimitadores
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	// Palavras-chave BIRL
	START_PROG = "HORA DO SHOW"
	END_PROG   = "BIRL"
	PRINT      = "CE QUER VER ESSA PORRA?"
	READ       = "QUE QUE CE QUER MONSTRÃO?"
	IF         = "ELE QUE A GENTE QUER?"
	ELSE       = "NÃO VAI DAR NÃO"
	ELSE_IF    = "QUE NUM VAI DAR O QUE?"
	WHILE      = "NEGATIVA BAMBAM"
	FOR        = "MAIS QUERO MAIS"
	RETURN     = "BORA CUMPADE"
	BREAK      = "SAI FILHO DA PUTA"
	CONTINUE   = "VAMO MONSTRO"
	FUNC       = "OH O HOMEM AI PO"
	CALL       = "AJUDA O MALUCO TA DOENTE"

	// Tipos
	TYPE_INT   = "MONSTRO"
	TYPE_FLOAT = "TRAPÉZIO"
	TYPE_CHAR  = "FRANGO"
)

var keywords = map[string]TokenType{
	"HORA DO SHOW":            START_PROG,
	"BIRL":                    END_PROG,
	"CE QUER VER ESSA PORRA?": PRINT,
	"QUE QUE CE QUER MONSTRÃO?": READ,
	"ELE QUE A GENTE QUER?":    IF,
	"NÃO VAI DAR NÃO":         ELSE,
	"QUE NUM VAI DAR O QUE?":   ELSE_IF,
	"NEGATIVA BAMBAM":         WHILE,
	"MAIS QUERO MAIS":         FOR,
	"BORA CUMPADE":            RETURN,
	"SAI FILHO DA PUTA":       BREAK,
	"VAMO MONSTRO":            CONTINUE,
	"OH O HOMEM AI PO":        FUNC,
	"AJUDA O MALUCO TA DOENTE": CALL,
	"MONSTRO":                 TYPE_INT,
	"TRAPÉZIO":                TYPE_FLOAT,
	"FRANGO":                  TYPE_CHAR,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
