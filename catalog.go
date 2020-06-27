package pdfedit

// Catalog represents a PDF document's catalog dictionary, which is the root of the document's
// object hierarchy.
type Catalog struct{}

// Type returns the type of PDF object represented by the catalog. This is always the string
// `Catalog`. Type never returns an error. Type is implemented as part of the Object interface.
func (catalog *Catalog) Type() (string, error) {
	return "Catalog", nil
}

// Children returns a list of objects referenced by the catalog. Children is implemented as part
// of the Object interface.
func (catalog *Catalog) Children() ([]*Object, error) {
	return nil, nil
}

// Version returns the version of the PDF specification to which the document conforms, if that
// version is later than the version specified in the file's header.
func (catalog *Catalog) Version() (string, error) {
	return "", nil
}
