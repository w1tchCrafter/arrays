package arrays

/* 
# ForEach

 execute a callback function en each array element

 Example:

	names := arrays.New("john", "doe")
	names.ForEach(func(i int, s string) {
		fmt.Printf("index: %d, value: %s\n", i, s)
	})
*/

func (a Array[T]) ForEach(f func(int, T)){
  for i, v := range a.elements {
    f(i, v)
  }
}
