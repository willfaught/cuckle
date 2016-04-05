package cuckle

import "testing"

func TestConstantBoolean(t *testing.T) {
	for _, test := range []struct {
		b bool
		c Constant
	}{
		{false, "false"},
		{true, "true"},
	} {
		if a := ConstantBoolean(test.b); a != test.c {
			t.Errorf("Actual constant %v, expected %v", a, test.c)
		}
	}
}
