package arrays

func (a Array[T]) Every(cb func(item T, index int) bool) bool {
	for i, v := range a.elements {
		if !cb(v, i) {
			return false
		}

	}

	return true
}
