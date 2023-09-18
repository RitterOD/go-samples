package generic

import (
	"slices"
	"testing"
)

func TestSerializationDeserialization(t *testing.T) {
	m := make(map[string]int)
	m["A"] = 1
	m["B"] = 2
	m["C"] = 3
	k := GetKeys[string, int](m)
	if slices.Contains(k, "A") {
		t.Error("Lose A value")
	}
	if !slices.Contains(k, "B") {
		t.Error("Lose B value")
	}
	if !slices.Contains(k, "C") {
		t.Error("Lose C value")
	}
}
