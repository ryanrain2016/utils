package collections

import "errors"

var (
	ErrKeyError                = errors.New("key error")
	ErrIndexOutofRange         = errors.New("index out of range")
	ErrElementNotFound         = errors.New("element not found")
	ErrUnsupportType           = errors.New("unsupport type")
	ErrPopFromEmptyCollections = errors.New("pop or remove from empty collections")
)
