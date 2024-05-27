package arrays

/*
# ToSlice method returns a copy of the array as a slice

# This method can return the entire array as a slice or just a portion of it

# Return the entire array:

	arr := arrays.New("john", "doe")
	slc, err := arr.ToSlice(FULL_COPY)

	if err != nil {
			handle error
	}

# Return from a specific index:

	arr := arrays.New(1, 2, 3, 4, 5, 6, 7, 8)
	slc, _ := arr.ToSlice(USE_INDEX, 2, 5)

	mt.Println(slc) // should print [3, 4, 5, 6]

# Use the ARR_END constant:

ARR_END is a constant used to pass the last item in an array without the risk of going out of range

	arr := arrays.New("wake", "eat", "sleep", "repeat")
	slc, _ := arr.ToSlice(ARR_END, 2) // 2 here means the starting index, the second one do not need to be passed
	fmt.Println(slc) // should print ["sleep", "repeat"]
*/
func (a Array[T]) ToSlice(opt Options, indexes ...int) ([]T, error) {
	switch opt {
	case FULL_COPY:
		return a.elements, nil
	case USE_INDEX:
		if len(indexes) != 2 {
			return nil, ErrWrongUsage
		}

		startIndex := indexes[0]
		endIndex := indexes[1]

		if startIndex < 0 || startIndex >= a.Len() || endIndex < 0 || endIndex >= a.Len() {
			return nil, ErrIndexOutRange
		}

		return a.elements[startIndex : endIndex+1], nil
	case ARR_END:
		if len(indexes) != 1 {
			return nil, ErrWrongUsage
		}

		endIndex := indexes[0]

		if endIndex < 0 || endIndex >= len(a.elements) {
			return nil, ErrIndexOutRange
		}

		return a.elements[endIndex:], nil
	default:
		return nil, ErrWrongUsage
	}
}
