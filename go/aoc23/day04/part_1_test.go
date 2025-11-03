package day04_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day04"
)

const testData string = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

func TestParseLine(t *testing.T) {
	ticket, winners, picks := day04.ParseLine([]byte("Card 153: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"))
	if len(winners) != 5 {
		t.Errorf(`Expected winners to have a length of %d. Not %d`, 5, len(winners))
	}
	if len(picks) != 8 {
		t.Errorf(`Expected picks to have a length of %d. Not %d`, 8, len(picks))
	}
	if ticket != 153 {
		t.Errorf(`Expected ticket to be %d. Not %d`, 153, ticket)
	}
}

func TestGetScore(t *testing.T) {
	_, winners, picks := day04.ParseLine([]byte("Card 153: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"))
	if score := day04.GetScore(winners, picks); score != 8 {
		t.Errorf(`Expected winners to have a length of %d. Not %d`, 8, score)
	}
}

func TestPart1(t *testing.T) {
	result := day04.Part1(strings.NewReader(testData))
	if result != 13 {
		t.Errorf(`Expected winners to have a length of %d. Not %d`, 13, result)
	}
}

func BenchmarkPart1(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day04.Part1(reader)
	}
}
