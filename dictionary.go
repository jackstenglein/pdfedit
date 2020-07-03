package pdfedit

// Dictionary represents a PDF dictionary object. It implements the Object interface.
type Dictionary struct{}

// Type returns the value of the Dictionary's /Type entry, if present. Otherwise, it returns the
// string `Dictionary`.
func (d *Dictionary) Type() (string, error) {
	return "", nil
}

// IsIndirect returns true if the dictionary is an indirect object.
func (d *Dictionary) IsIndirect() (bool, error) {
	return false, nil
}

// ObjectNumber returns the object number for the Dictionary if it is an indirect object and -1 if
// it is a direct object.
func (d *Dictionary) ObjectNumber() (int, error) {
	return -1, nil
}

// GenerationNumber returns the generation number for the Dictionary if it is an indirect object and
// -1 if it is a direct object.
func (d *Dictionary) GenerationNumber() (int, error) {
	return -1, nil
}

// Children returns the list of objects used as values in the Dictionary. It is equivalent to the
// Values function.
func (d *Dictionary) Children() ([]*Object, error) {
	return d.Values()
}

// Subtype returns the value of the Dictionary's /Subtype or /S entry, if present. Otherwise, it
// returns an empty string.
func (d *Dictionary) Subtype() (string, error) {
	return "", nil
}

// Size returns the number of entries in the Dictionary.
func (d *Dictionary) Size() (int, error) {
	return -1, nil
}

// Entries returns a map of the key-value pairs contained in the Dictionary.
func (d *Dictionary) Entries() (map[string]*Object, error) {
	return nil, nil
}

// Keys returns a list of the keys contained in the Dictionary.
func (d *Dictionary) Keys() ([]string, error) {
	return nil, nil
}

// Values returns a list of the objects used as values in the Dictionary.
func (d *Dictionary) Values() ([]*Object, error) {
	return nil, nil
}

// Entry returns the object associated in the Dictionary with the given key, or nil if the key is
// not present.
func (d *Dictionary) Entry(key string) (*Object, error) {
	return nil, nil
}

// SetEntry sets the key-value pair within the Dictionary to the provided values. If the key already
// exists, its value is overwritten. If the key does not exist, it is added.
func (d *Dictionary) SetEntry(key string, value *Object) error {
	return nil
}

// RemoveEntry removes the key and its value from the Dictionary and returns the value.
func (d *Dictionary) RemoveEntry(key string) (*Object, error) {
	return nil, nil
}

// IsStream returns true if the Dictionary contains stream data.
func (d *Dictionary) IsStream() (bool, error) {
	return false, nil
}

// StreamLength returns the length of the stream data if the Dictionary is a stream and -1
// otherwise.
func (d *Dictionary) StreamLength() (int, error) {
	return -1, nil
}

// StreamData returns the stream data as a list of bytes. If the Dictionary is not a stream,
// a nil list will be returned.
func (d *Dictionary) StreamData() ([]byte, error) {
	return nil, nil
}

// TODO: stream filters and modifying stream data
