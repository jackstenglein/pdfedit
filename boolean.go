package pdfedit

// Boolean represents a PDF boolean object. It implements the object interface. The PDF object can
// be direct or indirect.
type Boolean struct {
	objectNumber     int
	generationNumber int
	value            bool
}

// Type returns the type of PDF object represented by the Boolean. This is always the string
// `Boolean`. Type never returns an error. Type is implemented as part of the Object interface.
func (b *Boolean) Type() (string, error) {
	return "Boolean", nil
}

// IsIndirect returns true if the Boolean is an indirect object.
func (b *Boolean) IsIndirect() (bool, error) {
	return b != nil && b.objectNumber > 0, nil
}

// ObjectNumber returns the object number for the Boolean if it is an indirect object and -1 if it
// is a direct object. If the Boolean is nil, 0 is returned.
func (b *Boolean) ObjectNumber() (int, error) {
	if b == nil {
		return 0, nil
	}
	return b.objectNumber, nil
}

// GenerationNumber returns the generation number for the Boolean if it is an indirect object and
// -1 if it is a direct object. If the Boolean is nil, -1 is returned.
func (b *Boolean) GenerationNumber() (int, error) {
	if b == nil {
		return -1, nil
	}
	return b.generationNumber, nil
}

// Children always returns a nil list because PDF booleans cannot have children. Children never
// returns an error.
func (b *Boolean) Children() ([]*Object, error) {
	return nil, nil
}

// Value returns the bool value of the PDF Boolean object.
func (b *Boolean) Value() bool {
	if b == nil {
		return false
	}
	return b.value
}

// SetValue updates the value of the Boolean object to the provided value.
func (b *Boolean) SetValue(value bool) {
	if b != nil {
		b.value = value
	}
}
