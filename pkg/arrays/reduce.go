package arrays

import (
	"reflect"
)

/*
# Reduce

The Reduce function runs a function on each array element to produce (reduce it to) a single value.

# Note that this function takes 3 arguments:

  - a Array[T]: the array itself
  - cb func(accumulator E, item T, index int) E: the callback function to be executed
  - initial ...E: initial value to be used (optional)

if multiple values are passed to initial only the first one will be used

initial is required if the generic types T and E are different

# This example reduces an array of integers to a single value:

	arr := arrays.New(25, 18, 32, 20, 15)

	max := arrays.Reduce(arr, func(acm, item, _ int) int {
		return int(math.Max(float64(acm), float64(item)))
	}) // returns 110

# This example reduces the same array to a single string:

	arr := arrays.New(25, 18, 32, 20, 15)

	str := arrays.Reduce(arr, func(acm string, item, _ int) string {
		return fmt.Sprint(acm, item)
	}, "") // returns "2518322015"
*/
func Reduce[T, E any](a Array[T], cb func(accumulator E, item T, index int) E, initial ...E) E {
	var acm E

	if len(initial) != 0 {
		acm = initial[0]
	} else {
		first, err := a.At(0)

		if err != nil {
			return acm
		}

		acmType := reflect.TypeOf(acm)
		firstType := reflect.TypeOf(first)

		if acmType != firstType {
			return acm
		}
	}

	for i := range a.Len() {
		current, _ := a.At(i)
		acm = cb(acm, current, i)
	}

	return acm
}
