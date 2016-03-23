package cuckle

import (
	"fmt"
	"strings"
)

type Term string

func TermConstant(c Constant) Term {
	return Term(c)
}

func TermFunc(function Identifier, arguments ...Term) Term {
	var ss []string

	for _, a := range arguments {
		ss = append(ss, string(a))
	}

	return Term(fmt.Sprintf("%v(%v)", function, strings.Join(ss, ", ")))
}

func TermIdentifier(i Identifier) Term {
	return Term(fmt.Sprint(i))
}

func TermIndex(i Identifier, t Term) Term {
	return Term(fmt.Sprintf("%v[%v]", i, t))
}

func TermList(t ...Term) Term {
	var ss []string

	for _, t := range t {
		ss = append(ss, string(t))
	}

	return Term(fmt.Sprintf("[%v]", strings.Join(ss, ", ")))
}

func TermMap(m map[Term]Term) Term {
	var ss []string

	for k, v := range m {
		ss = append(ss, fmt.Sprintf("%v: %v", k, v))
	}

	return Term(fmt.Sprintf("{%v}", strings.Join(ss, ", ")))
}

func TermOperation(left Term, o Operator, right Term) Term {
	return Term(fmt.Sprintf("%v %v %v", left, o, right))
}

func TermSet(t ...Term) Term {
	var ss []string

	for _, t := range t {
		ss = append(ss, string(t))
	}

	return Term(fmt.Sprintf("{%v}", strings.Join(ss, ", ")))
}

func TermTuple(t ...Term) Term {
	var ss []string

	for _, t := range t {
		ss = append(ss, string(t))
	}

	return Term(fmt.Sprintf("(%v)", strings.Join(ss, ", ")))
}
