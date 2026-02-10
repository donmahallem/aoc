package day23

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

const (
	TILE_EMPTY       byte = '.'
	TILE_FOREST      byte = '#'
	TILE_SLOPE_RIGHT byte = '>'
	TILE_SLOPE_LEFT  byte = '<'
	TILE_SLOPE_UP    byte = '^'
	TILE_SLOPE_DOWN  byte = 'v'
)

type direction = byte

const (
	DIR_UP    direction = '^'
	DIR_DOWN  direction = 'v'
	DIR_LEFT  direction = '<'
	DIR_RIGHT direction = '>'
)

/*
	encodes a cell

encoding is as follows (bits):
bit 0-3 - exit directions (up, down, left, right)
bit 4-7 - entry directions from (up, down, left, right)
*/
type cell = uint8

const (
	dirMaskNone    cell = 0
	dirExitTop     cell = 1 << 0
	dirExitBottom  cell = 1 << 1
	dirExitLeft    cell = 1 << 2
	dirExitRight   cell = 1 << 3
	dirExitAll     cell = dirExitTop | dirExitBottom | dirExitLeft | dirExitRight
	dirEntryTop    cell = 1 << 4
	dirEntryBottom cell = 1 << 5
	dirEntryLeft   cell = 1 << 6
	dirEntryRight  cell = 1 << 7
	dirEntryAll    cell = dirEntryTop | dirEntryBottom | dirEntryLeft | dirEntryRight
)

// parses the input and encodes exit nodes directly into the byte slice
// respectSlop indicates whether slope directions should be respected
func parseInput(in io.Reader, respectSlop bool) ([]cell, int, int, error) {
	scanner := bufio.NewScanner(in)

	width := 0

	var parsedData []cell
	var previousLine []byte
	height := 0
	for ; scanner.Scan(); height++ {
		raw := scanner.Bytes()
		// ensure consistent width
		if height == 0 {
			width = len(raw)
			parsedData = make([]cell, 0, width*width)
			previousLine = make([]byte, width)
		} else {
			if len(raw) != width {
				if len(raw) == 0 {
					return nil, 0, 0, aoc_utils.NewUnexpectedInputError(0)
				}
				return nil, 0, 0, aoc_utils.NewUnexpectedInputError(raw[0])
			}
		}
		for c := 0; c < width; c++ {
			tile := raw[c]
			if tile != TILE_EMPTY && tile != TILE_FOREST && tile != TILE_SLOPE_RIGHT && tile != TILE_SLOPE_LEFT &&
				tile != TILE_SLOPE_UP && tile != TILE_SLOPE_DOWN {
				return nil, 0, 0, aoc_utils.NewUnexpectedInputError(tile)
			}
			var currentMask cell = 0
			currentIndex := height*width + c
			previousIndex := currentIndex - 1
			if respectSlop {
				if c < width-1 && (tile == TILE_SLOPE_RIGHT || tile == TILE_EMPTY) {
					currentMask |= dirExitRight
				}
				if c > 0 {
					if tile == TILE_FOREST {
						parsedData[previousIndex] &^= dirExitRight
					} else if parsedData[previousIndex]&dirExitRight != 0 {
						currentMask |= dirEntryLeft
					}
					if raw[c-1] != TILE_FOREST && (tile == TILE_SLOPE_LEFT || tile == TILE_EMPTY) {
						currentMask |= dirExitLeft
						parsedData[previousIndex] |= dirEntryRight
					}
					if tile == TILE_SLOPE_RIGHT || tile == TILE_EMPTY {
						currentMask |= dirExitRight
					}
				}
				if height == 0 && (tile == TILE_SLOPE_DOWN || tile == TILE_EMPTY) {
					currentMask |= dirExitBottom
				} else if height > 0 {
					if tile == TILE_FOREST {
						parsedData[currentIndex-width] &^= dirExitBottom
					} else if parsedData[currentIndex-width]&dirExitBottom != 0 {
						currentMask |= dirEntryTop
					}
					if previousLine[c] != TILE_FOREST && (tile == TILE_SLOPE_UP || tile == TILE_EMPTY) {
						currentMask |= dirExitTop
						parsedData[currentIndex-width] |= dirEntryBottom
					}
					if tile == TILE_SLOPE_DOWN || tile == TILE_EMPTY {
						currentMask |= dirExitBottom
					}
				}
			} else {
				if tile == TILE_FOREST {
					if c > 0 {
						parsedData[previousIndex] &^= dirExitRight
						parsedData[previousIndex] &^= dirEntryRight
					}
					if height > 0 {
						parsedData[currentIndex-width] &^= dirExitBottom
						parsedData[currentIndex-width] &^= dirEntryBottom
					}
				} else {
					if c == 0 {
						currentMask |= dirExitRight
					} else {
						if parsedData[previousIndex]&dirExitRight != 0 {
							currentMask |= dirEntryLeft | dirExitLeft
							parsedData[previousIndex] |= dirEntryRight
						}
						currentMask |= dirExitRight
					}

					if height == 0 {
						currentMask |= dirExitBottom
					} else {
						if parsedData[currentIndex-width]&dirExitBottom != 0 {
							currentMask |= dirEntryTop | dirExitTop
							parsedData[currentIndex-width] |= dirEntryBottom
						}
						currentMask |= dirExitBottom
					}
				}
			}
			if c == width-1 {
				currentMask &^= dirExitRight
			}
			parsedData = append(parsedData, currentMask)
		}
		copy(previousLine, raw)
	}
	// fix last row
	if height > 0 {
		for c := 0; c < width; c++ {
			parsedData[(height-1)*width+c] &^= dirExitBottom
			parsedData[(height-1)*width+c] &^= dirEntryBottom
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, 0, 0, err
	}
	return parsedData, width, height, nil
}
