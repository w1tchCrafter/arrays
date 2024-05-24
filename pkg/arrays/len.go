package arrays

/* 
  # Len

 Returns the array lenght
*/

func (a Array[T]) Len() int {
  return len(a.elements)
}
