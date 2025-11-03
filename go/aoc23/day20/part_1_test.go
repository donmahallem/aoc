package day20_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day20"
)

var testDataSample1 string = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`
var testDataSample2 string = `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`

func TestParseInput(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		const expected int = 4
		reader := strings.NewReader(testDataSample1)
		if res := day20.ParseInput(reader); len(res.Modules) != expected {
			t.Errorf(`Expected %d to be %d`, res, expected)
		}

	})

}

func TestPart1(t *testing.T) {
	t.Run("test sample 1", func(t *testing.T) {
		expected := 32000000
		reader := strings.NewReader(testDataSample1)
		res := day20.Part1(reader)
		if res != expected {
			t.Errorf(`Expected number of blocks to be %d, got %d`, expected, res)
		}
	})
	t.Run("test sample 2", func(t *testing.T) {
		expected := 11687500
		reader := strings.NewReader(testDataSample2)
		res := day20.Part1(reader)
		if res != expected {
			t.Errorf(`Expected number of blocks to be %d, got %d`, expected, res)
		}
	})
}

func BenchmarkPart1(b *testing.B) {

	reader := strings.NewReader(testDataSample1)
	for b.Loop() {
		day20.Part1(reader)
		reader.Seek(0, 0)
	}
}
