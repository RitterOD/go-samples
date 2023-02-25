package interface_exercises

import (
	"fmt"
	"os"
	"testing"
)

func TestWrapping(t *testing.T) {
	var testString = "Test string"
	var counter = 0
	var testWriter = Wrap(os.Stdout, &counter)
	fmt.Fprintf(testWriter, testString)
	if counter != len(testString) {

		var errorMsg = fmt.Sprintf("failed. Expected %d given:  %d", len(testString), counter)
		t.Error(errorMsg)
	}

}
