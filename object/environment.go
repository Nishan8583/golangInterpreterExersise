package object

// Environment struct will hold the map object that sotres identifier and their values
type Environment struct {
	store map[string]Object // object could be anything, from simple things like integer, to boolean to functions itself
	outer *Environment      // pointer to the outer,global environment
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// NewEnvironment creates a new environment structure and returns it
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// Get is used to get value from the environment structure
func (e *Environment) Get(name string) (Object, bool) {

	// Now first try to ge the outer environment value
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set is used to store environment value to structure
func (e *Environment) Set(name string, val Object) Object {

	// store in local when inside funciton
	e.store[name] = val
	return val
}
