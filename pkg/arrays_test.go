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

func TestLen(t *testing.T) {
	arr := New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	if arr.Len() != 10 {
		t.Fail()
	}
}

func TestJoin(t *testing.T) {
	arr := New("john", "doe", "is", "foo", "bar")

	if arr.Join(" ") != "john doe is foo bar" {
		t.Fail()
	}
}

func TestAt(t *testing.T) {
	arr := New(1.2, 102.187, 45.0, 94.3)

	third, err := arr.At(2)
	if third != 45.0 || err != nil {
		t.Fail()
	}

	nan, err := arr.At(11)

	if nan != 0 || err == nil {
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	arr := New("i", "use", "arch", "btw")
	item, err := arr.Pop()

	if item != "btw" || err != nil || arr.Len() != 3 {
		t.Fail()
	}

	arr2 := New[string]()
	empty, err := arr2.Pop()

	if empty != "" || err == nil {
		t.Fail()
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

func TestShift(t *testing.T) {
	arr := New("john", "doe")
	first, err := arr.Shift()

	if first != "john" || err != nil || arr.Len() != 1 {
		t.Fail()
	}
}

func TestUnshift(t *testing.T) {
	arr := New(2, 3, 4, 5, 6)
	arr.Unshift(1)

	if arr[0] != 1 || arr.Len() != 6 {
		t.Fail()
	}
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
