package pdfedit

import (
	"bufio"
	"os"
	"reflect"
	"strings"
	"testing"
)

var getXrefOffsetTests = []struct {
	file   string
	offset int
}{
	{
		file:   "test_resources/cs380s_report.pdf",
		offset: 216,
	},
	{
		file:   "test_resources/icml2006.pdf",
		offset: 263752,
	},
}

func TestGetXrefOffset(t *testing.T) {
	for _, test := range getXrefOffsetTests {
		t.Run(test.file, func(t *testing.T) {
			file, err := os.Open(test.file)
			if err != nil {
				t.Error(err)
			}
			defer file.Close()
			doc := &Document{
				File: file,
			}

			err = getXrefOffset(doc)
			if err != nil {
				t.Error(err)
			}

			if doc.xrefOffset != test.offset {
				t.Errorf("Got incorrect offset: %d. Expected %d.", doc.xrefOffset, test.offset)
			}
		})
	}
}

var parseXrefTableTests = []struct {
	name      string
	table     string
	wantErr   bool
	wantTable *XrefTable
}{
	{
		name:    "MissingXrefKeyword",
		table:   "asdf\n",
		wantErr: true,
	},
	{
		name:    "IncorrectSubsectionFormat",
		table:   "xref\n0 6\nasdfasd 00000 n\n",
		wantErr: true,
	},
	{
		name:  "SingleSubsection",
		table: "xref\n0 1\n0000000003 00000 n\ntrailer\n",
		wantTable: &XrefTable{
			entries: map[int]*xrefEntry{
				0: &xrefEntry{
					inUse:            true,
					objectNumber:     0,
					generationNumber: 0,
					offset:           3,
				},
			},
		},
	},
	{
		name:  "MultipleSubsections",
		table: "xref\n0 1\n0000000000 00000 f\n3 1\n0000025325 00000 n\n23 2\n0000025518 00002 n\n0000025635 00000 n\n30 1\n0000025777 00000 n\ntrailer\n",
		wantTable: &XrefTable{
			entries: map[int]*xrefEntry{
				3: &xrefEntry{
					inUse:            true,
					objectNumber:     3,
					generationNumber: 0,
					offset:           25325,
				},
				23: &xrefEntry{
					inUse:            true,
					objectNumber:     23,
					generationNumber: 2,
					offset:           25518,
				},
				24: &xrefEntry{
					inUse:            true,
					objectNumber:     24,
					generationNumber: 0,
					offset:           25635,
				},
				30: &xrefEntry{
					inUse:            true,
					objectNumber:     30,
					generationNumber: 0,
					offset:           25777,
				},
			},
		},
	},
}

func TestParseXrefTable(t *testing.T) {
	for _, test := range parseXrefTableTests {
		t.Run(test.name, func(t *testing.T) {
			table, err := parseXrefTable(bufio.NewReader(strings.NewReader(test.table)))

			if (err != nil) != test.wantErr {
				if test.wantErr {
					t.Errorf("Expected an error but didn't get one.")
				} else {
					t.Errorf("Expected no error, but got `%v`", err)
				}
			}

			if err == nil && !reflect.DeepEqual(table, test.wantTable) {
				t.Errorf("Got table `%v`. Expected `%v`", table, test.wantTable)
			}
		})
	}
}

var isXrefSubsectionHeaderTests = []struct {
	name       string
	line       string
	wantResult bool
	wantObjNum int
	wantLength int
}{
	{
		name:       "IncorrectHeaderFormat",
		line:       "2 random\n",
		wantResult: false,
		wantObjNum: -1,
		wantLength: -1,
	},
	{
		name:       "CorrectHeaderFormat",
		line:       "28 5\n",
		wantResult: true,
		wantObjNum: 28,
		wantLength: 5,
	},
	{
		name:       "CorrectFormatWithWhiteSpace",
		line:       "29 13     \n",
		wantResult: true,
		wantObjNum: 29,
		wantLength: 13,
	},
}

func TestIsXrefSubsectionHeader(t *testing.T) {
	for _, test := range isXrefSubsectionHeaderTests {
		t.Run(test.name, func(t *testing.T) {
			result, objNum, length := isXrefSubsectionHeader(test.line)

			if result != test.wantResult {
				t.Errorf("Got result `%v`. Expected `%v`", result, test.wantResult)
			}

			if objNum != test.wantObjNum {
				t.Errorf("Got object number `%d`. Expected `%d`", objNum, test.wantObjNum)
			}

			if length != test.wantLength {
				t.Errorf("Got length `%d`. Expected `%d`", length, test.wantLength)
			}
		})
	}
}

var parseXrefSubsectionTests = []struct {
	name       string
	subsection string
	objectNum  int
	length     int
	wantErr    bool
	wantMap    map[int]*xrefEntry
}{
	{
		name:       "IncorrectLength",
		subsection: "Test string",
		length:     3,
		wantErr:    true,
	},
	{
		name:       "IncorrectEntry",
		subsection: "avcasdfadsf\n",
		length:     1,
		wantErr:    true,
	},
	{
		name:       "IncorrectOffset",
		subsection: "nnnnnnnnnn 00013 n\n0000000023 00042 f\n",
		length:     2,
		wantErr:    true,
	},
	{
		name:       "IncorrectGenNum",
		subsection: "0000000024 ggggg n\n0000000023 00042 f\n",
		length:     2,
		wantErr:    true,
	},
	{
		name:       "CorrectSubsection",
		subsection: "0000000003 65535 f\n0000000017 00000 n\n0000000081 00000 n\n0000000000 00007 f\n0000000331 00005 n\n",
		objectNum:  13,
		length:     5,
		wantErr:    false,
		wantMap: map[int]*xrefEntry{
			14: &xrefEntry{
				inUse:            true,
				objectNumber:     14,
				generationNumber: 0,
				offset:           17,
			},
			15: &xrefEntry{
				inUse:            true,
				objectNumber:     15,
				generationNumber: 0,
				offset:           81,
			},
			17: &xrefEntry{
				inUse:            true,
				objectNumber:     17,
				generationNumber: 5,
				offset:           331,
			},
		},
	},
}

func TestParseXrefSubsection(t *testing.T) {
	for _, test := range parseXrefSubsectionTests {
		t.Run(test.name, func(t *testing.T) {
			xrefTable := &XrefTable{}
			reader := bufio.NewReader(strings.NewReader(test.subsection))
			err := parseXrefSubsection(reader, xrefTable, test.objectNum, test.length)

			if (err != nil) != test.wantErr {
				if test.wantErr {
					t.Errorf("Expected an error but didn't get one.")
				} else {
					t.Errorf("Expected no error, but got `%v`", err)
				}
			}

			if err == nil && !reflect.DeepEqual(xrefTable.entries, test.wantMap) {
				t.Errorf("Got entry map `%v`. Expected `%v`", xrefTable.entries, test.wantMap)
			}
		})
	}
}
