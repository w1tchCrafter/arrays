package arrays_test

import (
	"strings"
	"testing"

	"github.com/w1tchCrafter/arrays/pkg/arrays"
)

func TestEvery(t *testing.T) {
	arr := arrays.New(1, 2, 3, 4, 5, 6, 7)
	gtrfive := arr.Every(func(item, _ int) bool {
		return item >= 8
	})

	if gtrfive {
		t.Fail()
	}

	arr2 := arrays.New("foo", "bar", "john", "doe")

	oCont := arr2.Every(func(item string, _ int) bool {
		return strings.Contains(item, "o")
	})

	if oCont {
		t.Fail()
	}

	gtrfive = arr.Every(func(item, _ int) bool {
		return item >= 0
	})

	if !gtrfive {
		t.Fail()
	}

}
