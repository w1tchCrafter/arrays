package arrays

import (
	"errors"
)

var ErrEmptyArray = errors.New("the array element is empty")
var ErrIndexOutRange = errors.New("index out of range")
var ErrNotFound = errors.New("item was not found in array")
var ErrWrongUsage = errors.New("no index passed to 'ToSlice' method")
