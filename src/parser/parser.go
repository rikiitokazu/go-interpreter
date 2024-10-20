package parser

import (
	"fmt"

	"github.com/rikiitokazu/go-interpreter/src/ast"
	"github.com/rikiitokazu/go-interpreter/src/lexer"
	"github.com/rikiitokazu/go-interpreter/src/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	errors    []string
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l,
		errors: []string{}}

	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{} // root node
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return nil
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// skippping expressions until we encounter semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt

}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token tob e %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
