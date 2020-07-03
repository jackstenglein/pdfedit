// Package pdfedit provides types and functions for parsing and editing PDF files.
package pdfedit

// Document represents a PDF file.
type Document struct{}

// NewDocument returns a Document representing the PDF at the given file path.
func NewDocument(filepath string) (*Document, error) {
	return nil, nil
}

// NewDocumentFromURL returns a Document representing the PDF at the given URL. The PDF will first
// be downloaded locally and then opened.
func NewDocumentFromURL(url string) (*Document, error) {
	return nil, nil
}

// Version returns the Document's version number as a string.
func (document *Document) Version() (string, error) {
	return "", nil
}

// Trailer returns the Document's trailer object.
func (document *Document) Trailer() (*Trailer, error) {
	return nil, nil
}

// Info returns the Document's information dictionary.
func (document *Document) Info() (*Info, error) {
	return nil, nil
}

// XrefTable returns the Document's xref table.
func (document *Document) XrefTable() (*XrefTable, error) {
	return nil, nil
}

// Catalog returns the Document's catalog object.
func (document *Document) Catalog() (*Catalog, error) {
	return nil, nil
}

// Object returns the PDF object with the given object number and generation number found within the
// Document.
func (document *Document) Object(objectNum int, generationNum int) (*Object, error) {
	return nil, nil
}

// AddObject inserts the given object at the end of the Document's xref table. The new object number
// of the object is returned.
func (document *Document) AddObject(object *Object) (int, error) {
	return 0, nil
}

// Validate checks the Document to ensure it meets the PDF specifications. If it does not, an error
// is returned describing the issue. Validate can be used to check whether any modifications made to
// the Document were incorrect.
func (document *Document) Validate() error {
	return nil
}

// Save overwrites the original version of the PDF on the file system with the version represented
// by the Document. If writing the Document fails, the original version of the PDF will be unchanged.
func (document *Document) Save() error {
	return nil
}

// SaveAs writes the Document to a PDF located at the given filepath. The original PDF file that the
// Document was first created from will be unchanged.
func (document *Document) SaveAs(filepath string) error {
	return nil
}
