package wrap

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Ptr is a generic wrapper for a pointer of type T.
type Ptr[T any] struct {
	x *T
}

// NewPtr creates a new Ptr instance with the provided pointer to T.
func NewPtr[T any](ptr *T) Ptr[T] {
	return Ptr[T]{
		x: ptr,
	}
}

// NewPtr creates a new Ptr instance of type T with no value.
func NewNilPtr[T any]() Ptr[T] {
	return Ptr[T]{
		x: nil,
	}
}

// Unwrap returns the underlying pointer of type T.
func (p *Ptr[T]) Unwrap() *T {
	return p.x
}

// GetValue returns the value pointed to by the Ptr and a boolean indicating if the pointer is non-nil.
func (p *Ptr[T]) GetValue() (T, bool) {
	var zero T
	if p.x == nil {
		return zero, false
	}
	return *p.x, true
}

// SetValue sets the value of the pointer, creating a new pointer if it was previously nil.
func (p *Ptr[T]) SetValue(value T) {
	if p.x == nil {
		p.x = new(T)
	}
	*p.x = value
}

// Clear sets the pointer to nil, effectively clearing its value.
func (p *Ptr[T]) Clear() {
	p.x = nil
}

// IsNil returns true if the pointer is nil, otherwise false.
func (p *Ptr[T]) IsNil() bool {
	return p.x == nil
}

// UnmarshalJSON unmarshals JSON data into the Ptr. It handles both "null" and non-null values.
func (p *Ptr[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		p.x = nil
		return nil
	}

	var value T
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	p.x = &value
	return nil
}

// MarshalJSON marshals the Ptr into JSON. If the pointer is nil, it serializes as "null".
func (p Ptr[T]) MarshalJSON() ([]byte, error) {
	if p.x == nil {
		return json.Marshal(p.x)
	}
	return json.Marshal(*p.x)
}

// UnmarshalXML unmarshals XML data into the Ptr. It handles character data and XML elements.
func (p *Ptr[T]) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error
	var token xml.Token
	var hasValue bool

	p.x = nil

	for {
		token, err = d.Token()
		if err != nil {
			return err
		}

		switch t := token.(type) {
		case xml.CharData:
			if len(t) == 0 {
				p.x = nil
				break
			}

			var value T
			err = xml.Unmarshal(t, &value)
			if err.Error() == "EOF" {
				err = json.Unmarshal(t, &value)
				if err != nil {
					return err
				}
			} else if err != nil {
				return err
			}

			hasValue = true
			p.x = &value
			break

		case xml.EndElement:
			if !hasValue {
				p.x = nil
			}
			return err
		}
	}
}

// MarshalXML marshals the Ptr into XML. If the pointer is nil, it serializes as a nil element.
func (p Ptr[T]) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if p.x == nil {
		return e.EncodeElement(nil, start)
	}
	return e.EncodeElement(*p.x, start)
}

// String returns a string representation of the Ptr, using the default formatting for its value.
func (s Ptr[T]) String() string {
	return fmt.Sprintf("%v", s.x)
}
