package arrays_test

import (
	"arrays/pkg/arrays"
	"testing"
)

func TestFind(t *testing.T) {
  arr := arrays.New(12, 54, 2, 0, 322)

  found, err := arr.Find(func(_, i2 int) bool {
		return i2 > 100
	})

	if err != nil || found != 322 {
		t.Fail()
	}

	found, err = arr.Find(func(_, i2 int) bool {
		return i2 < 0
	})

	if found != 0 || err == nil {
		t.Fail()
	}
}
