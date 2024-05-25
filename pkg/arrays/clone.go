package arrays

import "slices"

/*
# Clone

the Clone method returns a shallow copy of the array elements
*/
func (a Array[T]) Clone() []T {
	return slices.Clone(a.elements)
}
