package arrays

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrEmptyArray    = errors.New("error, the array element is empty")
	ErrIndexOutRange = errors.New("error, index out of range")
	ErrNotFound      = errors.New("item was not found in array")
	ErrWrongUsage    = errors.New("no index were passed to 'ToSlice' method")
)

const (
	FULL_COPY Options = iota
	USE_INDEX
	ARR_END
)

type Options int8
type Array[T any] []T

// Creates a new generic Array
//
// To create an empty array, the type must be passed
//
//	arr := arrays.New[string]()
//
// The type can be ommited if the array is initialized
//
//	arr := arrays.New("John", "Doe")
//
// You can also initialize the new Array with an existing slice
//
//	slc := []int{1, 2, 3, 4, 5, 6, 7}
//	arr := arrays.New(slc...)
func New[T any](items ...T) Array[T] {
	array := Array[T]{}

	for _, v := range items {
		array = append(array, v)
	}

	return array
}

func (a Array[T]) ForEach(f func(int, T)) {
	for i, v := range a {
		f(i, v)
	}
}

// Returns the array lenght
func (a Array[T]) Len() int {
	return len(a)
}

// Join returns a string made of the Array elements separated by the paramenter sep
func (a Array[T]) Join(sep string) string {
	strResult := make([]string, 0)

	for _, v := range a {
		strResult = append(strResult, fmt.Sprint(v))
	}

	return strings.Join(strResult, sep)
}

// At returns the element at the given index
//
// Returns error if index don't exists
func (a Array[T]) At(index int) (T, error) {
	var item T

	if index > a.Len()-1 {
		return item, ErrIndexOutRange
	}

	item = a[index]
	return item, nil
}

// The Pop method removes the last element then returns it
//
// Returns error and an zeroed generic value if array is empty
func (a *Array[T]) Pop() (T, error) {
	var deleted T
	tmp := New[T]()

	if a.Len() == 0 {
		return deleted, ErrEmptyArray
	}

	for i, v := range *a {
		if i == a.Len()-1 {
			deleted = v
			break
		}

		tmp = append(tmp, v)
	}

	*a = tmp
	return deleted, nil
}

// Push method adds one or more new items to the end of the array
func (a *Array[T]) Push(items ...T) {
	*a = append(*a, items...)
}

// The Shift method Removes the first item then returns it
//
// Returns error and an zeroed generic value if array is empty
func (a *Array[T]) Shift() (T, error) {
	var deleted T
	tmp := New[T]()

	if a.Len() == 0 {
		return deleted, ErrEmptyArray
	}

	for i, v := range *a {
		if i == 0 {
			deleted = v
			continue
		}

		tmp = append(tmp, v)
	}

	*a = tmp
	return deleted, nil

}

// Pushes an item to the beggining of the array
func (a *Array[T]) Unshift(item T) {
	tmp := Array[T]{}
	tmp = append(tmp, *a...)

	*a = Array[T]{item}
	*a = append(*a, tmp...)
}

// # ToSlice method returns a copy of the array as a slice
//
//	This method can return the entire array as a slice or just a portion of it
//
// Return the entire array:
//
//	arr := arrays.New("john", "doe")
//	slc, err := arr.ToSlice(FULL_COPY)
//
//	if err != nil {
//		// handle error
//	}
//
// Return from a specific index:
//
//	arr := arrays.New(1, 2, 3, 4, 5, 6, 7, 8)
//	slc, _ := arr.ToSlice(USE_INDEX, 2, 5)
//
//	fmt.Println(slc) // should print [3, 4, 5, 6]
//
// Use the ARR_END constant:
//
// ARR_END is a constant used to pass the last item in an array without the risk of going out of range
//
//	arr := arrays.New("wake", "eat", "sleep", "repeat")
//	slc, _ := arr.ToSlice(ARR_END, 2) // 2 here means the starting index, the second one do not need to be passed
//	fmt.Println(slc) // should print ["sleep", "repeat"]
func (a Array[T]) ToSlice(opt Options, indexes ...int) ([]T, error) {
	slc := make([]T, 0)
	l := a.Len() - 1

	switch opt {
	case FULL_COPY:
		slc = append(slc, a...)

	case USE_INDEX:
		// fails if more than 2 index values are passed or in case of wrong index values
		if len(indexes) != 2 {
			return slc, ErrWrongUsage
		} else if indexes[0] > l || indexes[1] > l || indexes[0] < 0 || indexes[1] < 0 {
			return slc, ErrIndexOutRange
		}

		for i := indexes[0]; i <= indexes[1]; i++ {
			slc = append(slc, a[i])
		}

	case ARR_END:

		if len(indexes) < 1 {
			return slc, ErrWrongUsage
		} else if indexes[0] > l {
			return slc, ErrIndexOutRange
		}

		for i := indexes[0]; i <= l; i++ {
			slc = append(slc, a[i])
		}
	}

	return slc, nil
}

// The Find method returns the value of the first array item that satisfies the provided test function
//
// returns error if item was not found
func (a Array[T]) Find(f func(int, T) bool) (T, error) {
	for i, v := range a {
		if f(i, v) {
			return v, nil
		}
	}

	return *new(T), ErrNotFound
}

/*

Iteration functions goes after here

*/

func Map[T, E any](a Array[T], f func(int, T) E) []E {
	result := make([]E, 0)

	for i, v := range a {
		result = append(result, f(i, v))
	}

	return result
}

func (a Array[T]) Filter(f func(int, T) bool) []T {
	result := make([]T, 0)

	for i, v := range a {
		if f(i, v) {
			result = append(result, v)
		}
	}

	return result
}
