package cuckle

import "testing"

func TestSelectorAlias(t *testing.T) {
	for _, test := range []struct {
		s Selector
		i Identifier
		e Selector
	}{
		{`""`, "", `"" as ""`},
		{`"a"`, "b", `"a" as "b"`},
	} {
		t.Log("Test:", test)

		if a := SelectorAlias(test.s, test.i); a != test.e {
			t.Errorf("Actual %v, expected %v", a, test.e)
		}
	}
}

func TestSelectorFunc(t *testing.T) {
	for _, test := range []struct {
		i Identifier
		a []Selector
		e Selector
	}{
		{"", nil, `""()`},
		{"", []Selector{}, `""()`},
		{"a", []Selector{}, `"a"()`},
		{"a", []Selector{`"b"`}, `"a"("b")`},
		{"a", []Selector{`"b"`, `"c"`}, `"a"("b", "c")`},
	} {
		t.Log("Test:", test)

		if a := SelectorFunc(test.i, test.a...); a != test.e {
			t.Errorf("Actual %v, expected %v", a, test.e)
		}
	}
}

func TestSelectorIdentifier(t *testing.T) {
	for _, test := range []struct {
		i Identifier
		e Selector
	}{
		{"", `""`},
		{"a", `"a"`},
	} {
		t.Log("Test:", test)

		if a := SelectorIdentifier(test.i); a != test.e {
			t.Errorf("Actual %v, expected %v", a, test.e)
		}
	}
}

func TestSelectorIndex(t *testing.T) {
	for _, test := range []struct {
		i Identifier
		t Term
		e Selector
	}{
		{"", "", `""[]`},
		{"a", "1", `"a"[1]`},
	} {
		t.Log("Test:", test)

		if a := SelectorIndex(test.i, test.t); a != test.e {
			t.Errorf("Actual %v, expected %v", a, test.e)
		}
	}
}

func TestSelectorTTL(t *testing.T) {
	for _, test := range []struct {
		i Identifier
		e Selector
	}{
		{"", `ttl("")`},
		{"a", `ttl("a")`},
	} {
		t.Log("Test:", test)

		if a := SelectorTTL(test.i); a != test.e {
			t.Errorf("Actual %v, expected %v", a, test.e)
		}
	}
}

func TestSelectorWriteTime(t *testing.T) {
	for _, test := range []struct {
		i Identifier
		e Selector
	}{
		{"", `writetime("")`},
		{"a", `writetime("a")`},
	} {
		t.Log("Test:", test)

		if a := SelectorWriteTime(test.i); a != test.e {
			t.Errorf("Actual %v, expected %v", a, test.e)
		}
	}
}
