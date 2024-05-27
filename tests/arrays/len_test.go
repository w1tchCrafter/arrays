package arrays_test

import "github.com/w1tchCrafter/arrays/pkg/arrays"
import "testing"

func TestLen(t *testing.T){
  arr := arrays.New(1,2,3,4,5,6,7,8,9,10)

  if arr.Len() != 10 {
    t.Fail()
  }
}
