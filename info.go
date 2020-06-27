package pdfedit

// Info represents a PDF document's information dictionary.
type Info struct{}

// Title returns the document's title.
func (info *Info) Title() (string, error) {
	return "", nil
}

// Author returns the name of the person who created the document.
func (info *Info) Author() (string, error) {
	return "", nil
}

// Subject returns the subject of the document.
func (info *Info) Subject() (string, error) {
	return "", nil
}

// Keywords returns the keywords associated with the document.
func (info *Info) Keywords() (string, error) {
	return "", nil
}

// Creator returns the name of the application that created the original document if the PDF was
// converted from another format.
func (info *Info) Creator() (string, error) {
	return "", nil
}

// Producer returns the name of the application that converted the original document to PDF, if the
// original document was in another format.
func (info *Info) Producer() (string, error) {
	return "", nil
}

// CreationDate returns the date and time the document was created.
func (info *Info) CreationDate() (string, error) {
	return "", nil
}

// ModifiedDate returns the date and time the document was most recently modified.
func (info *Info) ModifiedDate() (string, error) {
	return "", nil
}
