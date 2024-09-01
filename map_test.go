package wrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap_SetAndGet(t *testing.T) {
	tests := []struct {
		initial map[string]int
		setKey   string
		setValue int
		getKey   string
		expected int
		exists   bool
	}{
		{
			initial:  map[string]int{"a": 1, "b": 2},
			setKey:   "c",
			setValue: 3,
			getKey:   "c",
			expected: 3,
			exists:   true,
		},
		{
			initial:  map[string]int{"a": 1, "b": 2},
			setKey:   "b",
			setValue: 10,
			getKey:   "b",
			expected: 10,
			exists:   true,
		},
		{
			initial:  map[string]int{"a": 1},
			setKey:   "d",
			setValue: 4,
			getKey:   "d",
			expected: 4,
			exists:   true,
		},
		{
			initial:  map[string]int{},
			setKey:   "e",
			setValue: 5,
			getKey:   "e",
			expected: 5,
			exists:   true,
		},
	}

	for _, tt := range tests {
		m := NewMap(tt.initial)
		m.Set(tt.setKey, tt.setValue)
		value, exists := m.Get(tt.getKey)
		assert.Equal(t, tt.expected, value, "Value for key '%s' should be %d", tt.getKey, tt.expected)
		assert.Equal(t, tt.exists, exists, "Key '%s' existence should be %v", tt.getKey, tt.exists)
	}
}

func TestMap_Delete(t *testing.T) {
	tests := []struct {
		initial   map[string]int
		deleteKey string
		expected  map[string]int
	}{
		{
			initial:   map[string]int{"a": 1, "b": 2},
			deleteKey: "a",
			expected:  map[string]int{"b": 2},
		},
		{
			initial:   map[string]int{"a": 1, "b": 2},
			deleteKey: "c",
			expected:  map[string]int{"a": 1, "b": 2}, // No change
		},
		{
			initial:   map[string]int{"a": 1},
			deleteKey: "a",
			expected:  map[string]int{}, // All keys removed
		},
	}

	for _, tt := range tests {
		m := NewMap(tt.initial)
		m.Delete(tt.deleteKey)
		assert.Equal(t, tt.expected, m.Unwrap(), "Map should be %v after deletion", tt.expected)
	}
}

func TestMap_Contains(t *testing.T) {
	tests := []struct {
		initial   map[string]int
		checkKey  string
		expected  bool
	}{
		{
			initial:   map[string]int{"a": 1, "b": 2},
			checkKey:  "a",
			expected:  true,
		},
		{
			initial:   map[string]int{"a": 1, "b": 2},
			checkKey:  "c",
			expected:  false,
		},
		{
			initial:   map[string]int{},
			checkKey:  "a",
			expected:  false,
		},
	}

	for _, tt := range tests {
		m := NewMap(tt.initial)
		assert.Equal(t, tt.expected, m.Contains(tt.checkKey), "Map should %v contain key '%s'", tt.expected, tt.checkKey)
	}
}

func TestMap_Keys(t *testing.T) {
	tests := []struct {
		initial  map[string]int
		expected []string
	}{
		{
			initial:  map[string]int{"a": 1, "b": 2},
			expected: []string{"a", "b"},
		},
		{
			initial:  map[string]int{},
			expected: []string{},
		},
		{
			initial:  map[string]int{"x": 10, "y": 20, "z": 30},
			expected: []string{"x", "y", "z"},
		},
	}

	for _, tt := range tests {
		m := NewMap(tt.initial)
		keys := m.Keys()
		assert.ElementsMatch(t, tt.expected, keys, "Keys should be %v", tt.expected)
	}
}

func TestMap_Values(t *testing.T) {
	tests := []struct {
		initial  map[string]int
		expected []int
	}{
		{
			initial:  map[string]int{"a": 1, "b": 2},
			expected: []int{1, 2},
		},
		{
			initial:  map[string]int{},
			expected: []int{},
		},
		{
			initial:  map[string]int{"x": 10, "y": 20, "z": 30},
			expected: []int{10, 20, 30},
		},
	}

	for _, tt := range tests {
		m := NewMap(tt.initial)
		values := m.Values()
		assert.ElementsMatch(t, tt.expected, values, "Values should be %v", tt.expected)
	}
}

