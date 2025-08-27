package lox

// Environment is a data structure to store named bindings for a given Lox
// scope.
//
// Environments can be nested within each other. Each environment contains
// internally managed reference to the one enclosing it.
type Environment struct {
	outer  *Environment
	values map[string]any
}

// NewEnvironment creates a new environment.
func NewEnvironment() *Environment {
	return &Environment{
		outer:  nil,
		values: make(map[string]any),
	}
}

func newInnerEnvironment(outer *Environment) *Environment {
	return &Environment{
		outer:  outer,
		values: make(map[string]any),
	}
}

// Define creates a named binding.
func (env *Environment) Define(name string, value any) {
	env.values[name] = value
}

// Assign replaces the value of a named binding if it exists.
//
// This function will search through outer environments for the named binding
// and return a boolean value to indicate if the named binding exists.
func (env *Environment) Assign(name string, value any) bool {
	if _, ok := env.values[name]; ok {
		env.values[name] = value
		return true
	}

	if env.outer != nil {
		return env.outer.Assign(name, value)
	}

	return false
}

// Get retrieves the value of a named binding if it exists.
//
// This function will search through outer environments for the named binding
// and return a boolean value to indicate if the named binding exists.
func (env *Environment) Get(name string) (any, bool) {
	if value, ok := env.values[name]; ok {
		return value, ok
	}

	if env.outer != nil {
		return env.outer.Get(name)
	}

	return nil, false
}
