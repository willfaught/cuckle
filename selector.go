package cuckle

import (
	"fmt"
	"strings"
)

type Selector string

const (
	SelectorAll   Selector = "*"
	SelectorCount Selector = "count(*)"
)

func SelectorAlias(s Selector, alias Identifier) Selector {
	return Selector(fmt.Sprintf("%v as %v", s, alias))
}

func SelectorFunc(function Identifier, arguments ...Selector) Selector {
	var ss []string

	for _, a := range arguments {
		ss = append(ss, string(a))
	}

	return Selector(fmt.Sprintf("%v(%v)", function, strings.Join(ss, ", ")))
}

func SelectorIdentifier(i Identifier) Selector {
	return Selector(fmt.Sprint(i))
}

func SelectorIndex(i Identifier, t Term) Selector {
	return Selector(fmt.Sprintf("%v[%v]", i, t))
}

func SelectorTTL(i Identifier) Selector {
	return Selector(fmt.Sprintf("ttl(%v)", i))
}

func SelectorWriteTime(i Identifier) Selector {
	return Selector(fmt.Sprintf("writetime(%v)", i))
}
