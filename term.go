package cuckle

import (
	"fmt"
	"sort"
	"strings"
)

// Term is a value.
type Term string

// TermConstant returns a Term for a Constant.
func TermConstant(c Constant) Term {
	return Term(c)
}

// TermFunc returns a Term for calling a function with arguments.
func TermFunc(function Identifier, arguments ...Term) Term {
	var ss []string

	for _, a := range arguments {
		ss = append(ss, string(a))
	}

	return Term(fmt.Sprintf("%v(%v)", function, strings.Join(ss, ", ")))
}

// TermIdentifier returns a Term for an Identifier.
func TermIdentifier(i Identifier) Term {
	return Term(fmt.Sprint(i))
}

// TermIndex returns a Term for indexing a column.
func TermIndex(i Identifier, t Term) Term {
	return Term(fmt.Sprintf("%v[%v]", i, t))
}

// TermList returns a Term for a list.
func TermList(t ...Term) Term {
	var ss []string

	for _, t := range t {
		ss = append(ss, string(t))
	}

	return Term(fmt.Sprintf("[%v]", strings.Join(ss, ", ")))
}

// TermMap returns a Term for a map.
func TermMap(m map[Term]Term) Term {
	var ss []string
	var ks = make([]string, 0, len(m))

	for k := range m {
		ks = append(ks, string(k))
	}

	sort.Strings(ks)

	for _, k := range ks {
		ss = append(ss, fmt.Sprintf("%v: %v", k, m[Term(k)]))
	}

	return Term(fmt.Sprintf("{%v}", strings.Join(ss, ", ")))
}

// TermRelation returns a Term for a Relation.
func TermRelation(r Relation) Term {
	return Term(r)
}

// TermSet returns a Term for a set.
func TermSet(t ...Term) Term {
	var ss []string

	for _, t := range t {
		ss = append(ss, string(t))
	}

	return Term(fmt.Sprintf("{%v}", strings.Join(ss, ", ")))
}

// TermTuple returns a Term for a tuple.
func TermTuple(t ...Term) Term {
	var ss []string

	for _, t := range t {
		ss = append(ss, string(t))
	}

	return Term(fmt.Sprintf("(%v)", strings.Join(ss, ", ")))
}
