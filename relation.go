package cuckle

import "fmt"

// Relation is an infix operation with an operator and two operands.
type Relation string

// NewRelation returns a new Relation for left, o, and right.
func NewRelation(left Term, o Operator, right Term) Relation {
	return Relation(fmt.Sprintf("%v %v %v", left, o, right))
}
