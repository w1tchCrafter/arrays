package arrays_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/w1tchCrafter/arrays/pkg/arrays"
)

func TestReduce(t *testing.T) {
	arr := arrays.New(25, 18, 32, 20, 15)

	sum := arrays.Reduce(arr, func(acm, item, _ int) int {
		return acm + item
	})

	if sum != 110 {
		t.Fail()
	}

	max := arrays.Reduce(arr, func(acm, item, _ int) int {
		return int(math.Max(float64(acm), float64(item)))
	})

	if max != 32 {
		t.Fail()
	}

	min := arrays.Reduce(arr, func(acm, item, _ int) int {
		return int(math.Min(float64(acm), float64(item)))
	})

	if min != 15 {
		t.Log(min)
		t.Fail()
	}

	str := arrays.Reduce(arr, func(acm string, item, _ int) string {
		return fmt.Sprint(acm, item)
	})

	if str != "2518322015" {
		t.Fail()
	}
}
