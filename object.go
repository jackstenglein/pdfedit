package pdfedit

// Object is the interface that wraps methods common to all types of PDF objects.
// PDF supports eight basic types of objects:
//    - Boolean values
//    - Integer and real numbers
//    - Strings
//    - Names
//    - Arrays
//    - Dictionaries
//    - Streams
//    - The null object
type Object interface {
	// Type returns the type of this object.
	Type() (string, error)

	// IsIndirect returns true if the object is labeled as an indirect object.
	IsIndirect() (bool, error)

	// ObjectNumber returns the object number for an indirect object and -1 for a direct object.
	ObjectNumber() (int, error)

	// GenerationNumber returns the generation number for an indirect object and -1 for a direct
	// object.
	GenerationNumber() (int, error)

	// Children returns a list of objects referenced by this object.
	Children() ([]*Object, error)
}
