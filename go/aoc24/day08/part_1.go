package day08

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type point aoc_utils.Point[int16]
type inputData struct {
	Antennas map[byte][]point
	Width    int16
	Height   int16
}

func readSource(reader io.Reader) (*inputData, error) {
	data := make(map[byte][]point, 8)
	s := bufio.NewScanner(reader)
	y := int16(0)
	width := int16(0)
	for s.Scan() {
		line := s.Bytes()
		if width == 0 {
			width = int16(len(line))
		} else if width != int16(len(line)) {
			return nil, aoc_utils.NewParseError("line length is uneven", nil)
		}
		for idx, character := range line {
			if character == '.' {
				continue
			}
			if charData, ok := data[character]; ok {
				data[character] = append(charData, point{Y: y, X: int16(idx)})
			} else {
				data[character] = []point{{Y: y, X: int16(idx)}}
			}
		}
		y++
	}
	return &inputData{
		Antennas: data,
		Width:    width,
		Height:   y,
	}, nil
}

func OutOfBounds(x int16, y int16, width int16, height int16) bool {
	return x < 0 || y < 0 || x >= width || y >= height
}

func Part1(in io.Reader) (int, error) {
	input, err := readSource(in)
	if err != nil {
		return 0, err
	}
	var antennaListLen int
	var newPoint point
	echos := make(map[point]bool, 0)
	for antenna := range input.Antennas {
		antennaList := input.Antennas[antenna]
		antennaListLen = len(antennaList)
		for i := 0; i < antennaListLen; i++ {
			for j := 0; j < antennaListLen; j++ {
				if i == j {
					continue
				}
				newPoint.X = 2*antennaList[i].X - antennaList[j].X
				newPoint.Y = 2*antennaList[i].Y - antennaList[j].Y
				if !OutOfBounds(newPoint.X, newPoint.Y, input.Width, input.Height) {
					echos[newPoint] = true
				}
			}
		}
	}
	return len(echos), nil
}
