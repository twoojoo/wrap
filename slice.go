package wrap

import (
	"encoding/json"
	"slices"
)

// Slice is a generic wrapper for a slice of type T.
type Slice[T any] struct {
	X []T
}

// NewSlice creates a new Slice instance with the provided initial slice.
func NewSlice[T any](slice []T) Slice[T] {
	return Slice[T]{
		X: slice,
	}
}

// Unwrap returns the underlying slice of type T.
func (s *Slice[T]) Unwrap() []T {
	return s.X
}

// ValueAt retrieves the value at the specified index and a boolean indicating success.
func (s *Slice[T]) ValueAt(index int) (T, bool) {
	var zero T
	if index < 0 || index >= len(s.X) {
		return zero, false
	}
	return s.X[index], true
}

// SetValueAt sets the value at the specified index and returns whether the operation was successful.
func (s *Slice[T]) SetValueAt(index int, value T) bool {
	if index < 0 || index >= len(s.X) {
		return false
	}
	s.X[index] = value
	return true
}

// Append adds one or more values to the end of the slice.
func (s *Slice[T]) Append(values ...T) {
	s.X = append(s.X, values...)
}

// Prepend adds one or more values to the beginning of the slice.
func (s *Slice[T]) Prepend(values ...T) {
	s.X = append(values, s.X...)
}

// Pop removes and returns the last value from the slice, or false if the slice is empty.
func (s *Slice[T]) Pop() (T, bool) {
	var zero T
	if len(s.X) == 0 {
		return zero, false
	}
	value := s.X[len(s.X)-1]
	s.X = s.X[:len(s.X)-1]
	return value, true
}

// Shift removes and returns the first value from the slice, or false if the slice is empty.
func (s *Slice[T]) Shift() (T, bool) {
	var zero T
	if len(s.X) == 0 {
		return zero, false
	}
	value := s.X[0]
	s.X = s.X[1:]
	return value, true
}

// InsertAt inserts one or more values at the specified index and returns whether the operation was successful.
func (s *Slice[T]) InsertAt(index int, values ...T) bool {
	if index < 0 || index > len(s.X) {
		return false
	}
	s.X = append(s.X[:index], append(values, s.X[index:]...)...)
	return true
}

// RemoveAt removes a specified number of elements starting from the index and returns the removed elements as a new Slice.
func (s *Slice[T]) RemoveAt(index int, count ...int) Slice[T] {
	if index < 0 || index >= len(s.X) {
		return NewSlice([]T{})
	}

	c := 1
	if len(count) > 0 {
		c = count[0]
	}

	if c > 0 && index+c > len(s.X) {
		c = len(s.X) - index
	}

	if c <= 0 {
		return NewSlice([]T{})
	}

	removedSl := s.X[index : index+c]
	var removed = make([]T, len(removedSl))
	copy(removed, removedSl)

	s.X = append(s.X[:index], s.X[index+c:]...)

	return NewSlice(removed)
}

// Remove deletes all elements that satisfy the provided comparison function and returns removed elements as a new Slice.
func (s *Slice[T]) Remove(compare func(T) bool) Slice[T] {
	removed := NewSlice([]T{})

	for i := 0; i < len(s.X); i++ {
		if compare(s.X[i]) {
			removed.Append(s.X[i])
			s.X = append(s.X[:i], s.X[i+1:]...)
			i--
		}
	}

	return removed
}

// Length returns the number of elements in the slice.
func (s *Slice[T]) Length() int {
	return len(s.X)
}

// Capacity returns the capacity of the underlying slice.
func (s *Slice[T]) Capacity() int {
	return cap(s.X)
}

// Clear removes all elements from the slice.
func (s *Slice[T]) Clear() {
	s.X = []T{}
}

// SetCapacity changes the capacity of the slice if the new capacity is greater than the current length.
func (s *Slice[T]) SetCapacity(cap int) {
	if cap >= len(s.X) {
		extended := make([]T, len(s.X), cap)
		copy(extended, s.X)
		s.X = extended
	}
}

// Crop reduces the slice to contain only elements up to the specified index.
func (s *Slice[T]) Crop(index int) {
	if index < 0 || index >= len(s.X) {
		return
	}
	s.X = s.X[:index]
}

// Copy returns a new Slice that is a copy of the current slice.
func (s *Slice[T]) Copy() Slice[T] {
	copied := make([]T, len(s.X))
	copy(copied, s.X)
	return NewSlice(copied)
}

// Compact removes consecutive duplicate elements from the slice based on the provided comparison function.
func (s *Slice[T]) Compact(compare func(a, b T) bool) Slice[T] {
	return NewSlice(slices.CompactFunc(s.X, compare))
}

// IndexOf returns the index of the first element that satisfies the provided equality function, or -1 if not found.
func (s *Slice[T]) IndexOf(equals func(T) bool) (int, bool) {
	for i, v := range s.X {
		if equals(v) {
			return i, true
		}
	}
	return -1, false
}

// Find returns the first element that satisfies the provided predicate function, or a zero value if not found.
func (s *Slice[T]) Find(predicate func(T) bool) (T, bool) {
	var zero T
	for _, v := range s.X {
		if predicate(v) {
			return v, true
		}
	}
	return zero, false
}

// Filter returns a new Slice containing only elements that satisfy the provided predicate function.
func (s *Slice[T]) Filter(predicate func(T) bool) Slice[T] {
	filtered := make([]T, 0, len(s.X))
	for _, v := range s.X {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}
	return NewSlice(filtered)
}

// Contains returns true if there is an element that satisfies the provided equality function.
func (s *Slice[T]) Contains(equals func(T) bool) bool {
	_, exists := s.IndexOf(equals)
	return exists
}

// UnmarshalJSON unmarshals JSON data into the Slice.
func (s *Slice[T]) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &s.X); err != nil {
		return err
	}
	return nil
}

// MarshalJSON marshals the Slice into JSON.
func (s Slice[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.X)
}
