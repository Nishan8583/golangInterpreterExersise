package parser

import (
	"testing"

	"../ast"
	"../lexer"
)

/*
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

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if len(program.Statements) != 1 {
		t.Errorf("ERROR could not get any statements from programs %v", len(program.Statements))
	}

	t.Log("Here is the whole program statements", program.Statements)

}
*/
/*
TEST CASE FOR PREFIX PARSING
func TestParsingPrefixExpressions(t *testing.T) {

	// Creating some table testing
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5", "!", 5},
		{"-15", "-", 15},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)

		program := p.ParseProgram()

		if len(program.Statements) != 1 {
			t.Errorf("Wrong number of statements got=%T", len(program.Statements))
		}

		t.Logf("The type is %T", program.Statements[0])

		stmt := program.Statements[0].(*ast.ExpressionStatement)

		// If no expression was observed
		if stmt.Expression == nil {
			t.Fatalf("NIL Statments")
		}

		// Type cast it as PrefixExpression
		exp := stmt.Expression.(*ast.PrefixExpression)

		t.Log(exp)
		if exp.Operator != tt.operator {
			t.Fatalf("ERROR Operator type not matched Expected=%s got=%s ", tt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
			return
		}

	}

}
*/

// TEST CASE FOR INFIX PARSING
func TestParsingInfixOperator(t *testing.T) {
	tests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5+5", 5, "+", 5},
		{"5-5", 5, "-", 5},
		{"5*5", 5, "*", 5},
		{"5/5", 5, "/", 5},
		{"5==5", 5, "==", 5},
		{"5!=5", 5, "!=", 5},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		program := p.ParseProgram()

		if len(program.Statements) != 1 {
			t.Errorf("Wrong number of statements got=%T", len(program.Statements))
		}

		t.Logf("The type is %T", program.Statements[0])

		stmt := program.Statements[0].(*ast.ExpressionStatement)

		// If no expression was observed
		if stmt.Expression == nil {
			t.Fatalf("NIL Statments")
		}

		// Type cast it as PrefixExpression
		exp := stmt.Expression.(*ast.InfixExpression)

		t.Log(exp)
		if exp.Operator != tt.operator {
			t.Fatalf("ERROR Operator type not matched Expected=%s got=%s ", tt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
			return
		}

		if !testIntegerLiteral(t, exp.Right, tt.rightValue) {
			return
		}

	}

}
func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ := il.(*ast.IntegerLiteral)
	if integ.Value != value {
		t.Fatalf("ERROR value does not match")
		return false
	}
	return true
}
