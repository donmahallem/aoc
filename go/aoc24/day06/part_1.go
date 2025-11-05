package day06

import (
	"bufio"
	"io"
)

const (
	DIR_UP    uint16 = 0
	DIR_RIGHT        = 1
	DIR_DOWN         = 2
	DIR_LEFT         = 3
)

type Guard struct {
	x   uint16
	y   uint16
	dir uint16
}

func (a *Guard) cmp(g *Guard) bool {
	return (*a).x == (*g).x && (*a).y == (*g).y && (*a).dir == (*g).dir
}
func translateToIndex(x, y uint16, width uint16) uint16 {
	return y*width + x
}

type Field struct {
	width, height uint16
	obstacles     map[uint16]struct{}
}

func NewField(width uint16, height uint16, obstacles map[uint16]struct{}) Field {
	return Field{width: width, height: height, obstacles: obstacles}
}

func (f *Field) HasObstacle(x, y uint16) bool {
	_, ok := f.obstacles[(y*f.width)+x]
	return ok
}

func (f *Field) Put(x, y uint16) {
	f.obstacles[(y*f.width)+x] = struct{}{}
}

func NewGuard(x, y uint16, dir uint16) *Guard {
	return &Guard{x: x, y: y, dir: dir}
}

func ReadSource(reader io.Reader) (Field, Guard, error) {
	obstacles := make(map[uint16]struct{})
	var guard *Guard
	s := bufio.NewScanner(reader)
	y := uint16(0)
	width := uint16(0)
	for s.Scan() {
		line := s.Bytes()
		if width == 0 {
			width = uint16(len(line))
		} else if width != uint16(len(line)) {
			panic("Line length is uneven")
		}
		for idx := uint16(0); idx < width; idx++ {
			character := line[idx]
			switch character {
			case '^':
				guard = NewGuard(idx, y, DIR_DOWN)
			case '>':
				guard = NewGuard(idx, y, DIR_RIGHT)
			case '<':
				guard = NewGuard(idx, y, DIR_LEFT)
			case 'v':
				guard = NewGuard(idx, y, DIR_UP)
			case '#':
				obstacles[translateToIndex(idx, y, width)] = struct{}{}
			}
		}
		y++
	}
	if guard == nil {
		panic("guard not found in input")
	}
	return NewField(width, y, obstacles), *guard, nil
}

type PathStep struct {
	Index    uint16
	PreGuard Guard
}

type PathInfo struct {
	UniqueCount int
	Steps       []PathStep
}

/*
resolves the path taken by the guard until it leaves the area
*/
func leaveArea(field *Field, guard Guard) PathInfo {
	total := int(field.width) * int(field.height)
	visited := make([]bool, total)
	steps := make([]PathStep, 0, total)

	startIdx := translateToIndex(guard.x, guard.y, field.width)
	visited[int(startIdx)] = true
	unique := 1

	for {
		switch guard.dir {
		case DIR_DOWN:
			if guard.y == 0 {
				return PathInfo{UniqueCount: unique, Steps: steps}
			}
			nextY := guard.y - 1
			if field.HasObstacle(guard.x, nextY) {
				guard.dir = DIR_RIGHT
				continue
			}
			targetIdx := translateToIndex(guard.x, nextY, field.width)
			if !visited[int(targetIdx)] {
				visited[int(targetIdx)] = true
				steps = append(steps, PathStep{Index: targetIdx, PreGuard: guard})
				unique++
			}
			guard.y = nextY

		case DIR_UP:
			if guard.y+1 >= field.height {
				return PathInfo{UniqueCount: unique, Steps: steps}
			}
			nextY := guard.y + 1
			if field.HasObstacle(guard.x, nextY) {
				guard.dir = DIR_LEFT
				continue
			}
			targetIdx := translateToIndex(guard.x, nextY, field.width)
			if !visited[int(targetIdx)] {
				visited[int(targetIdx)] = true
				steps = append(steps, PathStep{Index: targetIdx, PreGuard: guard})
				unique++
			}
			guard.y = nextY

		case DIR_RIGHT:
			if guard.x+1 >= field.width {
				return PathInfo{UniqueCount: unique, Steps: steps}
			}
			nextX := guard.x + 1
			if field.HasObstacle(nextX, guard.y) {
				guard.dir = DIR_UP
				continue
			}
			targetIdx := translateToIndex(nextX, guard.y, field.width)
			if !visited[int(targetIdx)] {
				visited[int(targetIdx)] = true
				steps = append(steps, PathStep{Index: targetIdx, PreGuard: guard})
				unique++
			}
			guard.x = nextX

		case DIR_LEFT:
			if guard.x == 0 {
				return PathInfo{UniqueCount: unique, Steps: steps}
			}
			nextX := guard.x - 1
			if field.HasObstacle(nextX, guard.y) {
				guard.dir = DIR_DOWN
				continue
			}
			targetIdx := translateToIndex(nextX, guard.y, field.width)
			if !visited[int(targetIdx)] {
				visited[int(targetIdx)] = true
				steps = append(steps, PathStep{Index: targetIdx, PreGuard: guard})
				unique++
			}
			guard.x = nextX
		}
	}
}

func Part1(in io.Reader) int {
	field, guard, err := ReadSource(in)
	if err != nil {
		panic(err)
	}
	info := leaveArea(&field, guard)
	return info.UniqueCount
}
