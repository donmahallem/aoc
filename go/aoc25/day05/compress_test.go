package day05

import (
	"reflect"
	"testing"
)

func Test_compressValidRanges(t *testing.T) {
	tests := []struct {
		name   string
		input  []validRange
		output []validRange
	}{
		{
			name: "overlapping ranges",
			input: []validRange{
				{Min: 1, Max: 5},
				{Min: 4, Max: 10},
				{Min: 12, Max: 15},
			},
			output: []validRange{
				{Min: 1, Max: 10},
				{Min: 12, Max: 15},
			},
		},
		{
			name: "adjacent ranges",
			input: []validRange{
				{Min: 1, Max: 5},
				{Min: 6, Max: 10},
				{Min: 12, Max: 15},
			},
			output: []validRange{
				{Min: 1, Max: 10},
				{Min: 12, Max: 15},
			},
		},
		{
			name: "disjoint ranges",
			input: []validRange{
				{Min: 1, Max: 5},
				{Min: 7, Max: 10},
				{Min: 12, Max: 15},
			},
			output: []validRange{
				{Min: 1, Max: 5},
				{Min: 7, Max: 10},
				{Min: 12, Max: 15},
			},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			compressValidRanges(&testCase.input)
			if !reflect.DeepEqual(testCase.input, testCase.output) {
				t.Errorf("Expected %+v, got %+v", testCase.output, testCase.input)
			}
		})
	}
}
