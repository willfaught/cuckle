package cuckle

import (
	"fmt"
	"testing"
)

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

func TestConstantHex(t *testing.T) {
	for _, test := range []struct {
		b []byte
		c Constant
	}{
		{nil, ""},
		{[]byte{1}, "01"},
		{[]byte{1, 2}, "0102"},
	} {
		if a, e := ConstantHex(test.b), Constant(fmt.Sprintf("0x%v", test.c)); a != e {
			t.Errorf("Actual constant %v, expected %v", a, e)
		}
	}
}

func TestConstantString(t *testing.T) {
	for _, test := range []struct {
		s string
		c Constant
	}{
		{"", "''"},
		{"a", "'a'"},
		{"ab", "'ab'"},
	} {
		if a := ConstantString(test.s); a != test.c {
			t.Errorf("Actual constant %v, expected %v", a, test.c)
		}
	}
}

func TestConstantStringEscaped(t *testing.T) {
	for _, test := range []struct {
		s string
		c Constant
	}{
		{"", "$$$$"},
		{"a", "$$a$$"},
		{"ab", "$$ab$$"},
		{"a\nb", "$$a\nb$$"},
		{"'ab'", "$$'ab'$$"},
	} {
		if a := ConstantStringEscaped(test.s); a != test.c {
			t.Errorf("Actual constant %v, expected %v", a, test.c)
		}
	}
}

func TestConstantUUID(t *testing.T) {
	for _, test := range []struct {
		s string
		c Constant
	}{
		{"", ""},
		{"a", "a"},
	} {
		if a := ConstantUUID(test.s); a != test.c {
			t.Errorf("Actual constant %v, expected %v", a, test.c)
		}
	}
}
