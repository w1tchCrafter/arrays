package arrays_test

import(
  "testing"
  "github.com/w1tchCrafter/arrays/pkg/arrays"
)

func TestJoin(t *testing.T) {
  arr := arrays.New("john", "doe", "is", "foo", "bar")

  if arr.Join(" ") != "john doe is foo bar" {
    t.Fail()
  }

  person1 := map[string]any{"name": "John", "age": 19}
  person2 := map[string]any{"name": "Mary", "age": 23}

  arr2 := arrays.New(person1, person2)

  t.Log(arr2.Join(" "))
}
