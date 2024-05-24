package arrays

/*
# At
The At method returns the element at the given index

Returns error if index don't exists
*/

func (a Array[T]) At(index int) (T, error) {
  var item T
  
  if index < 0 || index >= a.Len() {
    return item, ErrNotFound
  }
  
  item = a.elements[index]

  return item, nil
}

