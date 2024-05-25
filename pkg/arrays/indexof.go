package arrays

import "reflect"

/*
# IndexOf

The indexOf method searches an array for an element value and returns its position.

returns -1 if the item is not found.

If the item is present more than once, it returns the position of the first occurrence.
*/
func (a Array[T]) IndexOf(item T) int {
	result := -1

	for i, v := range a.elements {
		if reflect.DeepEqual(v, item) {
			return i
		}
	}

	return result
}
