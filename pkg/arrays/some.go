package arrays

/*
# Some

The some method checks if any array elements pass a test provided as a callback function.

it returns true if the callback returns true for one of the array elements

in case the callback function does not return true or the array is empty, this method returns false
*/
func (a Array[T]) Some(cb func(item T, index int) bool) bool {
	for i, v := range a.elements {
		if cb(v, i) {
			return true
		}
	}

	return false
}
