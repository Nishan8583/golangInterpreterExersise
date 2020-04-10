package evaluator

import (
	"testing"

	"../lexer"
	"../object"
	"../parser"
)

/*
func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}

	for _, tt := range tests {
		evaluate := testEval(tt.input)
		testIntegerObject(t, evaluate, tt.expected)
	}

}*/

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	return Eval(program, env)

}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	if obj == nil {
		t.Fatal("FATAL, nil object got")
	}

	r, ok := obj.(*object.Integer)
	if !(ok) {
		t.Errorf("ERROR could not assert the interface value as integer got=%T ", obj)
		return false
	}

	if r.Value != expected {
		t.Errorf("ERROR the value does not match got=%d expected=%d", r.Value, expected)
		return false
	}

	t.Log("The integer values match")
	return true
}

/*
func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!false", true},
		{"!true", false},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
	}

	for _, tt := range tests {
		evaulated := testEval(tt.input)
		booleanObj := evaulated.(*object.Boolean)
		if booleanObj.Value != tt.expected {
			t.Errorf("ERROR value expected=%v got=%v", tt.expected, booleanObj.Value)
		}
	}
}
*/

func testLetEvaluation(t *testing.T) {
	tests := []struct {
		input  string
		output int64
	}{
		{"let a=5;a", 5},
		{"let a=5*5;a", 25},
		{"let a=5; let b=6; let c = a*b;c", 30},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.output)
	}
}
