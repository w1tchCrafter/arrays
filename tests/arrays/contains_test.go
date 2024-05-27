package arrays_test

import (
	"github.com/w1tchCrafter/arrays/pkg/arrays"
	"testing"
)

func TestContains(t *testing.T) {
	arr := arrays.New("john", "JoHn", "PIzzA", "FoO", "bar")

	if arr.Contains("not in the array") {
		t.Fail()
	}

	if !arr.Contains("PIzzA") {
		t.Fail()
	}

	if !arr.Contains("bar") {
		t.Fail()
	}

	if arr.Contains("Doe") {
		t.Fail()
	}
}
