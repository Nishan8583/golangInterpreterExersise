package object

import (
	"fmt"
	"strings"

	"../ast"
)

// Function struct is a object to hold functions
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	return fmt.Sprintf(`fn(%s) {\n%s\n}`, strings.Join(params, ", "), f.Body.String())
}
