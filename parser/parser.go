package parser

import (
	"strconv"
	"birl-go/ast"
	"birl-go/lexer"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  lexer.Token
	peekToken lexer.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != lexer.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case lexer.START_PROG:
		return p.parseBlockStatement(lexer.END_PROG)
	case lexer.TYPE_INT, lexer.TYPE_FLOAT, lexer.TYPE_CHAR:
		return p.parseVarStatement()
	case lexer.PRINT:
		return p.parsePrintStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseBlockStatement(endToken lexer.TokenType) *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for p.curToken.Type != endToken && p.curToken.Type != lexer.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	return block
}

func (p *Parser) parseVarStatement() *ast.VarStatement {
	stmt := &ast.VarStatement{Token: p.curToken}

	if p.peekToken.Type != lexer.IDENT {
		return nil
	}
	p.nextToken()

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if p.peekToken.Type != lexer.ASSIGN {
		return nil
	}
	p.nextToken()
	p.nextToken()

	stmt.Value = p.parseExpression()

	if p.peekToken.Type == lexer.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parsePrintStatement() *ast.PrintStatement {
	stmt := &ast.PrintStatement{Token: p.curToken}

	if p.peekToken.Type != lexer.LPAREN {
		return nil
	}
	p.nextToken()
	p.nextToken()

	stmt.Expression = p.parseExpression()

	if p.peekToken.Type != lexer.RPAREN {
		return nil
	}
	p.nextToken()

	if p.peekToken.Type == lexer.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression()
	if p.peekToken.Type == lexer.SEMICOLON {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpression() ast.Expression {
	// Simplificação extrema: apenas literais e expressões infixas básicas
	left := p.parsePrimaryExpression()

	for p.peekToken.Type == lexer.PLUS || p.peekToken.Type == lexer.MINUS || p.peekToken.Type == lexer.ASTERISK || p.peekToken.Type == lexer.SLASH {
		p.nextToken()
		op := p.curToken.Literal
		p.nextToken()
		right := p.parsePrimaryExpression()
		left = &ast.InfixExpression{
			Left:     left,
			Operator: op,
			Right:    right,
		}
	}

	return left
}

func (p *Parser) parsePrimaryExpression() ast.Expression {
	switch p.curToken.Type {
	case lexer.IDENT:
		return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	case lexer.INT:
		val, _ := strconv.ParseInt(p.curToken.Literal, 0, 64)
		return &ast.IntegerLiteral{Token: p.curToken, Value: val}
	case lexer.STRING:
		return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
	default:
		return nil
	}
}
