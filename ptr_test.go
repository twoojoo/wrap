package wrap_test

import (
	"encoding/json"
	"encoding/xml"
	"testing"

	"wrap"

	"github.com/stretchr/testify/assert"
)

func TestNewPtr(t *testing.T) {
	value := 42
	ptr := wrap.NewPtr(&value)

	assert.NotNil(t, ptr.Unwrap())
	assert.Equal(t, 42, *ptr.Unwrap())

	nilPtr := wrap.NewPtr[int](nil)
	assert.Nil(t, nilPtr.Unwrap())
}

func TestUnwrap(t *testing.T) {
	value := 42
	ptr := wrap.NewPtr(&value)
	assert.Equal(t, &value, ptr.Unwrap())

	nilPtr := wrap.NewPtr[int](nil)
	assert.Nil(t, nilPtr.Unwrap())
}

func TestGetValue(t *testing.T) {
	value := 42
	ptr := wrap.NewPtr(&value)
	val, ok := ptr.GetValue()
	assert.True(t, ok)
	assert.Equal(t, 42, val)

	nilPtr := wrap.NewPtr[int](nil)
	val, ok = nilPtr.GetValue()
	assert.False(t, ok)
	assert.Equal(t, 0, val)
}

func TestSetValue(t *testing.T) {
	ptr := wrap.NewPtr[int](nil)
	ptr.SetValue(42)
	val, ok := ptr.GetValue()
	assert.True(t, ok)
	assert.Equal(t, 42, val)

	ptr.SetValue(100)
	val, ok = ptr.GetValue()
	assert.True(t, ok)
	assert.Equal(t, 100, val)
}

func TestClear(t *testing.T) {
	value := 42
	ptr := wrap.NewPtr(&value)
	assert.NotNil(t, ptr.Unwrap())

	ptr.Clear()
	assert.Nil(t, ptr.Unwrap())

	ptr.Clear() // Clear again to make sure no error
	assert.Nil(t, ptr.Unwrap())
}

func TestIsNil(t *testing.T) {
	ptr := wrap.NewPtr[int](nil)
	assert.True(t, ptr.IsNil())

	value := 42
	ptr = wrap.NewPtr(&value)
	assert.False(t, ptr.IsNil())
}

func TestMarshalJSON(t *testing.T) {
	value := 42
	ptr := wrap.NewPtr(&value)
	data, err := json.Marshal(ptr)
	assert.NoError(t, err)
	assert.Equal(t, "42", string(data))

	nilPtr := wrap.NewPtr[int](nil)
	data, err = json.Marshal(nilPtr)
	assert.NoError(t, err)
	assert.Equal(t, "null", string(data))
}

func TestUnmarshalJSON(t *testing.T) {
	var ptr wrap.Ptr[int]
	err := json.Unmarshal([]byte("42"), &ptr)
	assert.NoError(t, err)
	val, ok := ptr.GetValue()
	assert.True(t, ok)
	assert.Equal(t, 42, val)

	err = json.Unmarshal([]byte("null"), &ptr)
	assert.NoError(t, err)
	assert.True(t, ptr.IsNil())
}

func TestMarshalXML(t *testing.T) {
	type Test struct {
		Value wrap.Ptr[int] `xml:"value"`
	}

	value := 42
	ptr := wrap.NewPtr(&value)
	testStruct := Test{Value: ptr}
	data, err := xml.Marshal(testStruct)
	assert.NoError(t, err)
	assert.Contains(t, string(data), "<value>42</value>")

	nilPtr := wrap.NewPtr[int](nil)
	testStruct = Test{Value: nilPtr}
	data, err = xml.Marshal(testStruct)
	assert.NoError(t, err)
	assert.Equal(t, string(data), "<Test></Test>")
}

func TestUnmarshalXML(t *testing.T) {
	type Test struct {
		Value wrap.Ptr[int] `xml:"value"`
	}

	var testStruct Test
	err := xml.Unmarshal([]byte("<Test><value>42</value></Test>"), &testStruct)
	assert.NoError(t, err)
	val, ok := testStruct.Value.GetValue()
	assert.True(t, ok)
	assert.Equal(t, 42, val)

	err = xml.Unmarshal([]byte("<Test><value></value></Test>"), &testStruct)
	assert.NoError(t, err)
	assert.True(t, testStruct.Value.IsNil())
}