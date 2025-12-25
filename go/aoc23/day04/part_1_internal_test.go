package day04

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	ticket, winners, picks := parseLine([]byte("Card 153: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"))
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
	_, winners, picks := parseLine([]byte("Card 153: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"))
	if score := getScore(winners, picks); score != 8 {
		t.Errorf(`Expected winners to have a length of %d. Not %d`, 8, score)
	}
}
