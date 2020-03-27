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
		{"!5", "!", 6},
		{"-15", "-", 16},
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

/*
// TEST CASE FOR INFIX PARSING
func TestParsingInfixOperator(t *testing.T) {
	tests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 6", 5, "+", 6},
		{"5 - 6", 5, "-", 6},
		{"5 * 6", 5, "*", 6},
		{"5 / 6", 5, "/", 6},
		{"5 == 6", 5, "==", 6},
		{"5 != 6", 5, "!=", 6},
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

}*/
func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ := il.(*ast.IntegerLiteral)
	if integ.Value != value {
		t.Fatalf("ERROR value does not match")
		return false
	}
	return true
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	if exp == nil {
		t.Errorf("ERROR the expression passed was nil")
		return false
	}

	ident := exp.(*ast.Identifier)
	if ident.Value != value {
		t.Errorf("ERROR ident.value = %s but value is %s \n", ident.Value, value)
		return false
	}

	lit := ident.TokenLiteral()
	if lit != value {
		t.Errorf("ERROR token literal = %s and value = %s does not match", lit, value)
		return false
	}
	return true
}

func testLiteralExpression(
	t *testing.T,
	exp ast.Expression,
	expected interface{},
) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case string:
		return testIdentifier(t, exp, v)
	case bool:
		return testBooleanLiteral(t, exp, v)
	}
	t.Errorf("ERROR the type of the expected value passed could not be handled got=%T", exp)
	return false
}

func testInfixExpression(t *testing.T,
	exp ast.Expression,
	left interface{},
	operator string,
	right interface{}) bool {

	if exp == nil {
		t.Error("ERROR got nil expression")
		return false
	}

	opExp := exp.(*ast.InfixExpression)

	if !(testLiteralExpression(t, opExp.Left, left)) {
		return false
	}

	if opExp.Operator != operator {
		t.Errorf("ERROR operator type not matched expected=%s got=%s", operator, opExp.Operator)
		return false
	}

	if !(testLiteralExpression(t, opExp.Right, right)) {
		return false
	}
	return true
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) bool {
	if exp == nil {
		t.Error("ERROR the expression was nil")
		return false
	}

	bo := exp.(*ast.Boolean)

	if bo.Value != value {
		t.Errorf("ERROR bo.Value=%v got=%v", bo.Value, value)
		return false
	}
	return true
}

func TestParsingInfixExpression(t *testing.T) {
	tests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{"true == true", true, "==", true},
		{"true != true", true, "!=", true},
		{"true == false", true, "==", false},
		{"false == true", false, "==", true},
		{"false == false", false, "==", false},
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
		if !testLiteralExpression(t, exp.Left, tt.leftValue) {
			return
		}
	}
}
