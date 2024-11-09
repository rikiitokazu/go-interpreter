package parser

import (
	"testing"

	"github.com/rikiitokazu/go-interpreter/src/ast"
	"github.com/rikiitokazu/go-interpreter/src/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got =%s", name, letStmt.Name)
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("praser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

// func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
// 	ident, ok := exp.(*ast.Identifier)
// 	if !ok {
// 		t.Errorf("exp not *ast.Identifier. got=%T", exp)
// 		return false
// 	}
// 	if ident.Value != value {
// 		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
// 		return false
// 	}
// 	if ident.TokenLiteral() != value {
// 		t.Errorf("ident.TokenLiteral not %s. got=%s", value,
// 			ident.TokenLiteral())
// 		return false
// 	}
// 	return true
// }

// func testLiteralExpression(
// 	t *testing.T,
// 	exp ast.Expression,
// 	expected interface{},
// ) bool {
// 	switch v := expected.(type) {
// 	case int:
// 		return testIntegerLiteral(t, exp, int64(v))
// 	case int64:
// 		return testIntegerLiteral(t, exp, v)
// 	case string:
// 		return testIdentifier(t, exp, v)
// 	}
// 	t.Errorf("type of exp not handled. got=%T", exp)
// 	return false
// }

// func testInfixExpression(t *testing.T, exp ast.Expression, left interface{},
// 	operator string, right interface{}) bool {
// 	opExp, ok := exp.(*ast.InfixExpression)
// 	if !ok {
// 		t.Errorf("exp is not ast.OperatorExpression. got=%T(%s)", exp, exp)
// 		return false
// 	}
// 	if !testLiteralExpression(t, opExp.Left, left) {
// 		return false
// 	}
// 	if opExp.Operator != operator {
// 		t.Errorf("exp.Operator is not '%s'. got=%q", operator, opExp.Operator)
// 		return false
// 	}
// 	if !testLiteralExpression(t, opExp.Right, right) {
// 		return false
// 	}
// 	return true
// }
