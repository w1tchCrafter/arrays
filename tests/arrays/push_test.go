package arrays_test

import(
  "testing"
  "reflect"
  "github.com/w1tchCrafter/arrays/pkg/arrays"
)

func TestPush(t *testing.T) {
    arr := arrays.New[string]()

    arr.Push("witch")
    arr.Push("crafter")

    slice, _ := arr.ToSlice(arrays.FULL_COPY)
    expectedSlice := []string{"witch", "crafter"}

    if !reflect.DeepEqual(slice, expectedSlice) {
        t.Errorf("TestPush: Unexpected slice after initial Push. Expected: %v, got: %v", expectedSlice, slice)
    }

    arr.Push([]string{"john", "doe"}...)

    slice, _ = arr.ToSlice(arrays.FULL_COPY)
    expectedSlice = []string{"witch", "crafter", "john", "doe"}

    if !reflect.DeepEqual(slice, expectedSlice) {
        t.Errorf("TestPush: Unexpected slice after second Push. Expected: %v, got: %v", expectedSlice, slice)
    }

    // Convert Array to slice before passing it to Push
    spongebobArray := arrays.New("spongebob")
    spongebobSlice, _ := spongebobArray.ToSlice(arrays.FULL_COPY)
    arr.Push(spongebobSlice...)

    slice, _ = arr.ToSlice(arrays.FULL_COPY)
    expectedSlice = []string{"witch", "crafter", "john", "doe", "spongebob"}

    if !reflect.DeepEqual(slice, expectedSlice) {
        t.Errorf("TestPush: Unexpected slice after third Push. Expected: %v, got: %v", expectedSlice, slice)
    }
}

