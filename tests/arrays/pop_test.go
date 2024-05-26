package arrays_test

import(
  "github.com/w1tchCrafter/arrays/pkg/arrays"
  "testing"
)

func TestPop(t *testing.T) {
	arr := arrays.New("i", "use", "arch", "btw")
	item, err := arr.Pop()

	if item != "btw" || err != nil || arr.Len() != 3 {
		t.Fail()
	}

	arr2 := arrays.New[string]()
	empty, err := arr2.Pop()

	if empty != "" || err == nil {
		t.Fail()
	}
}
