package day03_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day03"
	"github.com/donmahallem/aoc/test_utils"
)

const testDataPart2 string = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func TestDoReader(t *testing.T) {
	t.Run("test irrelevant data", func(t *testing.T) {
		const sourceData = "{\"id\": 10, \"name\": \"Pie\"}"
		var data, _ = io.ReadAll(day03.NewDoReader(strings.NewReader(sourceData)))

		if i := strings.Compare(string(data), sourceData); i != 0 {
			t.Errorf(`Expected %s to match %s`, string(data), sourceData)
		}
	})
	t.Run("test striping don't() blocks", func(t *testing.T) {
		const sourceData = "asdfdo()yodon't()nodo()asdf"
		const targetData = "asdfyoasdf"
		var data, _ = io.ReadAll(day03.NewDoReader(strings.NewReader(sourceData)))

		if i := strings.Compare(string(data), targetData); i != 0 {
			t.Errorf(`Expected %s to match %s`, string(data), targetData)
		}
	})
}
func TestPart2(t *testing.T) {
	var data = day03.Part2(strings.NewReader(testDataPart2))
	expected := 48
	if data != expected {
		t.Errorf(`Expected %d to match %d`, data, expected)
	}
}

func BenchmarkPart2(b *testing.B) {
	if !test_utils.CheckTestDataExists(24, 3) {
		b.Skip("Couldn't retrieve test file data")
	}
	sourceData, _ := test_utils.GetTestData(24, 3)
	reader := strings.NewReader(sourceData)
	b.Run("test large input", func(b *testing.B) {
		for b.Loop() {
			reader.Seek(0, 0)
			day03.Part2(reader)
		}
	})
}
