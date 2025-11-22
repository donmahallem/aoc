package day23

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/test_utils"
)

//go:embed sample.txt
var testData string

func Benchmark_parseInput(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {

		reader := strings.NewReader(testData)
		for b.Loop() {
			parseInput(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {

		full_data, ok := test_utils.GetTestData(23, 24)
		if ok {

			reader := strings.NewReader(full_data)
			for b.Loop() {
				parseInput(reader)
				reader.Seek(0, 0)
			}
		} else {
			b.Skip("No full data available")
			return
		}
	})
}

func Test_parseInput(t *testing.T) {
	points := parseInput(strings.NewReader(testData))
	var keyTa []byte = []byte{'t', 'a'}
	var hashTa NodeHash = HashId(keyTa)
	if res := len(points[hashTa]); res != 0 {
		t.Errorf(`Expected %v %d to have length %d not 0`, keyTa, hashTa, res)
	}
	keyTa = []byte{'c', 'o'}
	hashTa = HashId(keyTa)
	if res := len(points[hashTa]); res != 4 {
		t.Errorf(`Expected %v %d to have length %d not 4`, keyTa, hashTa, res)
	}
}
