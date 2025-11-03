package day15_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day15"
)

const testDataBig string = `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`

const testDataSmall string = `########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`

func TestTranslateMovements(t *testing.T) {
	d := []byte("<vv>^<")
	movements := *day15.TranslateMovements(&d)
	if len(movements) != 6 {
		t.Errorf(`Expected %v to have length of %d`, movements, 6)
	}
}

func TestParseInput(t *testing.T) {
	field, player, moves := day15.ParseInput(strings.NewReader(testDataBig), false)
	if len(field) != 10 {
		t.Errorf(`Expected %d to match %d`, len(field), 10)
	}
	if player.Y != 4 || player.X != 4 {
		t.Errorf(`Expected %v to match [4,4]`, player)
	}
	if len(moves) != 700 {
		t.Errorf(`Expected %d to match %d`, len(moves), 700)
	}
}
func TestParseInputDoubleWide(t *testing.T) {
	field, player, moves := day15.ParseInput(strings.NewReader(testDataBig), true)
	if len(field) != 10 || len(field[0]) != 20 {
		t.Errorf(`Expected field to have size(10,20) not (%d,%d)`, len(field), len(field[0]))
	}
	if player.Y != 4 || player.X != 8 {
		t.Errorf(`Expected %v to match [4,8]`, player)
	}
	if len(moves) != 700 {
		t.Errorf(`Expected %d to match %d`, len(moves), 700)
	}
}
func TestFindNextEmptyCellOffset_SingleBlock(t *testing.T) {
	field, _, _ := day15.ParseInput(strings.NewReader(testDataBig), false)
	testPlayer := day15.Player{Y: 4, X: 4}
	testMove := day15.Move{Y: 0, X: -1}
	off, ok := day15.FindNextEmptyCellOffset(&field, &testPlayer, &testMove)
	if !ok || off != 1 {
		t.Errorf(`Expected offset to be ok and 1 not %d`, off)
	}
}
func TestFindNextEmptyCellOffset_DoubleBlock(t *testing.T) {
	field, _, _ := day15.ParseInput(strings.NewReader(testDataBig), false)
	testPlayer := day15.Player{Y: 3, X: 4}
	testMove := day15.Move{Y: 0, X: -1}
	off, ok := day15.FindNextEmptyCellOffset(&field, &testPlayer, &testMove)
	if !ok || off != 2 {
		t.Errorf(`Expected offset to be ok and 2 not %d`, off)
	}
}
func TestFindNextEmptyCellOffset_TripleBlock(t *testing.T) {
	field, _, _ := day15.ParseInput(strings.NewReader(testDataBig), false)
	testPlayer := day15.Player{Y: 3, X: 4}
	field[4][4] = day15.CELL_BOX
	field[5][4] = day15.CELL_BOX
	testMove := day15.Move{Y: 1, X: 0}
	off, ok := day15.FindNextEmptyCellOffset(&field, &testPlayer, &testMove)
	if !ok || off != 3 {
		t.Errorf(`Expected offset to be ok and 3 not %d`, off)
	}
}
func TestFindNextEmptyCellOffset_QuadBlock(t *testing.T) {
	// Testing 4 Blocks Space Wall
	field, _, _ := day15.ParseInput(strings.NewReader(testDataBig), false)
	testPlayer := day15.Player{Y: 3, X: 4}
	field[4][4] = day15.CELL_BOX
	field[5][4] = day15.CELL_BOX
	field[6][4] = day15.CELL_BOX
	field[7][4] = day15.CELL_BOX
	field[8][4] = day15.CELL_EMPTY
	testMove := day15.Move{Y: 1, X: 0}
	off, ok := day15.FindNextEmptyCellOffset(&field, &testPlayer, &testMove)
	if !ok || off != 4 {
		t.Errorf(`Expected offset to be ok and 4 not %d`, off)
	}
}
func TestFindNextEmptyCellOffset_BlocksAgainstWall(t *testing.T) {
	// Testing 5 Blocks Wall
	field, _, _ := day15.ParseInput(strings.NewReader(testDataBig), false)
	testPlayer := day15.Player{Y: 3, X: 4}
	field[4][4] = day15.CELL_BOX
	field[5][4] = day15.CELL_BOX
	field[6][4] = day15.CELL_BOX
	field[7][4] = day15.CELL_BOX
	field[8][4] = day15.CELL_BOX
	testMove := day15.Move{Y: 1, X: 0}
	off, ok := day15.FindNextEmptyCellOffset(&field, &testPlayer, &testMove)
	if ok {
		t.Errorf(`Expected offset to be ok and 4 not %d`, off)
	}
}
func TestPart1(t *testing.T) {
	t.Run("large", func(t *testing.T) {
		result := day15.Part1(strings.NewReader(testDataBig))
		if result != 10092 {
			t.Errorf(`Expected %d to match 10092`, result)
		}
	})
	t.Run("small", func(t *testing.T) {
		result := day15.Part1(strings.NewReader(testDataSmall))
		if result != 2028 {
			t.Errorf(`Expected %d to match 2028`, result)
		}
	})
}

func BenchmarkPart1(b *testing.B) {
	data := strings.NewReader(testDataBig)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		day15.Part1(data)
	}
}
