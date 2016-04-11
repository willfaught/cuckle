package cuckle

import "testing"

func TestNewRelation(t *testing.T) {
	for _, test := range []struct {
		l, r Term
		o    Operator
		e    Relation
	}{
		{"", "", "", "  "},
		{"1", "2", "<", "1 < 2"},
	} {
		t.Log("Test:", test)

		if a := NewRelation(test.l, test.o, test.r); a != test.e {
			t.Errorf("Actual %v, expected %v", a, test.e)
		}
	}
}
