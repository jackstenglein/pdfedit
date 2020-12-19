package pdfedit

// XrefTable represents a PDF document's cross-reference table.
type XrefTable struct {
	size    int
	objects map[int]Object
	entries map[int]*xrefEntry
}

// xrefEntry represents an entry in the PDF document's cross-reference table.
type xrefEntry struct {
	inUse            bool
	objectNumber     int
	generationNumber int
	offset           int
}

// Size returns the maximum object number contained within the xref table.
func (xref *XrefTable) Size() int {
	return xref.size
}

// Object returns the PDF object with the given object number and generation number found within the
// xref table. If the object number is free or not contained within the XrefTable, nil is returned.
// If the object number is in use but has a different generation number, nil is returned.
func (xref *XrefTable) Object(objectNum int, generationNum int) Object {
	// obj := xref.objects[objectNum]
	// if obj != nil && obj.GenerationNumber() != generationNum {
	// 	return nil
	// }
	// return obj
	return nil
}

// AddObject inserts the given object at the end of the xref table. The new object number of the
// object is returned.
func (xref *XrefTable) AddObject(object *Object) (int, error) {
	return 0, nil
}

// IsInUse returns true if the object number is marked as in use in the xref table.
func (xref *XrefTable) IsInUse(objectNum int) bool {
	_, present := xref.objects[objectNum]
	return present
}

// IsFree returns true if the object number is marked as free in the xref table.
func (xref *XrefTable) IsFree(objectNum int) bool {
	return !xref.IsInUse(objectNum)
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

func (xref *XrefTable) addEntry(entry *xrefEntry) {
	if xref.entries == nil {
		xref.entries = make(map[int]*xrefEntry)
	}

	xref.entries[entry.objectNumber] = entry
}
