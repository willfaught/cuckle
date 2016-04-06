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
