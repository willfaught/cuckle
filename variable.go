package cuckle

import "fmt"

// Variable is a CQL variable.
type Variable string

// VariableAnonymous is anonymous.
const VariableAnonymous Variable = "?"

// NewVariable returns a Variable for a name.
func NewVariable(name string) Variable {
	return Variable(fmt.Sprintf(":%v", name))
}
