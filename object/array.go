package object

import (
	"fmt"
	"strings"
)

// Array is a struct to hold array object
type Array struct {
	Elements []Object
}

func (ar *Array) Type() ObjectType { return ARRAY_OBJ }
func (ar *Array) Inspect() string {

	elements := []string{}
	for _, e := range ar.Elements {
		elements = append(elements, e.Inspect())
	}

	return fmt.Sprintf(`[%s]`, strings.Join(elements, ", "))
}
