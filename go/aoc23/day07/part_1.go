package day07

import (
	"bufio"
	"io"
	"slices"
	"sort"
)

type Game struct {
	Hand   [5]int8
	Bid    uint
	Rating int8
}

func parseInput(in io.Reader) []Game {
	s := bufio.NewScanner(in)
	games := make([]Game, 0, 16)
	for s.Scan() {
		data := s.Bytes()
		game := Game{Hand: [5]int8{0, 0, 0, 0, 0}, Bid: 0}
		for idx, ch := range data {
			if idx < 5 {
				switch ch {
				case 'T':
					game.Hand[idx] = 8
				case 'J':
					game.Hand[idx] = 9
				case 'Q':
					game.Hand[idx] = 10
				case 'K':
					game.Hand[idx] = 11
				case 'A':
					game.Hand[idx] = 12
				default:
					game.Hand[idx] = int8(ch - '2')
				}
			} else if ch >= '0' && ch <= '9' {
				game.Bid = (10 * game.Bid) + (uint(ch - '0'))
			}
		}
		games = append(games, game)
	}
	return games
}

func rateCards(games []Game) {
	var freq [13]int

	for idx := range games {
		for idx := range 13 {
			freq[idx] = 0
		}
		for _, card := range games[idx].Hand {
			freq[card]++
		}

		counts := []int{}
		for _, c := range freq {
			if c > 0 {
				counts = append(counts, c)
			}
		}
		sort.Slice(counts, func(i, j int) bool {
			return counts[i] > counts[j]
		})

		switch {
		case len(counts) == 1 && counts[0] == 5:
			games[idx].Rating = 7
		case len(counts) == 2 && counts[0] == 4:
			games[idx].Rating = 6
		case len(counts) == 2 && counts[0] == 3:
			games[idx].Rating = 5
		case len(counts) == 3 && counts[0] == 3:
			games[idx].Rating = 4
		case len(counts) == 3 && counts[0] == 2:
			games[idx].Rating = 3
		case len(counts) == 4:
			games[idx].Rating = 2
		default:
			games[idx].Rating = 1
		}
	}
}

func Part1(in io.Reader) uint {
	games := parseInput(in)
	rateCards(games)
	slices.SortFunc(games, func(a Game, b Game) int {
		if a.Rating == b.Rating {
			for i := range 5 {
				if a.Hand[i] == b.Hand[i] {
					continue
				}
				return int(a.Hand[i] - b.Hand[i])
			}
		}
		return int(a.Rating - b.Rating)
	})
	var total uint = 0
	for idx, game := range games {
		if game.Rating <= 1 {
			continue
		}
		total += uint(game.Bid * uint(idx+1))
	}
	return total
}
