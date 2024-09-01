package wrap

import (
	"encoding/json"
	"slices"
)

// Slice is a generic wrapper for a slice of type T.
type Slice[T any] struct {
	x []T
}

// NewSlice creates a new Slice instance with the provided initial slice.
func NewSlice[T any](slice []T) Slice[T] {
	return Slice[T]{
		x: slice,
	}
}

// Unwrap returns the underlying slice of type T.
func (s *Slice[T]) Unwrap() []T {
	return s.x
}

// ValueAt retrieves the value at the specified index and a boolean indicating success.
func (s *Slice[T]) ValueAt(index int) (T, bool) {
	var zero T
	if index < 0 || index >= len(s.x) {
		return zero, false
	}
	return s.x[index], true
}

// SetValueAt sets the value at the specified index and returns whether the operation was successful.
func (s *Slice[T]) SetValueAt(index int, value T) bool {
	if index < 0 || index >= len(s.x) {
		return false
	}
	s.x[index] = value
	return true
}

// Append adds one or more values to the end of the slice.
func (s *Slice[T]) Append(values ...T) {
	s.x = append(s.x, values...)
}

// Prepend adds one or more values to the beginning of the slice.
func (s *Slice[T]) Prepend(values ...T) {
	s.x = append(values, s.x...)
}

// Pop removes and returns the last value from the slice, or false if the slice is empty.
func (s *Slice[T]) Pop() (T, bool) {
	var zero T
	if len(s.x) == 0 {
		return zero, false
	}
	value := s.x[len(s.x)-1]
	s.x = s.x[:len(s.x)-1]
	return value, true
}

// Shift removes and returns the first value from the slice, or false if the slice is empty.
func (s *Slice[T]) Shift() (T, bool) {
	var zero T
	if len(s.x) == 0 {
		return zero, false
	}
	value := s.x[0]
	s.x = s.x[1:]
	return value, true
}

// InsertAt inserts one or more values at the specified index and returns whether the operation was successful.
func (s *Slice[T]) InsertAt(index int, values ...T) bool {
	if index < 0 || index > len(s.x) {
		return false
	}
	s.x = append(s.x[:index], append(values, s.x[index:]...)...)
	return true
}

// RemoveAt removes a specified number of elements starting from the index and returns the removed elements as a new Slice.
func (s *Slice[T]) RemoveAt(index int, count ...int) Slice[T] {
	if index < 0 || index >= len(s.x) {
		return NewSlice([]T{})
	}

	c := 1
	if len(count) > 0 {
		c = count[0]
	}

	if c > 0 && index+c > len(s.x) {
		c = len(s.x) - index
	}

	if c <= 0 {
		return NewSlice([]T{})
	}

	removedSl := s.x[index : index+c]
	var removed = make([]T, len(removedSl))
	copy(removed, removedSl)

	s.x = append(s.x[:index], s.x[index+c:]...)

	return NewSlice(removed)
}

// Remove deletes all elements that satisfy the provided comparison function and returns removed elements as a new Slice.
func (s *Slice[T]) Remove(compare func(T) bool) Slice[T] {
	removed := NewSlice([]T{})

	for i := 0; i < len(s.x); i++ {
		if compare(s.x[i]) {
			removed.Append(s.x[i])
			s.x = append(s.x[:i], s.x[i+1:]...)
			i--
		}
	}

	return removed
}

// Length returns the number of elements in the slice.
func (s *Slice[T]) Length() int {
	return len(s.x)
}

// Capacity returns the capacity of the underlying slice.
func (s *Slice[T]) Capacity() int {
	return cap(s.x)
}

// Clear removes all elements from the slice.
func (s *Slice[T]) Clear() {
	s.x = []T{}
}

// SetCapacity changes the capacity of the slice if the new capacity is greater than the current length.
func (s *Slice[T]) SetCapacity(cap int) {
	if cap >= len(s.x) {
		extended := make([]T, len(s.x), cap)
		copy(extended, s.x)
		s.x = extended
	}
}

// Crop reduces the slice to contain only elements up to the specified index.
func (s *Slice[T]) Crop(index int) {
	if index < 0 || index >= len(s.x) {
		return
	}
	s.x = s.x[:index]
}

// Copy returns a new Slice that is a copy of the current slice.
func (s *Slice[T]) Copy() Slice[T] {
	copied := make([]T, len(s.x))
	copy(copied, s.x)
	return NewSlice(copied)
}

// Compact removes consecutive duplicate elements from the slice based on the provided comparison function.
func (s *Slice[T]) Compact(compare func(a, b T) bool) Slice[T] {
	return NewSlice(slices.CompactFunc(s.x, compare))
}

// IndexOf returns the index of the first element that satisfies the provided equality function, or -1 if not found.
func (s *Slice[T]) IndexOf(equals func(T) bool) (int, bool) {
	for i, v := range s.x {
		if equals(v) {
			return i, true
		}
	}
	return -1, false
}

// Find returns the first element that satisfies the provided predicate function, or a zero value if not found.
func (s *Slice[T]) Find(predicate func(T) bool) (T, bool) {
	var zero T
	for _, v := range s.x {
		if predicate(v) {
			return v, true
		}
	}
	return zero, false
}

// Filter returns a new Slice containing only elements that satisfy the provided predicate function.
func (s *Slice[T]) Filter(predicate func(T) bool) Slice[T] {
	filtered := make([]T, 0, len(s.x))
	for _, v := range s.x {
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
	if err := json.Unmarshal(data, &s.x); err != nil {
		return err
	}
	return nil
}

// MarshalJSON marshals the Slice into JSON.
func (s Slice[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.x)
}
