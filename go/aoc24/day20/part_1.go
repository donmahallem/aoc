package day20

import (
	"bufio"
	"io"

	"github.com/donmahallem/aoc/go/aoc_utils"
)

type point = aoc_utils.Point[int]
type pointList = []point

var dirUp point = point{X: 0, Y: -1}
var dirLeft point = point{X: -1, Y: 0}
var dirRight point = point{X: 1, Y: 0}
var dirDown point = point{X: 0, Y: 1}
var DIRS_ALL []*point = []*point{&dirDown, &dirUp, &dirLeft, &dirRight}

func parseInput(in io.Reader) (pointList, error) {
	var startPoint, endPoint point
	pList := make([]point, 0)
	s := bufio.NewScanner(in)
	y := 0
	blockWidth := -1
	for s.Scan() {
		line := s.Bytes()
		if blockWidth <= 0 {
			blockWidth = len(line)
		} else if len(line) != blockWidth {
			return nil, aoc_utils.NewParseError("Inconsistent line sizes", nil)
		}
		for x := range len(line) {
			switch line[x] {
			case '.':
				pList = append(pList, point{X: x, Y: y})
			case 'S':
				startPoint = point{X: x, Y: y}
			case 'E':
				endPoint = point{X: x, Y: y}
			case '#':
				//call
				continue
			default:
				return nil, aoc_utils.NewParseError("Invalid character in input", nil)
			}
		}
		y++
	}
	if scanner := s.Err(); scanner != nil {
		return nil, scanner
	}
	if y <= 3 || blockWidth <= 3 {
		return nil, aoc_utils.NewParseError("Input too small", nil)
	}
	// Append endPoint twice
	// Prepend start to list without array duplications
	pList = append(pList, endPoint, endPoint)
	copy(pList[1:], pList)
	pList[0] = startPoint
	// Iterate and find non visited neighbours
	for i := 1; i < len(pList)-2; i++ {
		for j := i + 1; j < len(pList)-1; j++ {
			if pList[i-1].DistanceManhatten(pList[j]) == 1 {
				pList[i], pList[j] = pList[j], pList[i]
				break
			}
		}
	}
	return pList, nil
}

func countCheats(racewayPoints []point, minSavings int) int {
	cheatCount := 0
	for leftIdx := range len(racewayPoints) - minSavings {
		for rightIdx := len(racewayPoints) - 1; leftIdx+minSavings < rightIdx; rightIdx-- {
			if racewayPoints[leftIdx].DistanceManhatten(racewayPoints[rightIdx]) == 2 {
				cheatCount++
			}
		}
	}
	return cheatCount
}

func Part1(in io.Reader) (int, error) {
	patterns, err := parseInput(in)
	if err != nil {
		return 0, err
	}
	return countCheats(patterns, 100), nil
}
