package concurrentmap

import (
	"sync"
)

// ConcurrentMap is a concurrent map data structure.
type ConcurrentMap[K comparable, V any] struct {
    data map[K]V
    mu   sync.RWMutex
}

// New creates and returns a new ConcurrentMap instance.
func New[K comparable, V any]() *ConcurrentMap[K, V] {
    return &ConcurrentMap[K, V]{
        data: make(map[K]V),
    }
}

// Set adds or updates a key-value pair in the map.
func (cm *ConcurrentMap[K, V]) Set(key K, value V) {
    cm.mu.Lock()
    cm.data[key] = value
    cm.mu.Unlock()
}

// Get retrieves a value for a given key from the map.
// It returns the value and a boolean indicating whether the key exists.
func (cm *ConcurrentMap[K, V]) Get(key K) (V, bool) {
    cm.mu.RLock()
    value, exists := cm.data[key]
    cm.mu.RUnlock()
    return value, exists
}

// Len returns the number of key-value pairs in the map.
func (cm *ConcurrentMap[K, V]) Len() int {
    cm.mu.RLock()
    length := len(cm.data)
    cm.mu.RUnlock()
    return length
}

// Delete removes a key and its associated value from the map.
func (cm *ConcurrentMap[K, V]) Delete(key K) {
    cm.mu.Lock()
    delete(cm.data, key)
    cm.mu.Unlock()
}

// Keys returns a slice containing all the keys in the map.
func (cm *ConcurrentMap[K, V]) Keys() []K {
    cm.mu.RLock()
    keys := make([]K, 0, len(cm.data))
    for key := range cm.data {
        keys = append(keys, key)
    }
    cm.mu.RUnlock()
    return keys
}

// Values returns a slice containing all the values in the map.
func (cm *ConcurrentMap[K, V]) Values() []V {
    cm.mu.RLock()
    values := make([]V, 0, len(cm.data))
    for _, value := range cm.data {
        values = append(values, value)
    }
    cm.mu.RUnlock()
    return values
}

// Clear removes all key-value pairs from the map.
func (cm *ConcurrentMap[K, V]) Clear() {
    cm.mu.Lock()
    cm.data = make(map[K]V)
    cm.mu.Unlock()
}