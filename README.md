# Arrays - A package implementing js-like arrays in GO

Arrays brings types, methods, and functions to enhance Go slices, inspired by JavaScript arrays.

## Table of contents

- [Installation](#installation)
- [Usage](#usage)
  - [New](#new)
  - [ForEach](#foreach)
  - [Len](#len)
  - [Join](#join)
  - [At](#at)
  - [Pop](#pop)
  - [Push](#push)
  - [Shift](#shift)
  - [Unshift](#unshift)
  - [ToSlice](#toslice)
  - [Find](#find)
  - [Map](#map)
  - [Filter](#filter)

## Installation

This package uses generics, requiring Go `1.18` or later.

```sh
$ go get github.com/w1tchCrafter/arrays
```

# Usage

## New

Creates a new generic Array.

To create an empty array, the type must be passed:

```go

arr := arrays.New[string]()
```
The generic type can be omitted if the array is initialized:

```go

arr := arrays.New("John", "Doe")
```
You can also initialize the new Array with an existing slice:

```go

slc := []int{1, 2, 3, 4, 5, 6, 7}
arr := arrays.New(slc...)
```

## ForEach

Executes a callback function on each array element.

```go

func (a Array[T]) ForEach(f func(int, T))
```

Example:

```go

names := arrays.New("john", "doe")
names.ForEach(func(i int, s string) {
    fmt.Printf("index: %d, value: %s\n", i, s)
})
```

## Len

Returns the array length.

```go

func (a Array[T]) Len() int
```

Example:

```go

arr := arrays.New(1, 2, 3, 4)
length := arr.Len() // length: 4
```

## Join

Returns a string made of the Array elements separated by the parameter sep.

```go

func (a Array[T]) Join(sep string) string
```

Example:

```go

arr := arrays.New("I", "love", "pizza")
fmt.Println(arr.Join(" ")) // prints: "I love pizza"
```

## At

Returns the element at the given index.

```go

func (a Array[T]) At(index int) (T, error)
```

Example:

```go

arr := arrays.New(1, 2, 3, 4)
element, err := arr.At(2) // element: 3, err: nil
```

## Pop

Removes the last element from the array and returns it.

```go

func (a *Array[T]) Pop() (T, error)
```

Example:

```go

arr := arrays.New(1, 2, 3, 4)
last, err := arr.Pop() // last: 4, err: nil
```

## Push

Appends one or more new items to the end of the array.

```go

func (a *Array[T]) Push(items ...T)
```

Example:

```go

foods := arrays.New[string]()
foods.Push("pizza")
foods.Push("hamburger", "lasagna", "fries")
```

## Shift

Removes the first item from the array and returns it.

```go
func (a *Array[T]) Shift() (T, error)
```

Example:

```go

arr := arrays.New(1, 2, 3, 4)
first, err := arr.Shift() // first: 1, err: nil
```

## Unshift

Works the same way as the Push method, but pushes an item to the beginning of the array.

```go

func (a *Array[T]) Unshift(items ...T)
```

Example:

```go

arr := arrays.New(2, 3, 4)
arr.Unshift(1) // arr: [1, 2, 3, 4]
```

## ToSlice

Returns a copy of the array as a slice.

```go

func (a Array[T]) ToSlice(opt Options, indexes ...int) ([]T, error)
```

Example:

```go

arr := arrays.New("john", "doe")
slc, err := arr.ToSlice(FULL_COPY) // slc: ["john", "doe"]

arr := arrays.New(1, 2, 3, 4, 5, 6, 7, 8)
slc, _ := arr.ToSlice(USE_INDEX, 2, 5) // slc: [3, 4, 5, 6]
```

## Find

Returns the value of the first array item that satisfies the provided test function.

```go

func (a Array[T]) Find(f func(int, T) bool) (T, error)
```

Example:

```go

numbers := arrays.New(1, 2, 3, 4, 5, 6)
three, _ := numbers.Find(func(_, i2 int) bool {
    return i2 >= 3
}) // three: 3
```

## Map

Transforms a slice using a provided function.

```go

func Map[T, E any](a Array[T], f func(int, T) E) []E
```

Example:

```go

input := arrays.New(1, 2, 3, 4)
output := Map(input, func(_ int, x int) int { return x * 2 })
// output: [2, 4, 6, 8]
```

## Filter

Similar to the Find method, but returns a slice containing all matching results.

```go

func (a Array[T]) Filter(f func(int, T) bool) []T
```

Example:

```go

numbers := arrays.New(23, 12, 43, 6, 98)
grt := numbers.Filter(func(_, i2 int) bool {
    return i2 > 18
}) // grt: [23, 43, 98]
```

For further details, check the [source code](https://github.com/w1tchCrafter/arrays/blob/main/pkg/arrays.go).