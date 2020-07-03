package pdfedit

// Array represents a PDF array object. It implements the object interface.
type Array struct{}

// Type returns the type of PDF object represented by the Array. This is always the string
// `Array`. Type never returns an error. Type is implemented as part of the Object interface.
func (a *Array) Type() (string, error) {
	return "Array", nil
}

// IsIndirect returns true if the Array is an indirect object.
func (a *Array) IsIndirect() (bool, error) {
	return false, nil
}

// ObjectNumber returns the object number for the Array if it is an indirect object and -1 if it
// is a direct object.
func (a *Array) ObjectNumber() (int, error) {
	return -1, nil
}

// GenerationNumber returns the generation number for the Array if it is an indirect object and
// -1 if it is a direct object.
func (a *Array) GenerationNumber() (int, error) {
	return -1, nil
}

// Children returns the underlying list of PDF objects represented by the Array.
func (a *Array) Children() ([]*Object, error) {
	return nil, nil
}

// Length returns the number of objects in the Array.
func (a *Array) Length() (int, error) {
	return -1, nil
}

// AddFront inserts the given object at the front of the Array.
func (a *Array) AddFront(object *Object) error {
	return a.Insert(0, object)
}

// AddBack inserts the given object at the end of the Array.
func (a *Array) AddBack(object *Object) error {
	len, _ := a.Length()
	return a.Insert(len, object)
}

// Insert inserts the given object at the given position in the Array. All objects currently at that
// position or greater are shifted to the right by one.
func (a *Array) Insert(position int, object *Object) error {
	return nil
}

// Replace switches the object at the given position in the Array with the given object. No other
// objects in the Array are changed.
func (a *Array) Replace(position int, object *Object) error {
	return nil
}

// Remove removes the object at the given position from the Array. All objects to the right of the
// position are shifted to the left by one.
func (a *Array) Remove(position int) error {
	return nil
}
