package cuckle

import (
	"fmt"
	"testing"
)

func TestIdentifier(t *testing.T) {
	for _, test := range []struct {
		s string
		e string
	}{
		{"", `""`},
		{"a", `"a"`},
	} {
		t.Log("Test:", test)

		if a := fmt.Sprint(Identifier(test.s)); a != test.e {
			t.Errorf("Actual constant %v, expected %v", a, test.e)
		}
	}
}
