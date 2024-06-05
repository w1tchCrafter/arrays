package arrays

// unstable due to interface conversion
func Reduce[T, E any](a Array[T], cb func(accumulator E, item T, index int) E, initial ...E) E {
	var acm E
	count := 0

	if len(initial) != 0 {
		acm = initial[0]
	} else {
		first, err := a.At(0)

		if err != nil {
			return acm
		}

		acm = any(first).(E)
		count = 1
	}

	for i := count; i < a.Len(); i++ {
		current, _ := a.At(i)
		acm = cb(acm, current, count)
	}

	return acm
}