func TestMap_Len(t *testing.T) {
	tests := []struct {
		initial map[string]int
		expected int
	}{
		{
			initial: map[string]int{"a": 1, "b": 2},
			expected: 2,
		},
		{
			initial: map[string]int{},
			expected: 0,
		},
		{
			initial: map[string]int{"x": 10, "y": 20, "z": 30},
			expected: 3,
		},
	}

	for _, tt := range tests {
		m := NewMap(tt.initial)
		assert.Equal(t, tt.expected, m.Len(), "Length of map should be %d", tt.expected)
	}
}

func TestMap_IsEmpty(t *testing.T) {
	tests := []struct {
		initial map[string]int
		expected bool
	}{
		{
			initial: map[string]int{"a": 1},
			expected: false,
		},
		{
			initial: map[string]int{},
			expected: true,
		},
	}

	for _, tt := range tests {
		m := NewMap(tt.initial)
		assert.Equal(t, tt.expected, m.IsEmpty(), "Map should be empty: %v", tt.expected)
	}
}

func TestMap_Clear(t *testing.T) {
	tests := []struct {
		initial  map[string]int
		expected map[string]int
	}{
		{
			initial:  map[string]int{"a": 1, "b": 2},
			expected: map[string]int{},
		},
		{
			initial:  map[string]int{},
			expected: map[string]int{},
		},
	}

	for _, tt := range tests {
		m := NewMap(tt.initial)
		m.Clear()
		assert.Equal(t, tt.expected, m.Unwrap(), "Map should be %v after clearing", tt.expected)
	}
}

func TestMap_MarshalJSON(t *testing.T) {
	tests := []struct {
		initial  map[string]int
		expected string
	}{
		{
			initial:  map[string]int{"a": 1, "b": 2},
			expected: `{"a":1,"b":2}`,
		},
		{
			initial:  map[string]int{},
			expected: `{}`,
		},
	}

	for _, tt := range tests {
		m := NewMap(tt.initial)
		data, err := m.MarshalJSON()
		assert.NoError(t, err, "Marshalling should not return an error")
		assert.JSONEq(t, tt.expected, string(data), "JSON output should be %s", tt.expected)
	}
}

func TestMap_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{
			input:    `{"a":1,"b":2}`,
			expected: map[string]int{"a": 1, "b": 2},
		},
		{
			input:    `{}`,
			expected: map[string]int{},
		},
	}

	for _, tt := range tests {
		var m Map[string, int]
		err := m.UnmarshalJSON([]byte(tt.input))
		assert.NoError(t, err, "Unmarshalling should not return an error")
		assert.Equal(t, tt.expected, m.Unwrap(), "Map should be %v after unmarshalling", tt.expected)
	}
}


func TestMap_Find(t *testing.T) {
	tests := []struct {
		initial  map[string]int
		compare  func(int) bool
		expectedKey string
		expectedExists bool
	}{
		{
			initial:      map[string]int{"a": 1, "b": 2, "c": 3},
			compare:      func(v int) bool { return v == 2 }, // Match value 2
			expectedKey:  "b",
			expectedExists: true,
		},
		{
			initial:      map[string]int{"a": 1, "b": 2, "c": 3},
			compare:      func(v int) bool { return v == 4 }, // No match
			expectedKey:  "",
			expectedExists: false,
		},
		{
			initial:      map[string]int{},
			compare:      func(v int) bool { return v == 1 }, // Empty map
			expectedKey:  "",
			expectedExists: false,
		},
		{
			initial:      map[string]int{"x": 10, "y": 20, "z": 30},
			compare:      func(v int) bool { return v > 15 }, // Match value > 15
			expectedKey:  "y", // 'y' is the first key that matches the condition
			expectedExists: true,
		},
		{
			initial:      map[string]int{"a": 5, "b": 10, "c": 15},
			compare:      func(v int) bool { return v%5 == 0 }, // Match any value divisible by 5
			expectedKey:  "a", // 'a' is the first key that matches the condition
			expectedExists: true,
		},
	}

	for _, tt := range tests {
		m := NewMap(tt.initial)
		key, exists := m.Find(tt.compare)
		assert.Equal(t, tt.expectedKey, key, "Key should be %s", tt.expectedKey)
		assert.Equal(t, tt.expectedExists, exists, "Existence should be %v", tt.expectedExists)
	}
}