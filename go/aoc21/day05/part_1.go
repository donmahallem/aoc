package day05

import (
	"io"
)

func overlapsPart1(a, b coords, overlaps map[point]struct{}) {
	if a[4] == b[4] {
		switch a[4] {
		case DirectionHorizontal:
			if a[1] != b[1] {
				return
			}
			startX := max(a[0], b[0])
			endX := min(a[2], b[2])
			y := a[1]
			for x := startX; x <= endX; x++ {
				overlaps[point{x: x, y: y}] = struct{}{}
			}
			return

		case DirectionVertical:
			if a[0] != b[0] {
				return
			}
			startY := max(min(a[1], a[3]), min(b[1], b[3]))
			endY := min(max(a[1], a[3]), max(b[1], b[3]))
			x := a[0]
			for y := startY; y <= endY; y++ {
				overlaps[point{x: x, y: y}] = struct{}{}
			}
			return

		}
		return
	}

	if a[4] == DirectionVertical {
		a, b = b, a
	}

	// a is horizontal (a[0] <= a[2], a[1] == a[3]), b is vertical (b[0] == b[2])
	byMin, byMax := min(b[1], b[3]), max(b[1], b[3])
	vx, hy := b[0], a[1]

	if a[0] <= vx && vx <= a[2] && byMin <= hy && hy <= byMax {
		overlaps[point{x: vx, y: hy}] = struct{}{}
	}
}

func Part1(in io.Reader) (int, error) {
	arrows, err := parseInput(in)
	if err != nil {
		return 0, err
	}

	overlaps := make(map[point]struct{}, len(arrows)*2)

	for i := 0; i < len(arrows)-1; i++ {
		a := arrows[i]
		switch a[4] {
		case DirectionHorizontal, DirectionVertical:
		default:
			continue
		}

		for j := i + 1; j < len(arrows); j++ {
			b := arrows[j]
			switch b[4] {
			case DirectionHorizontal, DirectionVertical:
			default:
				continue
			}

			overlapsPart1(a, b, overlaps)
		}
	}

	return len(overlaps), nil
}
