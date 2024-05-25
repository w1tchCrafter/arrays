package arrays

import "fmt"

func (a Array[T]) LastIndexOF(t T) int {
	result := -1

	for i := a.lastPos; i >= 0; i-- {
		if fmt.Sprint(a.elements[i]) == fmt.Sprint(t) {
			result = i
			break
		}
	}

	return result
}
