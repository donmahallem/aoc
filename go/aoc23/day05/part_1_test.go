package day05_test

import (
	"io"
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day05"
)

const testData string = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func TestParseMap(t *testing.T) {
	testData := "123 321 292"
	testResult := day05.ParseMapping(&testData)
	expected := day05.AlmanacRange{
		From: day05.Interval{Start: 321, End: 321 + 292},
		To:   day05.Interval{Start: 123, End: 123 + 292},
	}
	if testResult != expected {
		t.Errorf(`Expected %d to be %d`, testResult, expected)
	}
}
func TestParseSeeds(t *testing.T) {
	testData := "123 321 292"
	testResult := day05.ParseSeeds(testData)
	expected := []int{123, 321, 292}
	if !slices.Equal(testResult, expected) {
		t.Errorf(`Expected %d to be %d`, testResult, expected)
	}
}

func TestParseTranslateMap(t *testing.T) {
	testData := []int{50, 98, 2}
	testResult := make(map[int]int)
	day05.TranslateMap(testData, &testResult)
	expected := make(map[int]int)
	expected[98] = 50
	expected[99] = 51
	if !reflect.DeepEqual(testResult, expected) {
		t.Errorf(`Expected %d to be %d`, testResult, expected)
	}
	testData = []int{4, 10, 3}
	expected[10] = 4
	expected[11] = 5
	expected[12] = 6
	day05.TranslateMap(testData, &testResult)
	if !reflect.DeepEqual(testResult, expected) {
		t.Errorf(`Expected %d to be %d`, testResult, expected)
	}
}

func TestGetPosition(t *testing.T) {
	almanac := day05.ParseAlmanac(strings.NewReader(testData))
	tests := [][2]int{{79, 82}, {14, 43}, {55, 86}, {13, 35}}
	for _, test := range tests {
		if res := day05.GetPosition(almanac, test[0]); res != test[1] {
			t.Errorf(`Expected %d to be %d for Input %d`, res, test[1], test[0])
		}
	}
}

func TestPart1(t *testing.T) {
	result := day05.Part1(strings.NewReader(testData))
	if result != 35 {
		t.Errorf(`Expected %d to be %d`, result, 35)
	}
}

func BenchmarkPart1(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day05.Part1(reader)
	}
}
