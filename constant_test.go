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

func TestConstantInteger(t *testing.T) {
	for _, test := range []struct {
		i int64
		c Constant
	}{
		{0, "0"},
		{1, "1"},
	} {
		if a := ConstantInteger(test.i); a != test.c {
			t.Errorf("Actual constant %v, expected %v", a, test.c)
		}
	}
}

func TestConstantFloat(t *testing.T) {
	for _, test := range []struct {
		f float64
		c Constant
	}{
		{0, "0"},
		{1, "1"},
	} {
		if a := ConstantFloat(test.f); a != test.c {
			t.Errorf("Actual constant %v, expected %v", a, test.c)
		}
	}
}
