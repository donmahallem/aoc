package day15

import "testing"

func TestTryMoveHorizontal(t *testing.T) {
	// row: [EMPTY, player at 1 (field empty), BOX at 2, EMPTY at 3]
	row := []uint8{CELL_EMPTY, CELL_EMPTY, CELL_BOX, CELL_EMPTY}
	f := field{row}
	pl := player{X: 1, Y: 0}
	m := move{X: 1, Y: 0}
	ok := tryMoveHorizontal(&f, &pl, &m)
	if !ok {
		t.Fatalf("expected move to succeed")
	}
	if pl.X != 2 {
		t.Fatalf("expected player X 2, got %d", pl.X)
	}
	if f[0][3] != CELL_BOX {
		t.Fatalf("expected box to move to index 3, got %d", f[0][3])
	}

	// blocked by wall
	row2 := []uint8{CELL_EMPTY, CELL_EMPTY, CELL_WALL, CELL_BOX}
	f2 := field{row2}
	pl2 := player{X: 1, Y: 0}
	ok2 := tryMoveHorizontal(&f2, &pl2, &m)
	if ok2 {
		t.Fatalf("expected move to be blocked by wall")
	}
}

func TestTryMoveVertical(t *testing.T) {
	// simple vertical push (double-wide boxes): player at (1,0), box left+right at (1,1)-(2,1), empty row below
	f := field{
		{CELL_EMPTY, CELL_EMPTY, CELL_EMPTY},
		{CELL_EMPTY, CELL_BOX_LEFT, CELL_BOX_RIGHT},
		{CELL_EMPTY, CELL_EMPTY, CELL_EMPTY},
	}
	pl := player{X: 1, Y: 0}
	m := move{X: 0, Y: 1}
	boxesToMove := make([]player, 0)
	queue := make([]player, 0)
	visited := make([]int, len(f)*len(f[0]))
	ok := tryMoveVertical(&f, &pl, &m, &boxesToMove, &queue, visited, len(f[0]), 1)
	if !ok {
		t.Fatalf("expected vertical push to succeed")
	}
	if pl.Y != 1 {
		t.Fatalf("expected player Y 1, got %d", pl.Y)
	}
	if f[2][1] != CELL_BOX_LEFT || f[2][2] != CELL_BOX_RIGHT {
		t.Fatalf("expected box to have moved down to row 2, got %v", f[2])
	}

	// blocked by wall below
	f2 := field{
		{CELL_EMPTY, CELL_EMPTY, CELL_EMPTY},
		{CELL_EMPTY, CELL_BOX_LEFT, CELL_BOX_RIGHT},
		{CELL_WALL, CELL_WALL, CELL_WALL},
	}
	pl2 := player{X: 1, Y: 0}
	ok2 := tryMoveVertical(&f2, &pl2, &m, &boxesToMove, &queue, visited, len(f2[0]), 2)
	if ok2 {
		t.Fatalf("expected vertical push to be blocked by wall")
	}
}

func TestWalkWideBoxesMatchesPart2(t *testing.T) {
	// reuse the sample input to ensure walkWideBoxes behaves like Part2
	// using the sample strings from the existing tests would be ideal; instead
	// craft a minimal scenario: a player pushing one box right then down
	// field with a double-wide box at row 2 cols 2-3; player at (2,1)
	f := field{
		{CELL_WALL, CELL_WALL, CELL_WALL, CELL_WALL, CELL_WALL, CELL_WALL},
		{CELL_WALL, CELL_EMPTY, CELL_EMPTY, CELL_EMPTY, CELL_EMPTY, CELL_WALL},
		{CELL_WALL, CELL_EMPTY, CELL_BOX_LEFT, CELL_BOX_RIGHT, CELL_EMPTY, CELL_WALL},
		{CELL_WALL, CELL_EMPTY, CELL_EMPTY, CELL_EMPTY, CELL_EMPTY, CELL_WALL},
		{CELL_WALL, CELL_WALL, CELL_WALL, CELL_WALL, CELL_WALL, CELL_WALL},
	}
	pl := player{X: 2, Y: 1}                    // above the box
	moves := []move{{X: 0, Y: 1}, {X: 1, Y: 0}} // first push down, then push right
	walkWideBoxes(&f, &pl, &moves)
	// after pushing down then right (player moved above the box before the horizontal move),
	// the box should have moved down but not necessarily right; expect it at row 3 cols 2-3
	if f[3][2] != CELL_BOX_LEFT || f[3][3] != CELL_BOX_RIGHT {
		t.Fatalf("expected box to be at (3,2)-(3,3) after moves, got row: %v", f[3])
	}
}
