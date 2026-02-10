package day03

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/test_utils"
)

const testData = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

func TestNewMulReader(t *testing.T) {
	const sourceData = "mul(1,2)mul(1amul(2,3)"
	var data, _ = io.ReadAll(newMulReader(strings.NewReader(sourceData)))

	if i := strings.Compare(string(data), "8"); i != 0 {
		t.Errorf(`Expected %s to match %s`, string(data), "8")
	}
}

func TestPart1(t *testing.T) {
	t.Run("test sample 1", func(t *testing.T) {
		expected := 161
		reader := strings.NewReader(testDataPart2)
		result, err := Part1(reader)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != expected {
			t.Errorf(`Expected number of blocks to be %d, got %d`, expected, result)
		}
	})

	t.Run("test real data", func(t *testing.T) {
		expected := 179571322
		result, ok := test_utils.TestFullDataForDate(t, 24, 3, Part1)
		if !ok || result != expected {
			t.Errorf(`Expected %d to be %d`, result, expected)
		}
	})
}

func BenchmarkPart1(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {

		reader := strings.NewReader(testDataPart2)
		for b.Loop() {
			Part1(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 25, 7, Part1)
	})
}
