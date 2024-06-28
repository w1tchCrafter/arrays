package arrays_test

import (
	"testing"

	"github.com/w1tchCrafter/arrays/pkg/arrays"
)

func TestUnique(t *testing.T) {
	test := arrays.New(1, 2, 2, 1, 3, 3, 4, 5, 3)
	expected := arrays.New(1, 2, 3, 4, 5)

	test.Unique()

	for i := range 4 {
		f, _ := test.At(i)
		s, _ := expected.At(i)

		if f != s {
			t.Fail()
		}
	}
}
