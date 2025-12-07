package day02

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed sample.txt
var testData string

func Test_parseInputGen(t *testing.T) {
	t.Run("test parseInputGen", func(t *testing.T) {
		reader := strings.NewReader(testData)
		var result []intInterval
		parseInputGen(reader)(func(data intInterval) bool {
			result = append(result, data)
			return true
		})
		//11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
		//1698522-1698528,446443-446449,38593856-38593862,565653-565659,
		//824824821-824824827,2121212118-2121212124
		expectedValues := []intInterval{
			{Min: 11, Max: 22},
			{Min: 95, Max: 115},
			{Min: 998, Max: 1012},
			{Min: 1188511880, Max: 1188511890},
			{Min: 222220, Max: 222224},
			{Min: 1698522, Max: 1698528},
			{Min: 446443, Max: 446449},
			{Min: 38593856, Max: 38593862},
			{Min: 565653, Max: 565659},
			{Min: 824824821, Max: 824824827},
			{Min: 2121212118, Max: 2121212124},
		}
		if len(result) != len(expectedValues) {
			t.Fatalf("Unexpected number of parsed values: got %d, want %d", len(result), len(expectedValues))
		}
		for i := range result {
			if result[i] != expectedValues[i] {
				t.Errorf("Unexpected parsed value at index %d: got %+v, want %+v", i, result[i], expectedValues[i])
			}
		}

	})
}
