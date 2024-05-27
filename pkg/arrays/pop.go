package arrays

/*
# Pop

the Pop method removes the last element from array then returns it

Returns error and an zeroed generic value if array is empty
*/
func (a *Array[T]) Pop() (T, error) {
	var deleted T

	if a.Len() == 0 {
		return deleted, ErrEmptyArray
	}

	deleted = a.elements[a.lastPos]
	a.elements = a.elements[:a.lastPos]
	a.lastPos--

	return deleted, nil
}
