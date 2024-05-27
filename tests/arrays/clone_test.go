package arrays_test

import (
	"github.com/w1tchCrafter/arrays/pkg/arrays"
	"testing"
)

func TestCLone(t *testing.T) {
	arr := arrays.New(1, 2, 3, 4, 5)
	slc := arr.Clone()

	arr.Push(6, 7, 8)

	// slc should not be affected
	if len(slc) == arr.Len() {
		t.Fail()
	}
}
