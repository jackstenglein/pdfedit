package pdfedit

// XrefTable represents a PDF document's cross-reference table.
type XrefTable struct{}

// Size returns the maximum object number contained within the xref table.
func (xref *XrefTable) Size() (int, error) {
	return 0, nil
}

// Object returns the PDF object with the given object number and generation number found within the
// xref table.
func (xref *XrefTable) Object(objectNum int, generationNum int) (*Object, error) {
	return nil, nil
}

// AddObject inserts the given object at the end of the xref table. The new object number of the
// object is returned.
func (xref *XrefTable) AddObject(object *Object) (int, error) {
	return 0, nil
}

// IsInUse returns true if the object number is marked as in use in the xref table.
func (xref *XrefTable) IsInUse(objectNum int) (bool, error) {
	return false, nil
}

// IsFree returns true if the object number is marked as free in the xref table.
func (xref *XrefTable) IsFree(objectNum int) (bool, error) {
	return false, nil
}

// ByteOffset returns the byte offset of the given object number if the object is in use and -1 if
// the object is free.
func (xref *XrefTable) ByteOffset(objectNum int) (int, error) {
	return 0, nil
}

// GenerationNumber returns the generation number of the given object number in the xref table.
func (xref *XrefTable) GenerationNumber(objectNum int) (int, error) {
	return 0, nil
}
