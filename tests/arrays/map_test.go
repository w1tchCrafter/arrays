package arrays_test

import (
	"github.com/w1tchCrafter/arrays/pkg/arrays"
	"testing"
)

func TestMap(t *testing.T) {
	arr := arrays.New(2, 3, 5, 7)

	// should double the given element
	result := arrays.Map(arr, func(_, n int) int {
		return n * 2
	})

	if result[0] != 4 || result[1] != 6 || result[2] != 10 || result[3] != 14 {
		t.Log(result)
		t.Fail()
	}
}
