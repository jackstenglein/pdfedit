package pdfedit

// Trailer represents a PDF document's trailer object.
type Trailer struct{}

// Size returns the value of the trailer's /Size object, which is the number of objects in the PDF
// document's xref table.
func (trailer *Trailer) Size() (int, error) {
	return 0, nil
}

// Catalog returns the trailer's /Root object.
func (trailer *Trailer) Catalog() (*Catalog, error) {
	return nil, nil
}

// SetCatalog sets the trailer's /Root object. This can be used to change the PDF file's catalog
// object.
func (trailer *Trailer) SetCatalog(*Catalog) error {
	return nil
}

// Info returns the trailer's /Info object.
func (trailer *Trailer) Info() (*Info, error) {
	return nil, nil
}

// SetInfo sets the trailer's /Info object.
func (trailer *Trailer) SetInfo(*Info) error {
	return nil
}

// ID returns the value of /ID, which is a list of two strings that uniquely identifies the file
// within a workflow.
func (trailer *Trailer) ID() ([]string, error) {
	return nil, nil
}
