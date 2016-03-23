package cuckle

import "fmt"

type Relation string

func NewRelation(left Term, o Operator, right Term) Relation {
	return Relation(fmt.Sprintf("%v %v %v", left, o, right))
}
