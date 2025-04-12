package day18

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/donmahallem/aoc/aoc_utils"
)

type Register = [3]int
type Program = []int
type Point = aoc_utils.Point[int]

func ParseInput(in io.Reader) *[]Point {
	points := make([]Point, 0)
	s := bufio.NewScanner(in)
	for s.Scan() {
		line := s.Text()
		commaIndex := strings.Index(line, ",")
		point := Point{}
		point.X, _ = strconv.Atoi(line[0:commaIndex])
		point.Y, _ = strconv.Atoi(line[commaIndex+1:])
		points = append(points, point)
	}
	return &points
}
