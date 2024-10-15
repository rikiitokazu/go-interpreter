package lexer

// lexer takes source code as input and outputs the tokens that represents src code
// will output next token it recognizes, no need for buffer or save toekns since there will
// only be one method called NextToken()

import (
	"testing"

	"github.com/rikiitokazu/go-interpreter/src/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input(points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination we are using ASCII
}

// it is necessary to "peek" further into the input, hence why we have readPosition
// as it always points to the next positiion, whereas 'position' corresponds with ch
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// ASCII CODE for "NUL" char and signifies either "we havent read anything"
		// or end of file
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1

}

func (l *Lexer) NextToken() token.Token {
	// TODO
}
func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - ltieral wrong. expected=%q, got %q", i, tt.expectedLiteral, tok.Literal)
		}

	}
}
