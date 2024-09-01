package wrap

import (
	"encoding/json"
)

// Map is a generic wrapper for a map with keys of type K and values of type V.
type Map[K comparable, V any] struct {
	x map[K]V
}

// Object is a shortcut for Map[string, any]
type Object Map[string, any]

// MapEntry represents a key-value pair in the map.
type MapEntry[K comparable, V any] struct {
	Key   K
	Value V
}

// NewMap creates a new Map instance with the provided initial map.
func NewMap[K comparable, V any](m map[K]V) Map[K, V] {
	return Map[K, V]{
		x: m,
	}
}

// Unwrap returns the underlying map of type map[K]V.
func (m *Map[K, V]) Unwrap() map[K]V {
	return m.x
}

// Get retrieves the value associated with the specified key and a boolean indicating if the key exists.
func (m *Map[K, V]) Get(key K) (V, bool) {
	var zero V
	value, exists := m.x[key]
	if !exists {
		return zero, false
	}
	return value, true
}

// Set adds or updates the value for the specified key and returns the Map instance.
func (m Map[K, V]) Set(key K, value V) Map[K, V] {
	m.x[key] = value
	return m
}

// Delete removes the key-value pair associated with the specified key and returns the Map instance.
func (m Map[K, V]) Delete(key K) Map[K, V] {
	delete(m.x, key)
	return m
}

// Contains checks if the specified key exists in the map.
func (m Map[K, V]) Contains(key K) bool {
	_, exists := m.x[key]
	return exists
}

// Keys returns a slice of all keys in the map.
func (m Map[K, V]) Keys() []K {
	keys := make([]K, 0, len(m.x))
	for key := range m.x {
		keys = append(keys, key)
	}
	return keys
}

// Values returns a slice of all values in the map.
func (m *Map[K, V]) Values() Slice[V] {
	values := make([]V, 0, len(m.x))
	for _, value := range m.x {
		values = append(values, value)
	}
	return NewSlice(values)
}

// Find returns the key of the first value that satisfies the provided comparison function, or zero value and false if not found.
func (m Map[K, V]) Find(compare func(V) bool) (K, bool) {
	var zero K
	for key, value := range m.x {
		if compare(value) {
			return key, true
		}
	}
	return zero, false
}

// Len returns the number of key-value pairs in the map.
func (m Map[K, V]) Len() int {
	return len(m.x)
}

// IsEmpty returns true if the map is empty, otherwise false.
func (m Map[K, V]) IsEmpty() bool {
	return len(m.x) == 0
}

// Clear removes all key-value pairs from the map.
func (m Map[K, V]) Clear() {
	for key := range m.x {
		delete(m.x, key)
	}
}

// UnmarshalJSON unmarshals JSON data into the Map. It expects a JSON object representation.
func (m *Map[K, V]) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &m.x); err != nil {
		return err
	}
	return nil
}

// MarshalJSON marshals the Map into JSON. It produces a JSON object representation.
func (m Map[K, V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.x)
}
