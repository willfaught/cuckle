package cuckle

import "testing"

func TestNewVariable(t *testing.T) {
	for _, test := range []struct {
		s string
		e Variable
	}{
		{"", ":"},
		{"a", ":a"},
	} {
		t.Log("Test:", test)

		if a := NewVariable(test.s); a != test.e {
			t.Errorf("Actual %v, expected %v", a, test.e)
		}
	}
}
