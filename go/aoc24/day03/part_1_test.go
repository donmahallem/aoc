package day03_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day03"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName32(t *testing.T) {
	const sourceData = "mul(1,2)mul(1amul(2,3)"
	var data, _ = io.ReadAll(day03.NewMulReader(strings.NewReader(sourceData)))

	if i := strings.Compare(string(data), "8"); i != 0 {
		t.Errorf(`Expected %s to match %s`, string(data), "8")
	}
}
