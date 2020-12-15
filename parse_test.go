package pdfedit

import (
	"os"
	"testing"
)

func TestGetXrefOffset(t *testing.T) {
	file, err := os.Open("test_resources/cs380s_report.pdf")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	offset, err := getXrefOffset(file)
	if err != nil {
		t.Error(err)
	}

	if offset != 216 {
		t.Errorf("Got incorrect offset: %d. Expected 216.", offset)
	}
}
