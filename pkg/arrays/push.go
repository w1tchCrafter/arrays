package arrays

/* 
# Push

 the push method appends one or more new items to the end of the array

 examples:

 # append a single element:

	foods := arrays.New[string]();
	foods.push("pizza")

 # append multiple elements:

 foods.Push("hamburguer", "lasagna", "fries")

 # append slice or Array:

	foods.Push([]string{"tomato", "potato"}...)
*/

func (a *Array[T]) Push(items ...T){
  a.elements = append(a.elements, items...)
  a.lastPos += len(items)
}


