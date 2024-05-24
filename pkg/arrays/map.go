package arrays

func Map [T, E any] (a Array[T], f func(int, T) E) []E {
  result := make([]E, 0)

  for i,v := range a.elements {
    result = append(result, f(i,v))
  }

  return result
}
