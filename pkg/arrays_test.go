package arrays

import (
	"fmt"
	"testing"
)

func TestForEach(t *testing.T) {
	arr := New("this", "is", "a", "test")

	arr.ForEach(func(i int, s string) {
		fmt.Println(i, s)
	})
}

func TestMap(t *testing.T) {
	arr := New(2, 3, 5, 7)

	// should double the given element
	result := Map(arr, func(_, n int) int {
		return n * 2
	})

	if result[0] != 4 || result[1] != 6 || result[2] != 10 || result[3] != 14 {
		t.Log(result)
		t.Fail()
	}
}

func TestFilter(t *testing.T) {
	arr := New([]int{23, 12, 17, 123, 18}...)

	// result should have only numbers greater than or equal to 18
	result := arr.Filter(func(_, n int) bool {
		return n >= 18
	})

	for _, v := range result {
		if v < 18 {
			t.Fail()
		}
	}
}

func TestPush(t *testing.T) {
	arr := New[string]()

	arr.Push("witch")
	arr.Push("crafter")

	if arr[0] != "witch" || arr[1] != "crafter" {
		t.Fail()
	}
}
