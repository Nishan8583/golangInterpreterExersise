package parser

import (
	"log"
	"testing"

	"../ast"
	"../lexer"
)

func TestLetStatements(t *testing.T) {

	// the input that will be fed into the testing funciton
	input := `
	let x = 5;
	let y = 10;
	let foobar = 8338;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("Error parsing program returned nil")
	}

	log.Println(len(program.Statements))
	if len(program.Statements) != 3 {
		t.Fatalf("Error wrong number of statements %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	log.Println("***********The output of parsing the program", program)
	for index, test := range tests {
		stmt := program.Statements[index]
		if !testLetStatement(t, stmt, test.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("ERROR the tokn litereal is something different, someone different and i am afraid Got %s", s.TokenLiteral())
		return false
	}

	letSmt, ok := s.(*ast.LetStatement) // Asserting that this type is a pointer to *ast.LetStatement
	if !ok {
		t.Errorf("ERROR, the type does not seem to be *ast.Statement infact the type was %T", s)
		return false
	}

	// Checking if the identifier value is as expected
	if letSmt.Name.Value != name {
		t.Errorf("The expected identifier was not achieved Expected = %s	Got =%s", name, letSmt.Name.Value)
		return false
	}

	// checking if the token literal is same
	if letSmt.Name.TokenLiteral() != name {
		t.Errorf("Error got unexpected token literal expected = %s got = %s", name, letSmt.Name.TokenLiteral())
		return false
	}

	return true

}
