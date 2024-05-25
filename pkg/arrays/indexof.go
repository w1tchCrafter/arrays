package arrays

import "fmt"

/*
# IndexOf

The indexOf method searches an array for an element value and returns its position.

returns -1 if the item is not found.

If the item is present more than once, it returns the position of the first occurrence.
*/
func (a Array[T]) IndexOf(t T) int {
	result := -1

	for i, v := range a.elements {
		if fmt.Sprint(v) == fmt.Sprint(t) {
			result = i
		}
	}

	return result
}
