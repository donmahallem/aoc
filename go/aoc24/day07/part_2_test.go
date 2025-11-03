package day07_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day07"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestOpConcat(t *testing.T) {
	if i := day07.OpConcat(5, 3); i != 53 {
		t.Errorf(`Expected %d and %d to be %d and not %d`, 5, 3, 53, i)
	}
	if i := day07.OpConcat(512, 355); i != 512355 {
		t.Errorf(`Expected %d and %d to be %d and not %d`, 512, 355, 512355, i)
	}
}

func TestPart2(t *testing.T) {
	if res := day07.Part2(strings.NewReader(testData)); res != 11387 {
		t.Errorf(`Expected %d to match %d`, res, 11387)
	}
}

func BenchmarkPart2(b *testing.B) {
	testData := strings.NewReader(testData)
	for b.Loop() {
		testData.Seek(0, io.SeekStart)
		day07.Part2(testData)
	}
}
