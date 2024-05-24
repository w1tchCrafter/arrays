package arrays


/*
# Shift

 the Shift method Removes the first item from the array then returns it

 returns error and an zeroed generic value if array is empty
*/

func (a *Array[T]) Shift() (T, error) {
  var deleted T

  if a.Len() == 0 {
    return deleted, ErrEmptyArray
  }

  deleted = a.elements[0]
  a.elements = a.elements[1:]
  a.lastPos--

  return deleted, nil
}
