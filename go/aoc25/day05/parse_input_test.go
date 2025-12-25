package day05

import (
	_ "embed"
	"reflect"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/test_utils"
)

//go:embed sample.txt
var testData string

func Test_parseInput(t *testing.T) {
	tests := taskData{
		validRanges: []validRange{
			{Min: 3, Max: 5},
			{Min: 10, Max: 14},
			{Min: 16, Max: 20},
			{Min: 12, Max: 18},
		},
		ingredients: []uint64{1, 5, 8, 11, 17, 32},
	}
	t.Run("test parseInput", func(t *testing.T) {
		reader := strings.NewReader(testData)
		parsedData, err := parseInput(reader)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(parsedData.validRanges, tests.validRanges) {
			t.Errorf("Expected validRanges to be %+v, got %+v", tests.validRanges, parsedData.validRanges)
		}
		if !reflect.DeepEqual(parsedData.ingredients, tests.ingredients) {
			t.Errorf("Expected ingredients to be %+v, got %+v", tests.ingredients, parsedData.ingredients)
		}
	})
}

func BenchmarkParseInput(b *testing.B) {

	b.Run("benchmark sample data", func(b *testing.B) {

		reader := strings.NewReader(testData)
		for b.Loop() {
			parseInput(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 25, 5, parseInput)
	})
}
