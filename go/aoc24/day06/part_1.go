package day06

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	DIR_UP    int16 = 0
	DIR_RIGHT       = 1
	DIR_DOWN        = 2
	DIR_LEFT        = 3
)

type Guard struct {
	x   int16
	y   int16
	dir int16
}

func (a *Guard) cmp(g *Guard) bool {
	return (*a).x == (*g).x && (*a).y == (*g).y && (*a).dir == (*g).dir
}

func NewGuard(x int16, y int16, dir int16) *Guard {
	return &Guard{x: x, y: y, dir: dir}
}

type Field struct {
	width, height int16
	field         [][]bool
}

func NewField(width int16, height int16, field [][]bool) Field {
	return Field{width: width, height: height, field: field}
}

func readSource(reader io.Reader) (Field, Guard, error) {
	obstacles := make([][]bool, 0)
	var guard *Guard
	s := bufio.NewScanner(reader)
	y := int16(0)
	width := int16(0)
	for s.Scan() {
		line := s.Bytes()
		if width == 0 {
			width = int16(len(line))
		} else if width != int16(len(line)) {
			panic("Line length is uneven")
		}
		row := make([]bool, len(line))
		for idx, character := range line {
			switch character {
			case '^':
				guard = NewGuard(int16(idx), y, DIR_DOWN)
			case '>':
				guard = NewGuard(int16(idx), y, DIR_RIGHT)
			case '<':
				guard = NewGuard(int16(idx), y, DIR_LEFT)
			case 'v':
				guard = NewGuard(int16(idx), y, DIR_UP)
			case '#':
				row[idx] = true
			}
		}
		obstacles = append(obstacles, row)
		y++
	}
	return NewField(width, y, obstacles), *guard, nil
}
func OutOfBounds(field *Field, x *int16, y *int16) bool {
	return *x < 0 || *y < 0 || *x >= (*field).width || *y >= (*field).height
}

func PrintField(field *Field) {
	for y := range field.height {
		if y > 0 {
			fmt.Print("\n")
		}
		for x := range field.width {
			if field.field[y][x] {
				fmt.Printf("1")
			} else {
				fmt.Printf("0")
			}
		}
	}
	fmt.Print("\n")
}

func leaveArea(field *Field, guard Guard) map[[2]int16]bool {
	stepsTaken := make(map[[2]int16]bool)
	stepsTaken[[2]int16{guard.y, guard.x}] = true
	for {
		if guard.dir == DIR_DOWN {
			nextY := guard.y - 1
			if OutOfBounds(field, &guard.x, &nextY) {
				return stepsTaken
			} else if (*field).field[nextY][guard.x] {
				guard.dir = DIR_RIGHT
				continue
			} else {
				guard.y = nextY
				stepsTaken[[2]int16{nextY, guard.x}] = true
			}
		} else if guard.dir == DIR_UP {
			nextY := guard.y + 1
			if OutOfBounds(field, &guard.x, &nextY) {
				return stepsTaken
			} else if (*field).field[nextY][guard.x] {
				guard.dir = DIR_LEFT
				continue
			} else {
				guard.y = nextY
				stepsTaken[[2]int16{nextY, guard.x}] = true
			}
		} else if guard.dir == DIR_RIGHT {
			nextX := guard.x + 1
			if OutOfBounds(field, &nextX, &guard.y) {
				return stepsTaken
			} else if (*field).field[guard.y][nextX] {
				guard.dir = DIR_UP
				continue
			} else {
				guard.x = nextX
				stepsTaken[[2]int16{guard.y, nextX}] = true
			}
		} else if guard.dir == DIR_LEFT {
			nextX := guard.x - 1
			if OutOfBounds(field, &nextX, &guard.y) {
				return stepsTaken
			} else if (*field).field[guard.y][nextX] {
				guard.dir = DIR_DOWN
				continue
			} else {
				guard.x = nextX
				stepsTaken[[2]int16{guard.y, nextX}] = true
			}
		}
	}
}

func Part1() {
	obstacles, guard, err := readSource(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", len(leaveArea(&obstacles, guard)))

}
