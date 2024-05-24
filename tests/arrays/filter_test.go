package arrays_test

import(
  "arrays/pkg/arrays"
  "testing"
)

func TestFilter(t *testing.T) {
	arr := arrays.New([]int{23, 12, 17, 123, 18}...)

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
