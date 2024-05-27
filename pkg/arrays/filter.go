package arrays

/*
# Filter

	similar to the Find method, but returns a slice of containing all matching resutls

	example:

# find all numbers greater than 18:

	numbers := arrays.New(23, 12, 43, 6, 98)
	grt := numbers.Filter(func(_, i2 int) bool {
		return i2 > 18
	})

	fmt.Println(grt) // [23, 43, 98]
*/
func (a Array[T]) Filter(f func(int, T) bool) []T {
	result := make([]T, 0)

	for i, v := range a.elements {
		if f(i, v) {
			result = append(result, v)
		}
	}

	return result
}
