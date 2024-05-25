package arrays

/*
# Contains

The Contains method allow users to check if an item is present in the array
*/
func (a Array[T]) Contains(v T) bool {
	return a.IndexOf(v) != -1
}
