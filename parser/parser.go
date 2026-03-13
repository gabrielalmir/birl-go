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
	case lexer.FOR:
		return p.parseForStatement()
	case lexer.BREAK:
		return p.parseBreakStatement()
	case lexer.CONTINUE:
		return p.parseContinueStatement()
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
		p.peekError(lexer.LPAREN)
		return nil
	}
	p.nextToken()
	p.nextToken()

	stmt.Condition = p.parseExpression(LOWEST)

	if p.peekToken.Type != lexer.RPAREN {
		p.peekError(lexer.RPAREN)
		return nil
	}
	p.nextToken()

	stmt.Consequence = p.parseBlockStatement(lexer.END_PROG)

	for p.peekToken.Type == lexer.ELSE_IF {
		p.nextToken()
		
		if p.peekToken.Type != lexer.LPAREN {
			p.peekError(lexer.LPAREN)
			return nil
		}
		p.nextToken()
		p.nextToken()
		
		condition := p.parseExpression(LOWEST)
		
		if p.peekToken.Type != lexer.RPAREN {
			p.peekError(lexer.RPAREN)
			return nil
		}
		p.nextToken()
		
		consequence := p.parseBlockStatement(lexer.END_PROG)
		stmt.ElseIfs = append(stmt.ElseIfs, &ast.IfElseIf{
			Condition:   condition,
			Consequence: consequence,
		})
	}

	if p.peekToken.Type == lexer.ELSE {
		p.nextToken()
		stmt.Alternative = p.parseBlockStatement(lexer.END_PROG)
	}

	return stmt
}

