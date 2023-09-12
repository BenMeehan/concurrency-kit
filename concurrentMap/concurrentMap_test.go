package concurrentmap

import (
	"fmt"
	"testing"
)

func TestConcurrentMap(t *testing.T) {
	// Create a new concurrent map
	cm := New[string, int]()

	// Test Set and Get
	cm.Set("one", 1)
	val, exists := cm.Get("one")
	if !exists {
		t.Errorf("Expected 'one' to exist in the map")
	}
	if val != 1 {
		t.Errorf("Expected value for 'one' to be 1, got %d", val)
	}

	// Test Len
	length := cm.Len()
	if length != 1 {
		t.Errorf("Expected length of the map to be 1, got %d", length)
	}

	// Test Delete
	cm.Delete("one")
	val, exists = cm.Get("one")
	if exists {
		t.Errorf("Expected 'one' to be deleted from the map")
	}

	// Test Keys and Values
	cm.Set("two", 2)
	cm.Set("three", 3)
	keys := cm.Keys()
	values := cm.Values()

	expectedKeys := []string{"two", "three"}
	expectedValues := []int{2, 3}

	for i, key := range expectedKeys {
		if keys[i] != key {
			t.Errorf("Expected key %s at index %d, got %s", key, i, keys[i])
		}
		if values[i] != expectedValues[i] {
			t.Errorf("Expected value %d at index %d, got %d", expectedValues[i], i, values[i])
		}
	}

	// Test non-existent key
	val, exists = cm.Get("four")
	if exists {
		t.Errorf("Expected 'four' not to exist in the map")
	}

	// Test concurrent Set
	for i := 0; i < 1000; i++ {
		go func(i int) {
			key := fmt.Sprintf("%d", i)
			cm.Set(key, i)
		}(i)
	}

	// Test Clear
	cm.Clear()
	length = cm.Len()
	if length != 0 {
		t.Errorf("Expected length of the map to be 0 after clearing, got %d", length)
	}
}