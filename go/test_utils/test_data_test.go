package test_utils

import (
	"testing"
)

func TestLoadDefaultTestData(t *testing.T) {
	data := LoadDefaultTestData(t)
	cases := data[23][1]
	if _, ok := cases["sample1"]; !ok {
		t.Fatalf("expected case %q for 23/01, got %v", "sample1", cases)
	}
}
