package arrays

/*
# Unique

The unique method removes duplicates from an array
*/
func (a *Array[T]) Unique() {
	uniq := New[T]()

	a.ForEach(func(t T, _ int) {
		if uniq.Contains(t) {
			return
		}

		uniq.Push(t)
	})

	*a = uniq
}
