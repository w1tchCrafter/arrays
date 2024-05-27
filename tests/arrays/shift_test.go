package arrays_test

import (
  "testing"
  "github.com/w1tchCrafter/arrays/pkg/arrays"
)

func TestShift(t *testing.T) {
	arr := arrays.New("john", "doe")
	first, err := arr.Shift()

	if first != "john" || err != nil || arr.Len() != 1 {
		t.Fail()
	}
}
