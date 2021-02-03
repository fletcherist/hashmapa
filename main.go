package main

import (
	"errors"
	"fmt"
)

// ErrKeyNotFound error when key not found
var ErrKeyNotFound = errors.New("key not found")

// KeyValue represents key-value pair
type KeyValue struct {
	Key   []byte
	Value []byte
}

// Hashmapa interface
type Hashmapa interface {
	Get(key []byte) ([]byte, error)
	Set(key []byte, value []byte) error
	Delete(key []byte) error
	GetRange(from []byte, to []byte) ([]KeyValue, error)
}

func main() {
	hashmapa := NewMemoryMap()
	fmt.Println("test", hashmapa)
}
