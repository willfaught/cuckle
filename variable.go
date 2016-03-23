package cuckle

import "fmt"

type Variable string

const VariableAnonymous Variable = "?"

func NewVariable(name string) Variable {
	return Variable(fmt.Sprintf(":%v", name))
}
