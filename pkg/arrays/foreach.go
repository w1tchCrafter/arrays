package arrays

/*
# ForEach

 execute a callback function en each array element

 Example:

	names := arrays.New("john", "doe")
	names.ForEach(func(s string, i int) {
		fmt.Printf("index: %d, value: %s\n", i, s)
	})
*/

func (a Array[T]) ForEach(f func(T, int)) {
	for i, v := range a.elements {
		f(v, i)
	}
}
