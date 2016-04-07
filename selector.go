package cuckle

import (
	"fmt"
	"strings"
)

// Selector selects the result row values.
type Selector string

// Simple selectors.
const (
	// SelectorAll selects all columns.
	SelectorAll Selector = "*"

	// SelectorCount selects the count of all result rows.
	SelectorCount Selector = "count(*)"
)

// SelectorAlias returns a Selector for aliasing s as alias.
func SelectorAlias(s Selector, alias Identifier) Selector {
	return Selector(fmt.Sprintf("%v as %v", s, alias))
}

// SelectorFunc returns a Selector for calling function with argumnets.
func SelectorFunc(function Identifier, arguments ...Selector) Selector {
	var ss []string

	for _, a := range arguments {
		ss = append(ss, string(a))
	}

	return Selector(fmt.Sprintf("%v(%v)", function, strings.Join(ss, ", ")))
}

// SelectorIdentifier returns a Selector for a column.
func SelectorIdentifier(column Identifier) Selector {
	return Selector(fmt.Sprint(column))
}

// SelectorIndex returns a Selector for indexing a column.
func SelectorIndex(column Identifier, index Term) Selector {
	return Selector(fmt.Sprintf("%v[%v]", column, index))
}

// SelectorTTL returns a Selector for the time-to-live of a column.
func SelectorTTL(i Identifier) Selector {
	return Selector(fmt.Sprintf("ttl(%v)", i))
}

// SelectorWriteTime returns a Selector for the write time of a column.
func SelectorWriteTime(i Identifier) Selector {
	return Selector(fmt.Sprintf("writetime(%v)", i))
}
