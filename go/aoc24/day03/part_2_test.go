package day03_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day03"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestNewDoReader(t *testing.T) {
	const sourceData = "{\"id\": 10, \"name\": \"Pie\"}"
	var data, _ = io.ReadAll(day03.NewDoReader(strings.NewReader(sourceData)))

	if i := strings.Compare(string(data), sourceData); i != 0 {
		t.Errorf(`Expected %s to match %s`, string(data), sourceData)
	}
}
func TestNewDoReader2(t *testing.T) {
	const sourceData = "asdfdo()yodon't()nodo()asdf"
	const targetData = "asdfyoasdf"
	var data, _ = io.ReadAll(day03.NewDoReader(strings.NewReader(sourceData)))

	if i := strings.Compare(string(data), targetData); i != 0 {
		t.Errorf(`Expected %s to match %s`, string(data), targetData)
	}
}

func TestPart2(t *testing.T) {
	const sourceData = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	var data = day03.Part2(strings.NewReader(sourceData))
	expected := 48
	if data != expected {
		t.Errorf(`Expected %d to match %d`, data, expected)
	}
}
