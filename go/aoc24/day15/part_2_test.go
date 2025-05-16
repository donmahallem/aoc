package day15_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day15"
)

const testDataSimple string = `#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######

<vv<<^^<<^^`

func createEmptyDoubleWideField(width, height int16) day15.Field {
	field := make(day15.Field, height)
	for y := range height {
		field[y] = make([]uint8, width)
		for x := range width {
			if (y == 0 || y == height-1) || x < 2 || x >= width-2 {
				field[y][x] = day15.CELL_WALL
			}
		}
	}
	return field
}

func setDoubleBox(field *day15.Field, x, y int16) {
	(*field)[y][x] = day15.CELL_BOX_LEFT
	(*field)[y][x+1] = day15.CELL_BOX_RIGHT
}

func IsWideBox(field *day15.Field, x int16, y int16) bool {
	return (*field)[y][x] == day15.CELL_BOX_LEFT && (*field)[y][x+1] == day15.CELL_BOX_RIGHT
}

func TestTryMoveHorizontal_Left_SingleBlock(t *testing.T) {
	field, _, _ := day15.ParseInput(strings.NewReader(testDataBig), true)
	testPlayer := day15.Player{Y: 4, X: 8}
	testMove := day15.Move{Y: 0, X: -1}
	testCandidates := make([]day15.Player, 0)
	ok := day15.TryMoveHorizontal(&field, &testPlayer, &testMove)
	if !ok {
		t.Errorf(`Expected offset to be ok and 1 not %d`, len(testCandidates))
	}
	if testPlayer.X != 7 || testPlayer.Y != 4 {
		t.Errorf(`Expected player to be at position (7,4) not %v`, testPlayer)
	}
	if len(testCandidates) == 1 {
		t.Errorf(`Expected offset to be ok and 1 not %d`, len(testCandidates))
	}
}
func TestTryMoveHorizontal_Right_SingleBlock(t *testing.T) {
	field := createEmptyDoubleWideField(10, 10)
	setDoubleBox(&field, 3, 4)
	testPlayer := day15.Player{Y: 4, X: 2}
	testMove := day15.Move{Y: 0, X: 1}
	ok := day15.TryMoveHorizontal(&field, &testPlayer, &testMove)
	if !ok {
		t.Errorf(`Expected to be ok`)
	}
	if !IsWideBox(&field, 4, 4) {
		t.Errorf(`Expected big box to be at position (4,4) not %v`, testPlayer)
	}
}
func TestTryMoveHorizontal_Right_DoubleBlock(t *testing.T) {
	field := createEmptyDoubleWideField(10, 10)
	setDoubleBox(&field, 3, 4)
	setDoubleBox(&field, 5, 4)
	testPlayer := day15.Player{Y: 4, X: 2}
	testMove := day15.Move{Y: 0, X: 1}
	ok := day15.TryMoveHorizontal(&field, &testPlayer, &testMove)
	if !ok {
		t.Errorf(`Expected to be ok`)
	}
	if !IsWideBox(&field, 4, 4) {
		t.Errorf(`Expected big box to be at position (4,4) not %v`, testPlayer)
	}
	if !IsWideBox(&field, 6, 4) {
		t.Errorf(`Expected big box to be at position (6,4) not %v`, testPlayer)
	}
	ok = day15.TryMoveHorizontal(&field, &testPlayer, &testMove)
	if ok {
		t.Errorf(`Expected to be not valid move`)
	}
}
func TestTryMoveHorizontal_DoubleBlock(t *testing.T) {
	field, _, _ := day15.ParseInput(strings.NewReader(testDataBig), true)
	testPlayer := day15.Player{Y: 3, X: 8}
	testMove := day15.Move{Y: 0, X: -1}
	testCandidates := make([]day15.Player, 0)
	ok := day15.TryMoveHorizontal(&field, &testPlayer, &testMove)
	if !ok {
		t.Errorf(`Expected offset to be ok and 1 not %d`, len(testCandidates))
	}
	if testPlayer.X != 7 || testPlayer.Y != 3 {
		t.Errorf(`Expected player to be at position (7,3) not %v`, testPlayer)
	}
	if len(testCandidates) == 1 {
		t.Errorf(`Expected offset to be ok and 1 not %d`, len(testCandidates))
	}
	if !IsWideBox(&field, 5, 3) {
		t.Errorf(`Expected a wide box at (6,3)`)
	}
	if !IsWideBox(&field, 3, 3) {
		t.Errorf(`Expected a wide box at (6,3)`)
	}
}
func TestTryMoveHorizontal_DoubleBlockAgainstWall(t *testing.T) {
	field, _, _ := day15.ParseInput(strings.NewReader(testDataBig), true)
	testPlayer := day15.Player{Y: 3, X: 8}
	testMove := day15.Move{Y: 0, X: -1}
	testCandidates := make([]day15.Player, 0)
	var ok bool = true
	for range 20 {
		ok = ok && day15.TryMoveHorizontal(&field, &testPlayer, &testMove)
	}
	if ok {
		t.Errorf(`Expected return to be false`)
	}
	if testPlayer.X != 6 || testPlayer.Y != 3 {
		t.Errorf(`Expected player to be at position (7,3) not %v`, testPlayer)
	}
	if len(testCandidates) == 1 {
		t.Errorf(`Expected offset to be ok and 1 not %d`, len(testCandidates))
	}
	if !IsWideBox(&field, 2, 3) {
		t.Errorf(`Expected a wide box at (2,3)`)
	}
	if !IsWideBox(&field, 4, 3) {
		t.Errorf(`Expected a wide box at (4,3)`)
	}
}
func TestFindMoveCandidates_SingleBlock(t *testing.T) {
	field := createEmptyDoubleWideField(10, 10)
	setDoubleBox(&field, 5, 5)
	testPlayer := day15.Player{Y: 5, X: 5}
	testMove := day15.Move{Y: -1, X: 0}
	day15.MoveBoxesVertically(&field, &testPlayer, &testMove)
	if !IsWideBox(&field, 5, 4) {
		t.Errorf(`Expected return value to be true`)
	}
}
func TestFindMoveCandidates_Triangle(t *testing.T) {
	testBoxes := make([]day15.Candidate, 0, 4)
	testBoxes = append(testBoxes,
		day15.Candidate{X: 5, Y: 5},
		day15.Candidate{X: 4, Y: 4},
		day15.Candidate{X: 6, Y: 4})
	field := createEmptyDoubleWideField(10, 10)
	for i := range len(testBoxes) {
		setDoubleBox(&field, testBoxes[i].X, testBoxes[i].Y)
	}
	testPlayer := day15.Player{Y: 6, X: 5}
	testMove := day15.Move{Y: -1, X: 0}
	ok := day15.IsMovableVertical(&field, &testPlayer, &testMove)
	if !ok {
		t.Errorf(`Expected return value to be true`)
	}
}
func TestFindMoveCandidates_Checkmark(t *testing.T) {
	testBoxes := make([]day15.Candidate, 0, 4)
	testBoxes = append(testBoxes,
		day15.Candidate{X: 5, Y: 5},
		day15.Candidate{X: 4, Y: 4},
		day15.Candidate{X: 6, Y: 4},
		day15.Candidate{X: 3, Y: 3})
	field := createEmptyDoubleWideField(10, 10)
	for i := range len(testBoxes) {
		setDoubleBox(&field, testBoxes[i].X, testBoxes[i].Y)
	}
	testPlayer := day15.Player{Y: 5, X: 5}
	testMove := day15.Move{Y: -1, X: 0}
	ok := day15.IsMovableVertical(&field, &testPlayer, &testMove)
	if !ok {
		t.Errorf(`Expected return value to be true`)
	}
}
func TestFindMoveCandidates_Checkmark_Blocked(t *testing.T) {
	testBoxes := make([]day15.Candidate, 0, 4)
	testBoxes = append(testBoxes,
		day15.Candidate{X: 5, Y: 5},
		day15.Candidate{X: 4, Y: 4},
		day15.Candidate{X: 6, Y: 4},
		day15.Candidate{X: 3, Y: 3})
	field := createEmptyDoubleWideField(10, 10)
	for i := range len(testBoxes) {
		setDoubleBox(&field, testBoxes[i].X, testBoxes[i].Y)
	}
	field[2][3] = day15.CELL_WALL
	testPlayer := day15.Player{Y: 5, X: 5}
	testMove := day15.Move{Y: -1, X: 0}
	ok := day15.IsMovableVertical(&field, &testPlayer, &testMove)
	if ok {
		t.Errorf(`Expected return value to be false`)
	}
}

