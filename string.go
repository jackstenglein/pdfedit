package pdfedit

// String represents a PDF string object.
type String struct {
	objectNumber     int
	generationNumber int
	value            string
	isHexadecimal    bool
}

// Type returns the type of PDF object represented by the String. This is always the string
// `String`. Type never returns an error. Type is implemented as part of the Object interface.
func (s *String) Type() (string, error) {
	return "String", nil
}

// IsIndirect returns true if the String is an indirect object.
func (s *String) IsIndirect() (bool, error) {
	return s != nil && s.objectNumber > 0, nil
}

// ObjectNumber returns the object number for the String if it is an indirect object and -1 if it
// is a direct object. If the String is nil, 0 is returned.
func (s *String) ObjectNumber() (int, error) {
	if s == nil {
		return 0, nil
	}
	return s.objectNumber, nil
}

// GenerationNumber returns the generation number for the String if it is an indirect object and
// -1 if it is a direct object. If the String is nil, -1 is returned.
func (s *String) GenerationNumber() (int, error) {
	if s == nil {
		return -1, nil
	}
	return s.generationNumber, nil
}

// Children always returns a nil list because PDF strings cannot have children.
func (s *String) Children() ([]*Object, error) {
	return nil, nil
}

// Value returns the string value of the PDF String object.
func (s *String) Value() string {
	if s == nil {
		return ""
	}
	return s.value
}

// SetValue sets the String object to contain the provided value.
func (s *String) SetValue(value string) {
	if s != nil {
		s.value = value
	}
}

// IsHexadecimal returns true if the String object is represented in hexadecimal form in the PDF.
func (s *String) IsHexadecimal() bool {
	if s == nil {
		return false
	}
	return s.isHexadecimal
}

// SetIsHexadecimal sets whether the String object will be represented in hexadecimal form when the
// PDF file is saved.
func (s *String) SetIsHexadecimal(isHexadecimal bool) {
	if s != nil {
		s.isHexadecimal = isHexadecimal
	}
}
