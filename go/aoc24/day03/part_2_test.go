package day03

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/test_utils"
)

const testDataPart2 string = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func TestDoReader(t *testing.T) {
	t.Run("test irrelevant data", func(t *testing.T) {
		const sourceData = "{\"id\": 10, \"name\": \"Pie\"}"
		var data, _ = io.ReadAll(NewDoReader(strings.NewReader(sourceData)))

		if i := strings.Compare(string(data), sourceData); i != 0 {
			t.Errorf(`Expected %s to match %s`, string(data), sourceData)
		}
	})
	t.Run("test striping don't() blocks", func(t *testing.T) {
		const sourceData = "asdfdo()yodon't()nodo()asdf"
		const targetData = "asdfyoasdf"
		var data, _ = io.ReadAll(NewDoReader(strings.NewReader(sourceData)))

		if i := strings.Compare(string(data), targetData); i != 0 {
			t.Errorf(`Expected %s to match %s`, string(data), targetData)
		}
	})
}

func TestPart2(t *testing.T) {
	t.Run("test sample 1", func(t *testing.T) {
		expected := 48
		reader := strings.NewReader(testDataPart2)
		result, err := Part2(reader)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != expected {
			t.Errorf(`Expected number of blocks to be %d, got %d`, expected, result)
		}
	})

	t.Run("test real data", func(t *testing.T) {
		expected := 103811193
		result, ok := test_utils.TestFullDataForDate(t, 24, 3, Part2)
		if !ok || result != expected {
			t.Errorf(`Expected %d to be %d`, result, expected)
		}
	})
}

func BenchmarkPart2(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {

		reader := strings.NewReader(testDataPart2)
		for b.Loop() {
			Part2(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 24, 3, Part2)
	})
}
