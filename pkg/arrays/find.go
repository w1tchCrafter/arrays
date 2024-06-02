package arrays

/*
# Find

the Find method returns the value of the first array item that satisfies the provided test function

examples:

# Find the first number equal or greater than 3:

numbers := arrays.New(1, 2, 3, 4, 5, 6)

	three, _ := numbers.Find(func(_, i2 int) bool {
		return i2 >= 3
	})

fmt.Println(third) // should print '3', the first matching result

# this method returns error if item was not found

	_, err := numbers.Find(func(_, i2 int) bool {
		return i2 == 10 // 10 don't exists in the array
	})

fmt.Println(err) // prints "item was not found in array"
*/
func (a Array[T]) Find(f func(int, T) bool) (T, error) {
	for i, v := range a.elements {
		if f(i, v) {
			return v, nil
		}
	}

	return *new(T), ErrNotFound
}
