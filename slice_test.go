package wrap

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSlice_NewSlice(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{}, []int{}},
	}

	for _, tt := range tests {
		s := NewSlice(tt.input)
		assert.ElementsMatch(t, tt.expected, s.Unwrap())
	}
}

func TestSlice_ValueAt(t *testing.T) {
	s := NewSlice([]int{1, 2, 3})

	tests := []struct {
		index    int
		expected int
		ok       bool
	}{
		{1, 2, true},
		{0, 1, true},
		{3, 0, false}, // out of bounds
		{-1, 0, false}, // negative index
	}

	for _, tt := range tests {
		got, ok := s.ValueAt(tt.index)
		assert.Equal(t, tt.ok, ok)
		assert.Equal(t, tt.expected, got)
	}
}

func TestSlice_SetValueAt(t *testing.T) {
	s := NewSlice([]int{1, 2, 3})

	tests := []struct {
		index    int
		value    int
		expected []int
		ok       bool
	}{
		{1, 10, []int{1, 10, 3}, true},
		{3, 10, []int{1, 10, 3}, false}, // out of bounds
		{-1, 10, []int{1, 10, 3}, false}, // negative index
	}

	for _, tt := range tests {
		ok := s.SetValueAt(tt.index, tt.value)
		assert.Equal(t, tt.ok, ok)
		assert.ElementsMatch(t, tt.expected, s.Unwrap())
	}
}

func TestSlice_Append(t *testing.T) {
	s := NewSlice([]int{1, 2})

	s.Append(3)
	expected := []int{1, 2, 3}
	assert.ElementsMatch(t, expected, s.Unwrap())
}

func TestSlice_Pop(t *testing.T) {
	s := NewSlice([]int{1, 2, 3})

	value, ok := s.Pop()
	expected := 3
	assert.True(t, ok)
	assert.Equal(t, expected, value)
	assert.Len(t, s.Unwrap(), 2)
}

func TestSlice_Shift(t *testing.T) {
	s := NewSlice([]int{1, 2, 3})

	value, ok := s.Shift()
	expected := 1
	assert.True(t, ok)
	assert.Equal(t, expected, value)
	assert.Len(t, s.Unwrap(), 2)
}

func TestSlice_InsertAt(t *testing.T) {
	s := NewSlice([]int{1, 3})

	ok := s.InsertAt(1, 2)
	expected := []int{1, 2, 3}
	assert.True(t, ok)
	assert.ElementsMatch(t, expected, s.Unwrap())
}

// func TestSlice_Delete(t *testing.T) {
// 	s := NewSlice([]int{1, 2, 3})

// 	newS := s.Delete(func(v int) bool { return v == 2 })
// 	expected := []int{1, 3}
// 	assert.ElementsMatch(t, expected, newS.Unwrap())
// 	assert.Len(t, s.Unwrap(), 2)
// }

func TestSlice_Length(t *testing.T) {
	s := NewSlice([]int{1, 2, 3})

	assert.Equal(t, 3, s.Length())
}

func TestSlice_Capacity(t *testing.T) {
	s := NewSlice([]int{1, 2, 3})

	assert.GreaterOrEqual(t, s.Capacity(), 3)
}

func TestSlice_Clear(t *testing.T) {
	s := NewSlice([]int{1, 2, 3})

	s.Clear()
	expected := []int{}
	assert.ElementsMatch(t, expected, s.Unwrap())
}

func TestSlice_SetCapacity(t *testing.T) {
	s := NewSlice([]int{1, 2, 3})

	s.SetCapacity(10)
	assert.GreaterOrEqual(t, s.Capacity(), 10)
}

func TestSlice_Crop(t *testing.T) {
	s := NewSlice([]int{1, 2, 3})

	s.Crop(2)
	expected := []int{1, 2}
	assert.ElementsMatch(t, expected, s.Unwrap())
}

func TestSlice_Copy(t *testing.T) {
	s := NewSlice([]int{1, 2, 3})

	copied := s.Copy()
	expected := []int{1, 2, 3}
	assert.ElementsMatch(t, expected, copied.Unwrap())
}

func TestSlice_Compact(t *testing.T) {
	s := NewSlice([]int{1, 1, 2, 2, 3})

	compact := s.Compact(func(a, b int) bool { return a == b })
	expected := []int{1, 2, 3}
	assert.ElementsMatch(t, expected, compact.Unwrap())
}

func TestSlice_IndexOf(t *testing.T) {
	s := NewSlice([]int{1, 2, 3})

	index, ok := s.IndexOf(func(v int) bool { return v == 2 })
	assert.True(t, ok)
	assert.Equal(t, 1, index)
}

func TestSlice_Find(t *testing.T) {
	s := NewSlice([]int{1, 2, 3})

	value, ok := s.Find(func(v int) bool { return v == 2 })
	assert.True(t, ok)
	assert.Equal(t, 2, value)
}

