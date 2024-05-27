package arrays

/*
# Map

the Map method applies a given function to each element of the array and returns a new slice
*/
func Map[T, E any](a Array[T], f func(int, T) E) []E {
	result := make([]E, 0)

	for i, v := range a.elements {
		result = append(result, f(i, v))
	}

	return result
}
