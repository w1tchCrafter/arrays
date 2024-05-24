package arrays

/* 
# Unshift

 works the same way as the Push method, but ushes an item to the beggining of the array
*/

func (a *Array[T]) Unshift(items ...T){
  newElements := make([]T, len(a.elements) + len(items))

  copy(newElements, items)

  copy(newElements[len(items):], a.elements)

  a.elements = newElements
  a.lastPos = len(newElements) - 1
}