func TestMoveBoxesVertically_Left(t *testing.T) {
	testBoxes := make([]day15.Candidate, 0, 4)
	testBoxes = append(testBoxes,
		day15.Candidate{X: 5, Y: 5},
		day15.Candidate{X: 4, Y: 4},
		day15.Candidate{X: 6, Y: 4},
		day15.Candidate{X: 3, Y: 3})
	field := createEmptyDoubleWideField(10, 10)
	for i := range len(testBoxes) {
		setDoubleBox(&field, testBoxes[i].X, testBoxes[i].Y)
	}
	testPlayer := day15.Player{Y: 5, X: 5}
	testMove := day15.Move{Y: -1, X: 0}
	day15.MoveBoxesVertically(&field, &testPlayer, &testMove)
	if len(testBoxes) != 4 {
		t.Errorf(`Expected to receive 2 move targets not %d`, len(testBoxes))
	}
	for i := range len(testBoxes) {
		if !IsWideBox(&field, testBoxes[i].X, testBoxes[i].Y-1) {
			t.Errorf(`Expected box to be at at (%d,%d)`, testBoxes[i].X, testBoxes[i].Y)
		}
	}
}
func TestTryMoveVertical_Left_Down(t *testing.T) {
	testBoxes := make([]day15.Candidate, 0, 4)
	testBoxes = append(testBoxes,
		day15.Candidate{X: 10, Y: 5},
		day15.Candidate{X: 8, Y: 6},
		day15.Candidate{X: 10, Y: 7},
		day15.Candidate{X: 10, Y: 8})
	field := createEmptyDoubleWideField(20, 10)
	for i := range len(testBoxes) {
		setDoubleBox(&field, testBoxes[i].X, testBoxes[i].Y)
	}
	testPlayer := day15.Player{Y: 4, X: 10}
	testMove := day15.Move{Y: 1, X: 0}
	ok := day15.TryMoveVertical(&field, &testPlayer, &testMove)
	if !ok {
		t.Error(`Expected to result to be true`)
	}
	if len(testBoxes) != 4 {
		t.Errorf(`Expected to recieve 2 move targets not %d`, len(testBoxes))
	}
	for i := 1; i < len(testBoxes[1:]); i++ {
		if !IsWideBox(&field, testBoxes[i].X, testBoxes[i].Y) {
			t.Errorf(`Expected box to be at at (%d,%d)`, testBoxes[i].X, testBoxes[i].Y)
		}
	}
	ok = day15.TryMoveVertical(&field, &testPlayer, &testMove)
	if ok {
		t.Error(`Expected to result to be false`)
	}
}
func TestTryMoveVertical_Left_Down_DoubleSpace(t *testing.T) {
	testBoxes := make([]day15.Candidate, 0, 4)
	testBoxes = append(testBoxes,
		day15.Candidate{X: 10, Y: 5},
		day15.Candidate{X: 8, Y: 6},
		day15.Candidate{X: 10, Y: 7},
		day15.Candidate{X: 10, Y: 8})
	field := createEmptyDoubleWideField(20, 10)
	for i := range len(testBoxes) {
		setDoubleBox(&field, testBoxes[i].X, testBoxes[i].Y)
	}
	testPlayer := day15.Player{Y: 3, X: 10}
	testMove := day15.Move{Y: 1, X: 0}
	ok := day15.TryMoveVertical(&field, &testPlayer, &testMove)
	if !ok {
		t.Error(`Expected to result to be true`)
	}
	if len(testBoxes) != 4 {
		t.Errorf(`Expected to recieve 2 move targets not %d`, len(testBoxes))
	}
	for i := range len(testBoxes[1:]) {
		if !IsWideBox(&field, testBoxes[i].X, testBoxes[i].Y) {
			t.Errorf(`Expected box to be at at (%d,%d)`, testBoxes[i].X, testBoxes[i].Y)
		}
	}
	ok = day15.TryMoveVertical(&field, &testPlayer, &testMove)
	if !ok {
		t.Error(`Expected to result to be false`)
	}
}
func TestTryMoveVertical_Right_Down(t *testing.T) {
	testBoxes := make([]day15.Candidate, 0, 4)
	testBoxes = append(testBoxes,
		day15.Candidate{X: 10, Y: 5},
		day15.Candidate{X: 8, Y: 6},
		day15.Candidate{X: 10, Y: 7},
		day15.Candidate{X: 10, Y: 8})
	field := createEmptyDoubleWideField(20, 10)
	for i := range len(testBoxes) {
		setDoubleBox(&field, testBoxes[i].X, testBoxes[i].Y)
	}
	testPlayer := day15.Player{Y: 4, X: 11}
	testMove := day15.Move{Y: 1, X: 0}
	ok := day15.TryMoveVertical(&field, &testPlayer, &testMove)
	if !ok {
		t.Error(`Expected to result to be true`)
	}
	if len(testBoxes) != 4 {
		t.Errorf(`Expected to recieve 2 move targets not %d`, len(testBoxes))
	}
	for i := 1; i < len(testBoxes[1:]); i++ {
		if !IsWideBox(&field, testBoxes[i].X, testBoxes[i].Y) {
			t.Errorf(`Expected box to be at at (%d,%d)`, testBoxes[i].X, testBoxes[i].Y)
		}
	}
	ok = day15.TryMoveVertical(&field, &testPlayer, &testMove)
	if ok {
		t.Error(`Expected to result to be false`)
	}
}
func TestTryMoveVertical_Left_Down_AcrossEdge(t *testing.T) {
	testBoxes := make([]day15.Candidate, 0, 4)
	testBoxes = append(testBoxes,
		day15.Candidate{X: 9, Y: 5},
		day15.Candidate{X: 8, Y: 6},
		day15.Candidate{X: 10, Y: 7},
		day15.Candidate{X: 10, Y: 8})
	field := createEmptyDoubleWideField(20, 10)
	for i := range len(testBoxes) {
		setDoubleBox(&field, testBoxes[i].X, testBoxes[i].Y)
	}
	testPlayer := day15.Player{Y: 4, X: 10}
	testMove := day15.Move{Y: 1, X: 0}
	ok := day15.TryMoveVertical(&field, &testPlayer, &testMove)
	if !ok {
		t.Error(`Expected to result to be true`)
	}
	if len(testBoxes) != 4 {
		t.Errorf(`Expected to recieve 2 move targets not %d`, len(testBoxes))
	}
	for i := 2; i < len(testBoxes[1:]); i++ {
		if !IsWideBox(&field, testBoxes[i].X, testBoxes[i].Y) {
			t.Errorf(`Expected box to be at at (%d,%d)`, testBoxes[i].X, testBoxes[i].Y)
		}
	}
	ok = day15.TryMoveVertical(&field, &testPlayer, &testMove)
	if ok {
		t.Error(`Expected to result to be false`)
	}
}
func TestPart2_large(t *testing.T) {
	result := day15.Part2(strings.NewReader(testDataBig))
	if result != 9021 {
		t.Errorf(`Expected %d to match 9021`, result)
	}
}

func BenchmarkPart2(b *testing.B) {
	data := strings.NewReader(testDataBig)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		day15.Part2(data)
	}
}
