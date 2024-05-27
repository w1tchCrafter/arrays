package arrays_test

import(
  "testing"
  "github.com/w1tchCrafter/arrays/pkg/arrays"
)

func TestUnshift(t *testing.T) {
    arr := arrays.New(2, 3, 4, 5, 6)
    arr.Unshift(1)

    if element, _ := arr.At(0); element != 1 || arr.Len() != 6 {
        t.Fail()
    }

    arr.Unshift(-1, 0)

    if element1, _ := arr.At(0); element1 != -1 {
        t.Fail()
    }
    if element2, _ := arr.At(1); element2 != 0 {
        t.Fail()
    }
}

