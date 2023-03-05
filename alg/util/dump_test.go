package util

import (
	"os"
	"reflect"
	"testing"
)

func TestDump1dimSlice(t *testing.T) {
	testData := []int64{int64(1), int64(2), int64(3)}
	Dump(testData, "test.txt")
	var rv = Restore1dimSlice("test.txt")
	if !reflect.DeepEqual(testData, rv) {
		t.Error("Not equal")
	}
	os.Remove("test.txt")
}
