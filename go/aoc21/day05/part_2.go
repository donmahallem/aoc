package day05

import (
	"io"
)

// assumes all paths be in +x direction and one atleast diagonal (other one can be horizontal,vertical or diagonal)
// assumes all paths are normalized in +x direction (x1 <= x2) and at least one is diagonal
func overlapsPart2(a, b coords, overlaps map[point]struct{}) {
	if a[4]&DirectionDiagonal != 0 && b[4]&DirectionDiagonal != 0 && a[4] == b[4] {
		// Collinear diagonals with the same direction
		m := 1
		if a[4] == DirectionDiagonalDown {
			m = -1
		}
		cA := a[1] - m*a[0]
		cB := b[1] - m*b[0]
		if cA == cB {
			startX := max(a[0], b[0])
			endX := min(a[2], b[2])
			for x := startX; x <= endX; x++ {
				overlaps[point{x: x, y: m*x + cA}] = struct{}{}
			}
		}
	} else {
		// Make a always the diagonal segment if only one is diagonal
		if a[4]&DirectionDiagonal == 0 {
			a, b = b, a
		}

		mA := 1
		if a[4] == DirectionDiagonalDown {
			mA = -1
		}
		cA := a[1] - mA*a[0]

		if b[4] == DirectionVertical {
			// Diagonal 'a' intersecting vertical 'b'
			vx := b[0]
			iy := mA*vx + cA
			if vx >= max(a[0], b[0]) && vx <= min(a[2], b[2]) &&
				iy >= max(min(a[1], a[3]), min(b[1], b[3])) &&
				iy <= min(max(a[1], a[3]), max(b[1], b[3])) {
				overlaps[point{x: vx, y: iy}] = struct{}{}
			}
		} else if b[4] == DirectionHorizontal {
			// Diagonal 'a' intersecting horizontal 'b'
			hy := b[1]
			ix := (hy - cA) * mA
			if ix >= max(a[0], b[0]) && ix <= min(a[2], b[2]) &&
				hy >= max(min(a[1], a[3]), min(b[1], b[3])) &&
				hy <= min(max(a[1], a[3]), max(b[1], b[3])) {
				overlaps[point{x: ix, y: hy}] = struct{}{}
			}
		} else if b[4]&DirectionDiagonal != 0 {
			// Opposite-slope diagonals: y = mA*x + cA and y = mB*x + cB where mB = -mA
			mB := -mA
			cB := b[1] - mB*b[0]
			diffC := cB - cA
			// Check if intersection is not fractional
			if diffC%2 == 0 {
				ix := diffC / (mA - mB)
				iy := mA*ix + cA
				if ix >= max(a[0], b[0]) && ix <= min(a[2], b[2]) &&
					iy >= max(min(a[1], a[3]), min(b[1], b[3])) &&
					iy <= min(max(a[1], a[3]), max(b[1], b[3])) {
					overlaps[point{x: ix, y: iy}] = struct{}{}
				}
			}
		}
	}
}
func Part2(in io.Reader) (int, error) {
	arrows, err := parseInput(in)
	if err != nil {
		return 0, err
	}

	overlaps := make(map[point]struct{}, 2048)

	for i := 0; i < len(arrows)-1; i++ {
		a := arrows[i]
		if a[4] == DirectionOther {
			continue
		}

		for j := i + 1; j < len(arrows); j++ {
			b := arrows[j]
			if b[4] == DirectionOther {
				continue
			} else if a[4]&DirectionDiagonal != 0 || b[4]&DirectionDiagonal != 0 {
				overlapsPart2(a, b, overlaps)
			} else {
				overlapsPart1(a, b, overlaps)
			}
		}
	}

	return len(overlaps), nil
}
