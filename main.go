package hmap

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

// Hmap interface
type Hmap interface {
	Get(key []byte) ([]byte, error)
	Set(key []byte, value []byte) error
	Delete(key []byte) error
	GetRange(from []byte, to []byte) ([]KeyValue, error)
}

func main() {
	hashmapa := NewMemoryMap()
	fmt.Println("test", hashmapa)
}
