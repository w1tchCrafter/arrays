package arrays

import (
	"errors"
)

var (
	ErrEmptyArray    = errors.New("the array element is empty")
	ErrIndexOutRange = errors.New("index out of range")
	ErrNotFound      = errors.New("item was not found in array")
	ErrWrongUsage    = errors.New("no index passed to 'ToSlice' method")
)
