package pdfedit

// String represents a PDF string object.
type String struct{}

// Type returns the type of PDF object represented by the String. This is always the string
// `String`. Type never returns an error. Type is implemented as part of the Object interface.
func (s *String) Type() (string, error) {
	return "String", nil
}

// IsIndirect returns true if the String is an indirect object.
func (s *String) IsIndirect() (bool, error) {
	return false, nil
}

// ObjectNumber returns the object number for the String if it is an indirect object and -1 if it
// is a direct object.
func (s *String) ObjectNumber() (int, error) {
	return -1, nil
}

// GenerationNumber returns the generation number for the String if it is an indirect object and
// -1 if it is a direct object.
func (s *String) GenerationNumber() (int, error) {
	return -1, nil
}

// Children always returns a nil list because PDF strings cannot have children. Children never
// returns an error.
func (s *String) Children() ([]*Object, error) {
	return nil, nil
}

// Value returns the string value of the PDF String object.
func (s *String) Value() (string, error) {
	return "", nil
}

// SetValue sets the String object to contain the provided value.
func (s *String) SetValue(value string) error {
	return nil
}

// IsHexadecimal returns true if the String object is represented in hexadecimal form in the PDF.
func (s *String) IsHexadecimal() (bool, error) {
	return false, nil
}

// SetIsHexadecimal sets whether the String object will be represented in hexadecimal form when the
// PDF file is saved.
func (s *String) SetIsHexadecimal(isHexadecimal bool) error {
	return nil
}
