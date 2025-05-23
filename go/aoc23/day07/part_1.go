package day07

import (
	"bufio"
	"fmt"
	"io"
	"slices"
)

type Game struct {
	/*
		Hashing the hand is faster in comparison than iterating over the hand
		base13 encoding
	*/
	HandHash int
	Bid      int
	Rating   int8
}

var cardRanks = map[byte]uint8{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

func parseInput(in io.Reader) []Game {
	s := bufio.NewScanner(in)
	freq := [13]uint8{}
	games := make([]Game, 0, 16)
	hand := [5]uint8{0, 0, 0, 0, 0}
	for s.Scan() {
		for idx := range 13 {
			freq[idx] = 0
		}
		data := s.Bytes()
		game := Game{Bid: 0}
		for idx, ch := range data {
			if idx < 5 {
				if val, ok := cardRanks[ch]; ok {
					hand[idx] = val
					freq[val]++
				} else {
					panic(fmt.Sprintf("Unexpected character in input. Got: %b", ch))
				}
			} else if ch >= '0' && ch <= '9' {
				game.Bid = (10 * game.Bid) + (int(ch - '0'))
			}
		}
		var ones, twos, threes, fours, fives int
		for _, count := range freq {
			switch count {
			case 1:
				ones++
			case 2:
				twos++
			case 3:
				threes++
			case 4:
				fours++
			case 5:
				fives++
			}
		}

		switch {
		case fives == 1:
			game.Rating = 7 // Five of a Kind
		case fours == 1:
			game.Rating = 6 // Four of a Kind
		case threes == 1 && twos == 1:
			game.Rating = 5 // Full House
		case threes == 1 && ones == 2:
			game.Rating = 4 // Three of a Kind
		case twos == 2 && ones == 1:
			game.Rating = 3 // Two Pair
		case twos == 1 && ones == 3:
			game.Rating = 2 // One Pair
		default:
			game.Rating = 1 // High Card
		}
		game.HandHash = int(game.Rating)
		for i := range 5 {
			game.HandHash = (game.HandHash * 13) + int(hand[i])
		}
		games = append(games, game)
	}
	return games
}

func Part1(in io.Reader) int {
	games := parseInput(in)
	slices.SortFunc(games, func(a Game, b Game) int {
		return a.HandHash - b.HandHash
	})
	var total int = 0
	for idx, game := range games {
		total += game.Bid * int(idx+1)
	}
	return total
}
