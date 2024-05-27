package arrays

import (
	"fmt"
	"strings"
)

/*
# Join

join returns a string made of the Array elements separated by the paramenter sep

# works with any datatype not only strings

	person1 := map[string]any{"name": "John", "age": 19}
	person2 := map[string]any{"name": "Mary", "age": 23}

	arr := arrays.New(person1, person2)

	fmt.Println(arr.Join(" ")) // prints: "map[age:19 name:John] map[age:23 name:Mary]"

# Joining array of strings with a comma

	arr := arrays.New("I", "love", "pizza")

	fmt.Println(arr.Join(" ")) // prints: "I love pizza"
*/
func (a Array[T]) Join(sep string) string {
	strResult := make([]string, 0)

	for _, v := range a.elements {
		strResult = append(strResult, fmt.Sprint(v))
	}

	return strings.Join(strResult, sep)
}
