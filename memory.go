package hmap

import (
	"sort"
	"sync"
)

// MemoryMap object itself
type MemoryMap struct {
	mut   *sync.Mutex
	kv    map[string][]byte
	index []string
}

// Get value by key
func (h MemoryMap) Get(key []byte) ([]byte, error) {
	h.mut.Lock()
	defer h.mut.Unlock()
	value, exists := h.kv[string(key)]
	if !exists {
		return nil, ErrKeyNotFound
	}
	return value, nil
}

// Set value by key
func (h MemoryMap) Set(key []byte, value []byte) error {
	// check if key exists in hashmap
	if _, err := h.Get(key); err != nil {
		// if key does not exist, then add it to index
		if err == ErrKeyNotFound {
			h.indexAdd(string(key))
		} else {
			panic(err)
		}
	}
	h.mut.Lock()
	defer h.mut.Unlock()
	h.kv[string(key)] = value
	return nil
}

// Delete removes key from hashmap
func (h MemoryMap) Delete(key []byte) error {
	h.mut.Lock()
	defer h.mut.Unlock()
	delete(h.kv, string(key))
	h.indexDelete(string(key))
	return nil
}

// GetRange returns list of sorted values
func (h MemoryMap) GetRange(from []byte, to []byte) ([]KeyValue, error) {
	h.mut.Lock()
	defer h.mut.Unlock()
	result := []KeyValue{}
	for _, key := range h.index {
		// value is in range
		if key >= string(from) && key <= string(to) {
			value, err := h.Get([]byte(key))
			if err != nil {
				return nil, err
			}
			result = append(result, KeyValue{
				Key:   []byte(key),
				Value: value,
			})
		}
	}
	return result, nil
}

func (h *MemoryMap) indexAdd(key string) {
	h.mut.Lock()
	defer h.mut.Unlock()
	h.index = append(h.index, string(key))
	h.indexSort()
}

func (h *MemoryMap) indexDelete(key string) {
	h.mut.Lock()
	defer h.mut.Unlock()
	newIndex := []string{}
	for _, indexKey := range h.index {
		// remove this key from index by not adding it
		if indexKey == key {
			continue
		}
		newIndex = append(newIndex, indexKey)
	}
	h.index = newIndex
}

func (h *MemoryMap) indexSort() {
	sort.Slice(h.index, func(i, j int) bool {
		return h.index[i] < h.index[j]
	})
}

// NewMemoryMap creates new MemoryMap instance
func NewMemoryMap() Hmap {
	hmap := MemoryMap{
		mut:   &sync.Mutex{},
		kv:    make(map[string][]byte),
		index: []string{},
	}
	return hmap
}
