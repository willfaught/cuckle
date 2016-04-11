package cuckle

import "testing"

func TestTermConstant(t *testing.T) {
	for _, test := range []struct {
		c Constant
		e Term
	}{
		{"", ""},
		{"1", "1"},
	} {
		t.Log("Test:", test)

		if a := TermConstant(test.c); a != test.e {
			t.Errorf("Actual %v, expected %v", a, test.e)
		}
	}
}

func TestTermFunc(t *testing.T) {
	for _, test := range []struct {
		i Identifier
		a []Term
		e Term
	}{
		{"", nil, `""()`},
		{"a", []Term{"1", "2"}, `"a"(1, 2)`},
	} {
		t.Log("Test:", test)

		if a := TermFunc(test.i, test.a...); a != test.e {
			t.Errorf("Actual %v, expected %v", a, test.e)
		}
	}
}

func TestTermIdentifier(t *testing.T) {
	for _, test := range []struct {
		i Identifier
		e Term
	}{
		{"", `""`},
		{"a", `"a"`},
	} {
		t.Log("Test:", test)

		if a := TermIdentifier(test.i); a != test.e {
			t.Errorf("Actual %v, expected %v", a, test.e)
		}
	}
}

func TestTermIndex(t *testing.T) {
	for _, test := range []struct {
		i Identifier
		t Term
		e Term
	}{
		{"", "", `""[]`},
		{"a", "1", `"a"[1]`},
	} {
		t.Log("Test:", test)

		if a := TermIndex(test.i, test.t); a != test.e {
			t.Errorf("Actual %v, expected %v", a, test.e)
		}
	}
}

func TestTermList(t *testing.T) {
	for _, test := range []struct {
		t []Term
		e Term
	}{
		{nil, "[]"},
		{[]Term{"1"}, "[1]"},
		{[]Term{"1", "2"}, "[1, 2]"},
	} {
		t.Log("Test:", test)

		if a := TermList(test.t...); a != test.e {
			t.Errorf("Actual %v, expected %v", a, test.e)
		}
	}
}

func TestTermMap(t *testing.T) {
	for _, test := range []struct {
		m map[Term]Term
		e Term
	}{
		{nil, "{}"},
		{map[Term]Term{"1": "2"}, "{1: 2}"},
		{map[Term]Term{"1": "2", "3": "4"}, "{1: 2, 3: 4}"},
	} {
		t.Log("Test:", test)

		if a := TermMap(test.m); a != test.e {
			t.Errorf("Actual %v, expected %v", a, test.e)
		}
	}
}

func TestTermRelation(t *testing.T) {
	for _, test := range []struct {
		r Relation
		e Term
	}{
		{"", ""},
		{"a", "a"},
	} {
		t.Log("Test:", test)

		if a := TermRelation(test.r); a != test.e {
			t.Errorf("Actual %v, expected %v", a, test.e)
		}
	}
}

func TestTermSet(t *testing.T) {
	for _, test := range []struct {
		t []Term
		e Term
	}{
		{nil, "{}"},
		{[]Term{"1"}, "{1}"},
		{[]Term{"1", "2"}, "{1, 2}"},
	} {
		t.Log("Test:", test)

		if a := TermSet(test.t...); a != test.e {
			t.Errorf("Actual %v, expected %v", a, test.e)
		}
	}
}

func TestTermTuple(t *testing.T) {
	for _, test := range []struct {
		t []Term
		e Term
	}{
		{nil, "()"},
		{[]Term{"1"}, "(1)"},
		{[]Term{"1", "2"}, "(1, 2)"},
	} {
		t.Log("Test:", test)

		if a := TermTuple(test.t...); a != test.e {
			t.Errorf("Actual %v, expected %v", a, test.e)
		}
	}
}
