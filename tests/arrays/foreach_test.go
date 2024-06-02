package arrays_test

import (
	"fmt"
	"testing"

	"github.com/w1tchCrafter/arrays/pkg/arrays"
)

func TestForEach(t *testing.T) {
	arr := arrays.New("this", "is", "a", "test")

	arr.ForEach(func(s string, i int) {
		fmt.Println(i, s)
	})
}
