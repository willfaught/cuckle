package cuckle

import "testing"

func TestTypeList(t *testing.T) {
	for _, test := range []struct {
		t Type
		e Type
	}{
		{"", "list<>"},
		{"a", "list<a>"},
	} {
		t.Log("Test:", test)

		if a := TypeList(test.t); a != test.e {
			t.Errorf("Actual %v, expected %v", a, test.e)
		}
	}
}

func TestTypeMap(t *testing.T) {
	for _, test := range []struct {
		k, v Type
		e    Type
	}{
		{"", "", "map<, >"},
		{"a", "b", "map<a, b>"},
	} {
		t.Log("Test:", test)

		if a := TypeMap(test.k, test.v); a != test.e {
			t.Errorf("Actual %v, expected %v", a, test.e)
		}
	}
}

func TestTypeSet(t *testing.T) {
	for _, test := range []struct {
		t Type
		e Type
	}{
		{"", "set<>"},
		{"a", "set<a>"},
	} {
		t.Log("Test:", test)

		if a := TypeSet(test.t); a != test.e {
			t.Errorf("Actual %v, expected %v", a, test.e)
		}
	}
}

func TestTypeTuple(t *testing.T) {
	for _, test := range []struct {
		t []Type
		e Type
	}{
		{nil, "tuple<>"},
		{[]Type{"a"}, "tuple<a>"},
		{[]Type{"a", "b"}, "tuple<a, b>"},
	} {
		t.Log("Test:", test)

		if a := TypeTuple(test.t...); a != test.e {
			t.Errorf("Actual %v, expected %v", a, test.e)
		}
	}
}
