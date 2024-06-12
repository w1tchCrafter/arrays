package arrays

/*
# Every

The every method checks if all array elements pass a test provided as a callback function.

it returns true if the callback returns true for all of the array elements

in case the callback function does not return true for one item or the array is empty, this method returns false
*/
func (a Array[T]) Every(cb func(item T, index int) bool) bool {
	for i, v := range a.elements {
		if !cb(v, i) {
			return false
		}

	}

	return true
}
