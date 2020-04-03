package object

import "fmt"

type ObjectType string

// Object is an interface type that will be used to represent everything from AT
type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
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
