package arrays_test

import (
	"arrays/pkg/arrays"
	"testing"
)

func TestLastIndexOf(t *testing.T) {
	arr := arrays.New(1, 2, 3, 3, 2, 1)
	last3 := arr.LastIndexOF(3) // should be 3
	last1 := arr.LastIndexOF(1) // should be 5

	if last3 != 3 || last1 != 5 {
		t.Fail()
	}

	ni := arr.LastIndexOF(8)

	if ni != -1 {
		t.Fail()
	}
}
