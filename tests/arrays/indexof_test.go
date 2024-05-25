package arrays_test

import(
  "github.com/w1tchCrafter/arrays/pkg/arrays"
  "testing"
)

func TestIndexOf(t *testing.T) {
	arr := arrays.New("a", "b", "c", "d", "e", "f", "g", "d", "g")
	indexD := arr.IndexOf("d") // should be 3
	indexG := arr.IndexOf("g") // should be 6

	if indexG != 6 || indexD != 3 {
		t.Log(indexD, indexG)
		t.Fail()
	}

	ni := arr.IndexOf("z")

	if ni != -1 {
		t.Fail()
	}
}
