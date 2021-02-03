package hmap

import (
	"testing"
)

func TestMemory(t *testing.T) {
	t.Run("set/get", func(t *testing.T) {
		hmap := NewMemoryMap()
		testKey := "foo"
		testValue := "bar"
		err := hmap.Set([]byte(testKey), []byte(testValue))
		if err != nil {
			t.Error(err)
			return
		}
		value, err := hmap.Get([]byte(testKey))
		if err != nil {
			t.Error(err)
			return
		}
		if testValue != string(value) {
			t.Error("value is invalid:", value)
		}
	})
}
