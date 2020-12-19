package pdfedit

import (
	"bufio"
	"errors"
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
func getXrefOffset(doc *Document) error {
	// Get the file size, which is therefore the max size of the xref offset
	fileInfo, err := doc.File.Stat()
	if err != nil {
		return err
	}
	maxLength := fileInfo.Size()
	doc.fileSize = maxLength

	// Seek to the correct position in the file
	offset := startXrefLength + eofLength + len(strconv.FormatInt(maxLength, 10)) + 1
	_, err = doc.File.Seek(int64(-offset), 2)
	if err != nil {
		return err
	}

	// Read the end of the file
	buffer := make([]byte, offset-eofLength)
	_, err = doc.File.Read(buffer)
	if err != nil {
		return err
	}

	// Extract the offset and convert it to an integer
	trailerString := string(buffer)
	startXrefIndex := strings.Index(trailerString, startXref)
	trailerString = trailerString[startXrefIndex+len(startXref):]
	endIndex := strings.Index(trailerString, "\n")
	offset, err = strconv.Atoi(trailerString[:endIndex])
	doc.xrefOffset = offset
	return err
}

func parseXrefTable(reader *bufio.Reader) (*XrefTable, error) {
	xrefTable := &XrefTable{}

	// Check that this is actually an xref table
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(line) != "xref" {
		return nil, errors.New("Xref offset does not point to xref keyword")
	}

	// Read xref table
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		isHeader, objNum, length := isXrefSubsectionHeader(line)
		if !isHeader {
			break
		}

		err = parseXrefSubsection(reader, xrefTable, objNum, length)
		if err != nil {
			return nil, err
		}
	}

	// TODO: read trailer

	return xrefTable, nil
}

func isXrefSubsectionHeader(line string) (bool, int, int) {
	line = strings.TrimSpace(line)
	tokens := strings.Split(line, " ")
	if len(tokens) != 2 {
		return false, -1, -1
	}

	objectNum, err := strconv.Atoi(tokens[0])
	if err != nil {
		return false, -1, -1
	}

	length, err := strconv.Atoi(tokens[1])
	if err != nil {
		return false, -1, -1
	}

	return true, objectNum, length
}

func parseXrefSubsection(reader *bufio.Reader, xrefTable *XrefTable, objectNum int, length int) error {

	for i := 0; i < length; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		tokens := strings.Split(strings.TrimSpace(line), " ")
		if len(tokens) != 3 {
			return errors.New("Incorrectly formatted xref line: " + line)
		}

		if tokens[2] != "n" {
			objectNum++
			continue
		}

		offset, err := strconv.Atoi(tokens[0])
		if err != nil {
			return err
		}

		generationNum, err := strconv.Atoi(tokens[1])
		if err != nil {
			return err
		}

		entry := &xrefEntry{
			inUse:            true,
			objectNumber:     objectNum,
			generationNumber: generationNum,
			offset:           offset,
		}
		xrefTable.addEntry(entry)
		objectNum++
	}

	return nil
}
