package arrays_test

import (
	"arrays/pkg/arrays"
	"testing"
)

func TestIndexOf(t *testing.T) {
	arr := arrays.New("a", "b", "c", "d", "e", "f", "g", "h", "i")
	indexD := arr.IndexOf("d") // should be 3
	indexG := arr.IndexOf("g") // should be 6

	if indexG != 6 || indexD != 3 {
		t.Fail()
	}

	ni := arr.IndexOf("z")

	if ni != -1 {
		t.Fail()
	}
}
