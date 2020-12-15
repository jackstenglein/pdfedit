package pdfedit

import (
	"os"
	"strconv"
	"strings"
)

const (
	startXref       = "startxref\n"
	startXrefLength = len(startXref)
	eofLength       = len("%%EOF\n")
)

func parseDocument(filepath string) error {

	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

// getXrefOffset returns the offset of the xref table or xref stream as
// specified in the given PDF file's trailer.
func getXrefOffset(file *os.File) (int, error) {
	// Get the file size, which is therefore the max size of the xref offset
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}
	maxLength := fileInfo.Size()

	// Seek to the correct position in the file
	offset := startXrefLength + eofLength + len(strconv.FormatInt(maxLength, 10)) + 1
	_, err = file.Seek(int64(-offset), 2)
	if err != nil {
		return 0, err
	}

	// Read the end of the file
	buffer := make([]byte, offset-eofLength)
	_, err = file.Read(buffer)
	if err != nil {
		return 0, err
	}

	// Extract the offset and convert it to an integer
	trailerString := string(buffer)
	startXrefIndex := strings.Index(trailerString, startXref)
	trailerString = trailerString[startXrefIndex+len(startXref):]
	endIndex := strings.Index(trailerString, "\n")
	return strconv.Atoi(trailerString[:endIndex])
}
