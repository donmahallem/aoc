package day20

import (
	_ "embed"
	"io"
	"strings"
	"testing"
)

//go:embed sample.txt
var testData string

var COUNT_CHEATS_TESTS [][2]int = [][2]int{{2, 14},
	{4, 14},
	{6, 2},
	{8, 4},
	{10, 2},
	{12, 3},
	{20, 1},
	{36, 1},
	{38, 1},
	{40, 1},
	{64, 1}}

func TestParseInput(t *testing.T) {
	raceWay, err := parseInput(strings.NewReader(testData))
	if err != nil {
		t.Errorf(`Unexpected error: %v`, err)
	}
	if len(raceWay) != 85 {
		t.Errorf(`Expected %d to match 85`, len(raceWay))
	}
}
func TestCountCheats(t *testing.T) {
	patterns, err := parseInput(strings.NewReader(testData))
	if err != nil {
		t.Errorf(`Unexpected error: %v`, err)
	}
	if result := countCheats(patterns, 2); result != 44 {
		t.Errorf(`Expected %d to save %d steps. Not %d`, 2, 44, result)
	}
}

func BenchmarkParseInput(b *testing.B) {
	data := strings.NewReader(testData)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		parseInput(data)
	}
}

func BenchmarkCountCheats(b *testing.B) {
	patterns, err := parseInput(strings.NewReader(testData))
	if err != nil {
		b.Errorf(`Unexpected error: %v`, err)
	}
	for b.Loop() {
		countCheats(patterns, 2)
	}
}

func FuzzPart1(f *testing.F) {
	f.Add(testData)
	f.Add("####\n#..#\n#..#\n####")
	f.Fuzz(func(t *testing.T, input string) {
		r := strings.NewReader(input)
		Part1(r)
	})
}
