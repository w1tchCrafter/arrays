package arrays_test

import (
	"arrays/pkg/arrays"
	"fmt"
  "testing"
)

func TestForEach(t *testing.T) {
  arr := arrays.New("this", "is", "a", "test")

  arr.ForEach(func(i int, s string) {
    fmt.Println(i, s)
  })
}
