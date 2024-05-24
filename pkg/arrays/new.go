package arrays

/* 
# New
	
  Creates a new generic Array
  
  To create an empty array, the type must be passed:

	arr := arrays.New[string]()

  The type can be ommited if the array is initialized:

  arr := arrays.New("John", "Doe")

  You can also initialize the new Array with an existing slice:

	slc := []int{1, 2, 3, 4, 5, 6, 7}
	arr := arrays.New(slc...)
*/

func New[T any](items  ...T) Array[T] {
  if len(items) > 0 {
    return Array[T]{elements: items, lastPos: len(items) - 1}
  }
  
  return Array[T]{elements: items, lastPos: -1}
}