func TestSlice_Filter(t *testing.T) {
	s := NewSlice([]int{1, 2, 3})

	filtered := s.Filter(func(v int) bool { return v > 1 })
	expected := []int{2, 3}
	assert.ElementsMatch(t, expected, filtered.Unwrap())
}

func TestSlice_Contains(t *testing.T) {
	s := NewSlice([]int{1, 2, 3})

	assert.True(t, s.Contains(func(v int) bool { return v == 2 }))
}


func TestSlice_Remove(t *testing.T) {
	tests := []struct {
		initial  []int
		compare  func(int) bool
		expected []int
		removed  []int
	}{
		// Remove some elements
		{
			initial:  []int{1, 2, 3, 4, 5},
			compare:  func(v int) bool { return v%2 == 0 }, // Remove even numbers
			expected: []int{1, 3, 5},
			removed:  []int{2, 4},
		},
		// Remove all elements
		{
			initial:  []int{1, 2, 3},
			compare:  func(v int) bool { return v > 0 }, // Remove all elements
			expected: []int{},
			removed:  []int{1, 2, 3},
		},
		// Remove no elements
		{
			initial:  []int{1, 2, 3},
			compare:  func(v int) bool { return v > 5 }, // No elements match
			expected: []int{1, 2, 3},
			removed:  []int{},
		},
		// Remove elements from an empty slice
		{
			initial:  []int{},
			compare:  func(v int) bool { return v == 1 }, // Empty slice, no elements
			expected: []int{},
			removed:  []int{},
		},
		// Remove elements with more complex condition
		{
			initial:  []int{10, 20, 30, 40, 50},
			compare:  func(v int) bool { return v%20 == 0 }, // Remove multiples of 20
			expected: []int{10, 30, 50},
			removed:  []int{20, 40},
		},
		// Remove elements with the same value
		{
			initial:  []int{5, 5, 5, 5, 5},
			compare:  func(v int) bool { return v == 5 }, // Remove all elements that are 5
			expected: []int{},
			removed:  []int{5, 5, 5, 5, 5},
		},
	}

	for _, tt := range tests {
		s := NewSlice(tt.initial)

		// Perform the remove operation
		removed := s.Remove(tt.compare)

		// Assert the removed elements
		assert.ElementsMatch(t, tt.removed, removed.Unwrap(), "Removed elements should be %v", tt.removed)

		// Assert the final state of the slice
		assert.ElementsMatch(t, tt.expected, s.Unwrap(), "Slice after Remove() should be %v", tt.expected)
	}
}

func TestSlice_RemoveAt(t *testing.T) {
	tests := []struct {
		initial  []int
		index    int
		count    int
		expected []int
		removed  []int
	}{
		// Remove a single element from the middle
		{
			initial:  []int{1, 2, 3, 4, 5},
			index:    2,
			count:    1,
			expected: []int{1, 2, 4, 5},
			removed:  []int{3},
		},
		// Remove multiple elements from the middle
		{
			initial:  []int{1, 2, 3, 4, 5},
			index:    1,
			count:    3,
			expected: []int{1, 5},
			removed:  []int{2, 3, 4},
		},
		// Remove from index greater than length
		{
			initial:  []int{1, 2, 3, 4, 5},
			index:    10,
			count:    1,
			expected: []int{1, 2, 3, 4, 5},
			removed:  []int{},
		},
		// Remove more elements than exist
		{
			initial:  []int{1, 2, 3, 4, 5},
			index:    3,
			count:    10,
			expected: []int{1, 2, 3},
			removed:  []int{4, 5},
		},
		// Remove elements from the start
		{
			initial:  []int{1, 2, 3, 4, 5},
			index:    0,
			count:    2,
			expected: []int{3, 4, 5},
			removed:  []int{1, 2},
		},
		// Remove elements from the end
		{
			initial:  []int{1, 2, 3, 4, 5},
			index:    3,
			count:    2,
			expected: []int{1, 2, 3},
			removed:  []int{4, 5},
		},
		// Remove zero elements
		{
			initial:  []int{1, 2, 3, 4, 5},
			index:    2,
			count:    0,
			expected: []int{1, 2, 3, 4, 5},
			removed:  []int{},
		},
		// Remove negative count
		{
			initial:  []int{1, 2, 3, 4, 5},
			index:    1,
			count:    -1,
			expected: []int{1, 2, 3, 4, 5},
			removed:  []int{},
		},
	}

	for _, tt := range tests {
		s := NewSlice(tt.initial)

		// Perform the remove operation
		removed := s.RemoveAt(tt.index, tt.count)

		// Assert the removed elements
		assert.ElementsMatch(t, tt.removed, removed.Unwrap(), "Removed elements should be %v", tt.removed)

		// Assert the final state of the slice
		assert.ElementsMatch(t, tt.expected, s.Unwrap(), "Slice after RemoveAt() should be %v", tt.expected)
	}
}
