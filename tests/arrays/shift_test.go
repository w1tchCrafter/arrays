package arrays_test

import (
  "testing"
  "arrays/pkg/arrays"
)

func TestShift(t *testing.T) {
	arr := arrays.New("john", "doe")
	first, err := arr.Shift()

	if first != "john" || err != nil || arr.Len() != 1 {
		t.Fail()
	}
}
