package cuckle

import (
	"reflect"
	"testing"
)

func TestOptionVars(t *testing.T) {
	for _, test := range []struct {
		a Option
		e option
	}{
		{OptionAllowFiltering, optionAllowFiltering},
		{OptionCalled, optionCalled},
		{OptionClusteringOrder, optionClusteringOrder},
		{OptionCompactStorage, optionCompactStorage},
		{OptionDistinct, optionDistinct},
		{OptionIfExists, optionIfExists},
		{OptionIfNotExists, optionIfNotExists},
		{OptionIndexKeys, optionIndexKeys},
		{OptionJSON, optionJSON},
		{OptionReplace, optionReplace},
	} {
		t.Log("Test:", test)

		if e := (Option{test.e: nil}); !reflect.DeepEqual(test.a, e) {
			t.Errorf("Actual %v, expected %v", test.a, e)
		}
	}
}

func TestOptionAliases(t *testing.T) {
	for _, test := range []map[Identifier]Identifier{nil, {}, {"a": "b"}} {
		t.Log("Test:", test)

		if a, e := OptionAliases(test), (Option{optionAliases: test}); !reflect.DeepEqual(a, e) {
			t.Errorf("Actual %v, expected %v", a, e)
		}
	}
}

func TestOptionConditions(t *testing.T) {
	for _, test := range [][]Relation{nil, {}, {"a"}} {
		t.Log("Test:", test)

		if a, e := OptionConditions(test...), (Option{optionConditions: test}); !reflect.DeepEqual(a, e) {
			t.Errorf("Actual %v, expected %v", a, e)
		}
	}
}

func TestOptionFinalFunc(t *testing.T) {
	for _, test := range []Identifier{"", "a"} {
		t.Log("Test:", test)

		if a, e := OptionFinalFunc(test), (Option{optionFinalFunc: test}); !reflect.DeepEqual(a, e) {
			t.Errorf("Actual %v, expected %v", a, e)
		}
	}
}

func TestOptionIndexName(t *testing.T) {
	for _, test := range []Identifier{"", "a"} {
		t.Log("Test:", test)

		if a, e := OptionIndexName(test), (Option{optionIndexName: test}); !reflect.DeepEqual(a, e) {
			t.Errorf("Actual %v, expected %v", a, e)
		}
	}
}

func TestOptionInitCond(t *testing.T) {
	for _, test := range []Term{"", "a"} {
		t.Log("Test:", test)

		if a, e := OptionInitCond(test), (Option{optionInitCond: test}); !reflect.DeepEqual(a, e) {
			t.Errorf("Actual %v, expected %v", a, e)
		}
	}
}

func TestOptionLimit(t *testing.T) {
	for _, test := range []int{0, 1} {
		t.Log("Test:", test)

		if a, e := OptionLimit(test), (Option{optionLimit: test}); !reflect.DeepEqual(a, e) {
			t.Errorf("Actual %v, expected %v", a, e)
		}
	}
}

func TestOptionOptions(t *testing.T) {
	for _, test := range []map[Term]Term{nil, {"a": "b"}} {
		t.Log("Test:", test)

		if a, e := OptionOptions(test), (Option{optionOptions: test}); !reflect.DeepEqual(a, e) {
			t.Errorf("Actual %v, expected %v", a, e)
		}
	}
}

func TestOptionOrder(t *testing.T) {
	for _, test := range []struct {
		i []Identifier
		o []Order
	}{
		{},
		{[]Identifier{}, []Order{}},
		{[]Identifier{"a"}, []Order{OrderAscending}},
	} {
		t.Log("Test:", test)

		if a, e := OptionOrder(test.i, test.o), (Option{optionOrderByColumns: test.i, optionOrderByDirections: test.o}); !reflect.DeepEqual(a, e) {
			t.Errorf("Actual %v, expected %v", a, e)
		}
	}
}

func TestOptionSelectors(t *testing.T) {
	for _, test := range [][]Selector{nil, {}, {"a"}} {
		t.Log("Test:", test)

		if a, e := OptionSelectors(test...), (Option{optionSelectors: test}); !reflect.DeepEqual(a, e) {
			t.Errorf("Actual %v, expected %v", a, e)
		}
	}
}

func TestOptionTTL(t *testing.T) {
	for _, test := range []int64{0, 1} {
		t.Log("Test:", test)

		if a, e := OptionTTL(test), (Option{optionTTL: test}); !reflect.DeepEqual(a, e) {
			t.Errorf("Actual %v, expected %v", a, e)
		}
	}
}

func TestOptionTimestamp(t *testing.T) {
	for _, test := range []int64{0, 1} {
		t.Log("Test:", test)

		if a, e := OptionTimestamp(test), (Option{optionTimestamp: test}); !reflect.DeepEqual(a, e) {
			t.Errorf("Actual %v, expected %v", a, e)
		}
	}
}

func TestOptionTriggerName(t *testing.T) {
	for _, test := range []Identifier{"", "a"} {
		t.Log("Test:", test)

		if a, e := OptionTriggerName(test), (Option{optionTriggerName: test}); !reflect.DeepEqual(a, e) {
			t.Errorf("Actual %v, expected %v", a, e)
		}
	}
}

func TestOptionUsing(t *testing.T) {
	for _, test := range []string{"", "a"} {
		t.Log("Test:", test)

		if a, e := OptionUsing(test), (Option{optionUsing: test}); !reflect.DeepEqual(a, e) {
			t.Errorf("Actual %v, expected %v", a, e)
		}
	}
}

func TestOptionWhere(t *testing.T) {
	for _, test := range [][]Relation{nil, {}, {"a"}} {
		t.Log("Test:", test)

		if a, e := OptionWhere(test...), (Option{optionWhere: test}); !reflect.DeepEqual(a, e) {
			t.Errorf("Actual %v, expected %v", a, e)
		}
	}
}

func TestOptionWith(t *testing.T) {
	for _, test := range []map[Identifier]Term{
		nil,
		{"": ""},
		{"a": "b"},
	} {
		t.Log("Test:", test)

		if a, e := OptionWith(test), (Option{optionWith: test}); !reflect.DeepEqual(a, e) {
			t.Errorf("Actual %v, expected %v", a, e)
		}
	}
}
