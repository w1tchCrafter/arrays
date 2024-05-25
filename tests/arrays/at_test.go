package arrays_test

import (
  "github.com/w1tchCrafter/arrays/pkg/arrays"
  "testing"
)

func TestAt(t *testing.T) {
  arr := arrays.New(1.2, 102.187, 45.0, 94.3)

	third, err := arr.At(2)
	if third != 45.0 || err != nil {
		t.Fail()
	}

	nan, err := arr.At(11)

	if nan != 0 || err == nil {
		t.Fail()
	}
}
