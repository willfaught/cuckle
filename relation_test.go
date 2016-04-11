package cuckle

import "testing"

func TestNewRelation(t *testing.T) {
	for _, test := range []struct {
		o    Operator
		l, r Term
		e    Relation
	}{
		{"", "", "", "  "},
		{"<", "1", "2", "1 < 2"},
	} {
		t.Log("Test:", test)

		if a := NewRelation(test.o, test.l, test.r); a != test.e {
			t.Errorf("Actual %v, expected %v", a, test.e)
		}
	}
}
