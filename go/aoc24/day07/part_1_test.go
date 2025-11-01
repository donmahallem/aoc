package day07_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day07"
)

const testData string = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestCheckLinePart1(t *testing.T) {
	result := 190
	terms := []int{10, 19}
	testData := day07.ParsedLineData{
		Result:     result,
		TestValues: terms,
	}
	if i := day07.CheckLinePart1(&testData); !i {
		t.Errorf(`Expected %v to be %d and not %t`, terms, result, i)
	}
}

func TestParseLine(t *testing.T) {
	rawLine := []byte("190: 10 19")
	expectedResult := 190
	expectedTerms := []int{10, 19}
	parsedLine := &day07.ParsedLineData{}
	day07.ParseLine(rawLine, parsedLine)
	if parsedLine.Result != expectedResult {
		t.Errorf(`Expected result to be %d, got %d`, expectedResult, parsedLine.Result)
	}
	if len(parsedLine.TestValues) != len(expectedTerms) {
		t.Errorf(`Expected number of terms to be %d, got %d`, len(expectedTerms), len(parsedLine.TestValues))
	}
	for i, term := range expectedTerms {
		if parsedLine.TestValues[i] != term {
			t.Errorf(`Expected term %d to be %d, got %d`, i, term, parsedLine.TestValues[i])
		}
	}
}

func TestPart1(t *testing.T) {
	if res := day07.Part1(strings.NewReader(testData)); res != 3749 {
		t.Errorf(`Expected %d to match %d`, res, 3749)
	}
}
func BenchmarkPart1(b *testing.B) {
	testData := strings.NewReader(testData)
	for b.Loop() {
		testData.Seek(0, io.SeekStart)
		day07.Part1(testData)
	}
}
