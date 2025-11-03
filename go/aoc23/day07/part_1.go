package day07

import (
	"bufio"
	"fmt"
	"io"
	"slices"

	"github.com/donmahallem/aoc/go/aoc_utils/bytes"
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
			} else if parsedInt, ok := bytes.ParseIntFromByte[int](ch); ok {
				game.Bid = (10 * game.Bid) + parsedInt
			}
		}
		slices.Sort(freq[:])
		tmp := freq[12]
		if tmp == 5 {
			game.Rating = 7
		} else if tmp == 4 {
			game.Rating = 6
		} else if tmp == 3 {
			if freq[11] == 2 {
				game.Rating = 5
			} else {
				game.Rating = 4
			}
		} else if freq[12] == 2 && freq[11] == 2 {
			game.Rating = 3
		} else if freq[12] == 2 && freq[11] == 1 {
			game.Rating = 2
		} else {
			game.Rating = 1
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
