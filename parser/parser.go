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
	case lexer.READ:
		return p.parseReadStatement()
	case lexer.IF:
		return p.parseIfStatement()
	case lexer.WHILE:
		return p.parseWhileStatement()
	case lexer.RETURN:
		return p.parseReturnStatement()
	case lexer.FUNC:
		return p.parseExpressionStatement() // Deixa o parseExpression lidar com FUNC
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseIfStatement() *ast.IfStatement {
	stmt := &ast.IfStatement{Token: p.curToken}

	if p.peekToken.Type != lexer.LPAREN {
		return nil
	}
	p.nextToken()
	p.nextToken()

	stmt.Condition = p.parseExpression()

	if p.peekToken.Type != lexer.RPAREN {
		return nil
	}
	p.nextToken()

	stmt.Consequence = p.parseBlockStatement(lexer.END_PROG)

	if p.peekToken.Type == lexer.ELSE {
		p.nextToken()
		stmt.Alternative = p.parseBlockStatement(lexer.END_PROG)
	}

	if p.curToken.Type == lexer.END_PROG {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseWhileStatement() *ast.WhileStatement {
	stmt := &ast.WhileStatement{Token: p.curToken}

	if p.peekToken.Type != lexer.LPAREN {
		return nil
	}
	p.nextToken()
	p.nextToken()

	stmt.Condition = p.parseExpression()

	if p.peekToken.Type != lexer.RPAREN {
		return nil
	}
	p.nextToken()

	stmt.Body = p.parseBlockStatement(lexer.END_PROG)

	if p.curToken.Type == lexer.END_PROG {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReadStatement() *ast.ReadStatement {
	stmt := &ast.ReadStatement{Token: p.curToken}

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

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	if p.peekToken.Type != lexer.SEMICOLON {
		p.nextToken()
		stmt.ReturnValue = p.parseExpression()
	}

	if p.peekToken.Type == lexer.SEMICOLON {
		p.nextToken()
	}

	return stmt
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

	for p.peekToken.Type == lexer.PLUS || p.peekToken.Type == lexer.MINUS || 
		p.peekToken.Type == lexer.ASTERISK || p.peekToken.Type == lexer.SLASH ||
		p.peekToken.Type == lexer.EQ || p.peekToken.Type == lexer.NOT_EQ ||
		p.peekToken.Type == lexer.LT || p.peekToken.Type == lexer.GT ||
		p.peekToken.Type == lexer.ASSIGN {
		
		p.nextToken()
		op := p.curToken.Literal
		p.nextToken()
		
		var right ast.Expression
		if op == "=" {
			right = p.parseExpression() // Recursivo para associatividade à direita
			if ident, ok := left.(*ast.Identifier); ok {
				return &ast.AssignmentExpression{
					Token: p.curToken,
					Name:  ident,
					Value: right,
				}
			}
		} else {
			right = p.parsePrimaryExpression()
		}

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
	case lexer.FUNC:
		return p.parseFunctionLiteral()
	case lexer.CALL:
		return p.parseCallExpression()
	default:
		return nil
	}
}

func (p *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: p.curToken}

	// Tipo (ignorado por enquanto)
	p.nextToken()
	
	// Nome da função (pula o tipo e vai para o nome)
	p.nextToken()
	
	if p.curToken.Type != lexer.IDENT {
		return nil
	}
	lit.Name = p.curToken.Literal

	if p.peekToken.Type != lexer.LPAREN {
		return nil
	}
	p.nextToken()

	lit.Parameters = p.parseFunctionParameters()

	lit.Body = p.parseBlockStatement(lexer.END_PROG)
	
	if p.curToken.Type == lexer.END_PROG {
		p.nextToken()
	}

	return lit
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}

	if p.peekToken.Type == lexer.RPAREN {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	for {
		// Ignora o tipo do parâmetro
		p.nextToken() 
		
		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		identifiers = append(identifiers, ident)

		if p.peekToken.Type != lexer.COMMA {
			break
		}
		p.nextToken()
		p.nextToken()
	}

	if p.peekToken.Type != lexer.RPAREN {
		return nil
	}
	p.nextToken()

	return identifiers
}

func (p *Parser) parseCallExpression() ast.Expression {
	exp := &ast.CallExpression{Token: p.curToken}
	
	p.nextToken()
	exp.Function = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if p.peekToken.Type != lexer.LPAREN {
		return nil
	}
	p.nextToken()

	exp.Arguments = p.parseCallArguments()

	return exp
}

func (p *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}

	if p.peekToken.Type == lexer.RPAREN {
		p.nextToken()
		return args
	}

	p.nextToken()
	args = append(args, p.parseExpression())

	for p.peekToken.Type == lexer.COMMA {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression())
	}

	if p.peekToken.Type != lexer.RPAREN {
		return nil
	}
	p.nextToken()

	return args
}
