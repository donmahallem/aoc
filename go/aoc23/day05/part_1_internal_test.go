package day05

import (
	_ "embed"
	"reflect"
	"slices"
	"strings"
	"testing"
)

//go:embed sample.txt
var testData string

func TestParseMap(t *testing.T) {
	testData := "123 321 292"
	testResult := parseMapping(&testData)
	expected := almanacRange{
		From: almanacInterval{Min: 321, Max: 321 + 292},
		To:   almanacInterval{Min: 123, Max: 123 + 292},
	}
	if testResult != expected {
		t.Errorf(`Expected %v to be %v`, testResult, expected)
	}
}
func TestParseSeeds(t *testing.T) {
	testData := "123 321 292"
	testResult := parseSeeds(testData)
	expected := []int{123, 321, 292}
	if !slices.Equal(testResult, expected) {
		t.Errorf(`Expected %v to be %v`, testResult, expected)
	}
}

func TestParseTranslateMap(t *testing.T) {
	testData := []int{50, 98, 2}
	testResult := make(map[int]int)
	translateMap(testData, &testResult)
	expected := make(map[int]int)
	expected[98] = 50
	expected[99] = 51
	if !reflect.DeepEqual(testResult, expected) {
		t.Errorf(`Expected %v to be %v`, testResult, expected)
	}
	testData = []int{4, 10, 3}
	expected[10] = 4
	expected[11] = 5
	expected[12] = 6
	translateMap(testData, &testResult)
	if !reflect.DeepEqual(testResult, expected) {
		t.Errorf(`Expected %v to be %v`, testResult, expected)
	}
}

func TestGetPosition(t *testing.T) {
	almanac := parseAlmanac(strings.NewReader(testData))
	tests := [][2]int{{79, 82}, {14, 43}, {55, 86}, {13, 35}}
	for _, test := range tests {
		if res := getPosition(almanac, test[0]); res != test[1] {
			t.Errorf(`Expected %d to be %d for Input %d`, res, test[1], test[0])
		}
	}
}

func TestParseMapping_Invalid(t *testing.T) {
	// too few tokens -> should not panic and should return zero value
	line := "123"
	mr := parseMapping(&line)
	if (mr != almanacRange{}) {
		t.Fatalf("expected zero AlmanacRange for invalid input, got %+v", mr)
	}
}

func TestParseAlmanac_NoSectionHeader(t *testing.T) {
	// mapping line before any section header should be ignored
	input := "1 2 3\n"
	a := parseAlmanac(strings.NewReader(input))
	// all sections should be empty
	if len(a.SeedToSoil) != 0 || len(a.SoilToFertilizer) != 0 || len(a.FertilizerToWater) != 0 || len(a.WaterToLight) != 0 || len(a.LightToTemperature) != 0 || len(a.TemperatureToHumidity) != 0 || len(a.HumidityToLocation) != 0 {
		t.Fatalf("expected empty sections, got %+v", a)
	}
}