func (p *Parser) peekError(t lexer.TokenType) {
	msg := "ESPERAVA " + string(t) + " MAS RECEBI " + string(p.peekToken.Type) + ", SEU FRANGO!"
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseWhileStatement() *ast.WhileStatement {
	stmt := &ast.WhileStatement{Token: p.curToken}

	if p.peekToken.Type != lexer.LPAREN {
		return nil
	}
	p.nextToken()
	p.nextToken()

	stmt.Condition = p.parseExpression(LOWEST)

	if p.peekToken.Type != lexer.RPAREN {
		return nil
	}
	p.nextToken()

	stmt.Body = p.parseBlockStatement(lexer.END_PROG)

	return stmt
}

func (p *Parser) parseBreakStatement() *ast.BreakStatement {
	stmt := &ast.BreakStatement{Token: p.curToken}
	if p.peekToken.Type == lexer.SEMICOLON {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseContinueStatement() *ast.ContinueStatement {
	stmt := &ast.ContinueStatement{Token: p.curToken}
	if p.peekToken.Type == lexer.SEMICOLON {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseForStatement() *ast.ForStatement {
	stmt := &ast.ForStatement{Token: p.curToken}

	if p.peekToken.Type != lexer.LPAREN {
		p.peekError(lexer.LPAREN)
		return nil
	}
	p.nextToken()
	p.nextToken() // Move para o início da Init

	// Parse Init
	if p.curToken.Type != lexer.SEMICOLON {
		stmt.Init = p.parseStatement()
		p.nextToken() // Pula o delimitador do statement para chegar na condição
	} else {
		p.nextToken() // Pula o ';' vazio
	}

	// Parse Condition
	if p.curToken.Type != lexer.SEMICOLON {
		stmt.Condition = p.parseExpression(LOWEST)
		p.nextToken() // Move para o próximo token
		if p.curToken.Type == lexer.SEMICOLON {
			p.nextToken() // Pula o ';'
		}
	} else {
		p.nextToken() // Pula o ';' vazio
	}

	// Parse Update
	if p.curToken.Type != lexer.RPAREN {
		stmt.Update = p.parseStatement()
		p.nextToken() // Move para o RPAREN (espero)
	}

	if p.curToken.Type != lexer.RPAREN {
		p.errors = append(p.errors, "ESPERAVA ) NO FINAL DO FOR, FRANGO! RECEBI: "+p.curToken.Literal)
		return nil
	}
	// Não precisamos de nextToken aqui, pois o parseBlockStatement já avança 1 token no início.
	
	stmt.Body = p.parseBlockStatement(lexer.END_PROG)

	return stmt
}

func (p *Parser) parseReadStatement() *ast.ReadStatement {
	stmt := &ast.ReadStatement{Token: p.curToken}

	if p.peekToken.Type != lexer.LPAREN {
		return nil
	}
	p.nextToken()
	p.nextToken()

	stmt.Expression = p.parseExpression(LOWEST)

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
		stmt.ReturnValue = p.parseExpression(LOWEST)
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

	stmt.Value = p.parseExpression(LOWEST)

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

	stmt.Expression = p.parseExpression(LOWEST)

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
	stmt.Expression = p.parseExpression(LOWEST)
	if p.peekToken.Type == lexer.SEMICOLON {
		p.nextToken()
	}
	return stmt
}

const (
	_ int = iota
	LOWEST
	ASSIGN      // =
	LOGICAL_OR  // ||
	LOGICAL_AND // &&
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X or &X or *X
	CALL        // myFunction(X)
	INDEX       // array[index]
)

var precedences = map[lexer.TokenType]int{
	lexer.ASSIGN:   ASSIGN,
	lexer.OR:       LOGICAL_OR,
	lexer.AND:      LOGICAL_AND,
	lexer.EQ:       EQUALS,
	lexer.NOT_EQ:   EQUALS,
	lexer.LT:       LESSGREATER,
	lexer.GT:       LESSGREATER,
	lexer.PLUS:     SUM,
	lexer.MINUS:    SUM,
	lexer.SLASH:    PRODUCT,
	lexer.ASTERISK: PRODUCT,
	lexer.LPAREN:   CALL,
	lexer.LBRACKET: INDEX,
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	left := p.parsePrefix()
	if left == nil {
		return nil
	}

	for p.peekToken.Type != lexer.SEMICOLON && p.peekToken.Type != lexer.RPAREN && precedence < p.peekPrecedence() {
		switch p.peekToken.Type {
		case lexer.PLUS, lexer.MINUS, lexer.ASTERISK, lexer.SLASH, lexer.EQ, lexer.NOT_EQ, lexer.LT, lexer.GT, lexer.ASSIGN, lexer.AND, lexer.OR:
			p.nextToken()
			left = p.parseInfixExpression(left)
		case lexer.LBRACKET:
			p.nextToken()
			left = p.parseIndexExpression(left)
		case lexer.LPAREN:
			return left
		default:
			return left
		}
	}

	return left
}

func (p *Parser) parsePrefix() ast.Expression {
	switch p.curToken.Type {
	case lexer.BANG, lexer.AMPERSAND, lexer.ASTERISK:
		return p.parsePrefixExpression()
	default:
		return p.parsePrimaryExpression()
	}
}

func (p *Parser) parsePrefixExpression() *ast.PrefixExpression {
	pe := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}
	p.nextToken()
	pe.Right = p.parseExpression(PREFIX)
	return pe
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	op := p.curToken.Literal
	precedence := p.curPrecedence()
	
	p.nextToken()
	right := p.parseExpression(precedence)
	
	if op == "=" {
		if ident, ok := left.(*ast.Identifier); ok {
			return &ast.AssignmentExpression{
				Token: p.curToken,
				Name:  ident,
				Value: right,
			}
		}
		// Suporte para atribuição em array: arr[0] = 1
		if _, ok := left.(*ast.IndexExpression); ok {
			// Por simplicidade, vamos tratar como uma atribuição especial ou ignorar por agora
			// BIRL é para ser rústico!
		}
	}

	return &ast.InfixExpression{
		Token:    p.curToken,
		Left:     left,
		Operator: op,
		Right:    right,
	}
}

func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	ie := &ast.IndexExpression{Token: p.curToken, Left: left}
	p.nextToken()
	ie.Index = p.parseExpression(LOWEST)
	if p.peekToken.Type != lexer.RBRACKET {
		return nil
	}
	p.nextToken()
	return ie
}

func (p *Parser) parsePrimaryExpression() ast.Expression {
	switch p.curToken.Type {
	case lexer.IDENT:
		return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	case lexer.INT:
		val, _ := strconv.ParseInt(p.curToken.Literal, 0, 64)
		return &ast.IntegerLiteral{Token: p.curToken, Value: val}
	case lexer.FLOAT:
		val, _ := strconv.ParseFloat(p.curToken.Literal, 64)
		return &ast.FloatLiteral{Token: p.curToken, Value: val}
	case lexer.STRING:
		return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
	case lexer.FUNC:
		return p.parseFunctionLiteral()
	case lexer.CALL:
		return p.parseCallExpression()
	case lexer.LBRACKET:
		return p.parseArrayLiteral()
	case lexer.LBRACE:
		return p.parseHashLiteral()
	default:
		return nil
	}
}

func (p *Parser) parseHashLiteral() ast.Expression {
	hash := &ast.HashLiteral{Token: p.curToken}
	hash.Pairs = make(map[ast.Expression]ast.Expression)

	for p.peekToken.Type != lexer.RBRACE {
		p.nextToken()
		key := p.parseExpression(LOWEST)

		if p.peekToken.Type != lexer.COLON {
			return nil
		}
		p.nextToken()
		p.nextToken()

		value := p.parseExpression(LOWEST)
		hash.Pairs[key] = value

		if p.peekToken.Type != lexer.RBRACE && p.peekToken.Type != lexer.COMMA {
			return nil
		}
		if p.peekToken.Type == lexer.COMMA {
			p.nextToken()
		}
	}

	if p.peekToken.Type != lexer.RBRACE {
		return nil
	}
	p.nextToken()

	return hash
}

func (p *Parser) parseArrayLiteral() ast.Expression {
	al := &ast.ArrayLiteral{Token: p.curToken}
	al.Elements = p.parseExpressionList(lexer.RBRACKET)
	return al
}

func (p *Parser) parseExpressionList(end lexer.TokenType) []ast.Expression {
	list := []ast.Expression{}
	if p.peekToken.Type == end {
		p.nextToken()
		return list
	}
	p.nextToken()
	list = append(list, p.parseExpression(LOWEST))
	for p.peekToken.Type == lexer.COMMA {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpression(LOWEST))
	}
	if p.peekToken.Type != end {
		return nil
	}
	p.nextToken()
	return list
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
	args = append(args, p.parseExpression(LOWEST))

	for p.peekToken.Type == lexer.COMMA {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}

	if p.peekToken.Type != lexer.RPAREN {
		return nil
	}
	p.nextToken()

	return args
}
