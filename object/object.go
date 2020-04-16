package object

import "fmt"

type ObjectType string

// Object is an interface type that will be used to represent everything from AT
type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
)

/*Integer object section starts here*/

// Integer struct is object type for integers
type Integer struct {
	Value int64 // actual value of the integer
}

// Inspect returns the literal value in string
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Type returns the type name
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

/*Boolean object section starts here*/

// Boolean struct for boolean values
type Boolean struct {
	Value bool // The actual boolean literal
}

// Type returns boolean type
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

// Inspect returns the boolean literal as string
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

/* NULL struct type implementation starts here */

// Null struct
type Null struct{}

// Type returns NULL_OBJ here
func (n *Null) Type() ObjectType { return NULL_OBJ }

// Inspect returns null string here
func (n *Null) Inspect() string { return "null" }

// ReturnValue struct to handle return statements, Value is a regular object type
type ReturnValue struct {
	Value Object
}

// Type gets the type, duh ..
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

// Inspect returns the literal value of object
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

// Error is the type for handling error
type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR :" + e.Message }

// Strings section
type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }
