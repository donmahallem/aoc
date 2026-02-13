package day05

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/test_utils"
)

const testData = "3-5\n10-14\n\n1\n5\n"

func Test_parseInput(t *testing.T) {
	t.Run("test parseInput", func(t *testing.T) {
		reader := strings.NewReader(testData)
		parsedData, err := parseInput(reader)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expectedRanges := []validRange{
			{Min: 3, Max: 5},
			{Min: 10, Max: 14},
		}
		if len(parsedData.validRanges) != len(expectedRanges) {
			t.Fatalf("expected %d ranges, got %d", len(expectedRanges), len(parsedData.validRanges))
		}
		for i := range expectedRanges {
			if parsedData.validRanges[i] != expectedRanges[i] {
				t.Errorf("range %d mismatch: got %+v want %+v", i, parsedData.validRanges[i], expectedRanges[i])
			}
		}
		expectedIngredients := []uint64{1, 5}
		if len(parsedData.ingredients) != len(expectedIngredients) {
			t.Fatalf("expected %d ingredients, got %d", len(expectedIngredients), len(parsedData.ingredients))
		}
		for i := range expectedIngredients {
			if parsedData.ingredients[i] != expectedIngredients[i] {
				t.Errorf("ingredient %d mismatch: got %d want %d", i, parsedData.ingredients[i], expectedIngredients[i])
			}
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
