package day07

import (
	"bufio"
	"io"
	"slices"

	"github.com/donmahallem/aoc/go/aoc_utils"
	"github.com/donmahallem/aoc/go/aoc_utils/bytes"
)

type game struct {
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

func parseInput(in io.Reader) ([]game, error) {
	s := bufio.NewScanner(in)
	freq := [13]uint8{}
	games := make([]game, 0, 16)
	hand := [5]uint8{0, 0, 0, 0, 0}
	for s.Scan() {
		for idx := range freq {
			freq[idx] = 0
		}
		data := s.Bytes()
		g := game{Bid: 0}
		for idx, ch := range data {
			if idx < 5 {
				if val, ok := cardRanks[ch]; ok {
					hand[idx] = val
					freq[val]++
				} else {
					return nil, aoc_utils.NewUnexpectedInputError(ch)
				}
			} else if parsedInt, ok := bytes.ParseIntFromByte[int](ch); ok {
				g.Bid = (10 * g.Bid) + parsedInt
			}
		}
		slices.Sort(freq[:])
		tmp := freq[12]
		if tmp == 5 {
			g.Rating = 7
		} else if tmp == 4 {
			g.Rating = 6
		} else if tmp == 3 {
			if freq[11] == 2 {
				g.Rating = 5
			} else {
				g.Rating = 4
			}
		} else if freq[12] == 2 && freq[11] == 2 {
			g.Rating = 3
		} else if freq[12] == 2 && freq[11] == 1 {
			g.Rating = 2
		} else {
			g.Rating = 1
		}
		g.HandHash = int(g.Rating)
		for i := range hand {
			g.HandHash = (g.HandHash * 13) + int(hand[i])
		}
		games = append(games, g)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return games, nil
}

var cardRanksPart2 = map[byte]uint8{
	'J': 0,
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

func parseInput2(in io.Reader) ([]game, error) {
	s := bufio.NewScanner(in)
	freq := [13]uint8{}
	games := make([]game, 0, 16)
	hand := [5]uint8{0, 0, 0, 0, 0}
	for s.Scan() {
		for idx := range freq {
			freq[idx] = 0
		}
		data := s.Bytes()
		g := game{Bid: 0}
		var countJokers uint8 = 0
		for idx, ch := range data {
			if idx < 5 {
				if val, ok := cardRanksPart2[ch]; ok {
					hand[idx] = val
					if val == 0 {
						countJokers++
					} else {
						freq[val]++
					}
				} else {
					return nil, aoc_utils.NewUnexpectedInputError(ch)
				}
			} else if parsedInt, ok := bytes.ParseIntFromByte[int](ch); ok {
				g.Bid = (10 * g.Bid) + parsedInt
			}
		}
		slices.Sort(freq[:])
		tmp := freq[12] + countJokers
		if tmp == 5 {
			g.Rating = 7
		} else if tmp == 4 {
			g.Rating = 6
		} else if tmp == 3 {
			if freq[11] == 2 {
				g.Rating = 5
			} else {
				g.Rating = 4
			}
		} else if freq[12] == 2 && (countJokers == 1 || freq[11] == 2) {
			g.Rating = 3
		} else if freq[12] == 2 || countJokers == 1 {
			g.Rating = 2
		} else {
			g.Rating = 1
		}
		g.HandHash = int(g.Rating)
		for i := range 5 {
			g.HandHash = (g.HandHash * 13) + int(hand[i])
		}
		games = append(games, g)
	}
	return games, nil
}
