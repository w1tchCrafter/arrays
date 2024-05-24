package arrays_test

import (
  "testing"
  "arrays/pkg/arrays"
)

func TestToSlice(t *testing.T) {
	arr := arrays.New("John", "Doe")
	slc, _ := arr.ToSlice(arrays.FULL_COPY)

	if len(slc) != 2 || slc[0] != "John" || slc[1] != "Doe" {
		t.Fail()
	}

	arr.Push("foo", "bar")
	slc2, _ := arr.ToSlice(arrays.USE_INDEX, 0, 1)

	if len(slc2) != 2 || slc2[0] != "John" || slc2[1] != "Doe" {
		t.Log(slc)
		t.Fail()
	}

	_, err := arr.ToSlice(arrays.USE_INDEX, 1, 1, 1)
	if err == nil {
		t.Fail()
	}

	_, err = arr.ToSlice(arrays.USE_INDEX, -1, 1)
	if err == nil {
		t.Fail()
	}

	_, err = arr.ToSlice(arrays.USE_INDEX, 0, 10)
	if err == nil {
		t.Fail()
	}

	_, err = arr.ToSlice(arrays.USE_INDEX)
	if err == nil {
		t.Fail()
	}

	slc3, _ := arr.ToSlice(arrays.USE_INDEX, 3, 3)
	if slc3[0] != "bar" {
		t.Fail()
	}

	slc4, err := arr.ToSlice(arrays.ARR_END, 1)
	if len(slc4) != 3 || err != nil {
		t.Fail()
	}
}
