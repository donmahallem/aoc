package day05

import "testing"

func TestOverlapsPart1(t *testing.T) {
	cases := []struct {
		name string
		a, b coords
		want int
	}{
		{"horizontal and vertical overlap", coords{0, 0, 5, 0, DirectionHorizontal}, coords{2, -1, 2, 1, DirectionVertical}, 1},
		{"no overlap", coords{0, 0, 5, 0, DirectionHorizontal}, coords{6, -1, 6, 1, DirectionVertical}, 0},
		{"horizontal - horizontal overlap", coords{0, 0, 5, 0, DirectionHorizontal}, coords{3, 0, 7, 0, DirectionHorizontal}, 3},
		{"vertical - vertical overlap", coords{2, 0, 2, 5, DirectionVertical}, coords{2, 3, 2, 7, DirectionVertical}, 3},
		{"horizontal - vertical overlap", coords{1, 1, 4, 1, DirectionHorizontal}, coords{3, 0, 3, 3, DirectionVertical}, 1},
		{"horizontal - diagonal overlap", coords{0, 0, 5, 0, DirectionHorizontal}, coords{2, -1, 5, 6, DirectionDiagonalUp}, 0},
		{"vertical - diagonal overlap", coords{2, 0, 2, 5, DirectionVertical}, coords{1, 2, 4, 5, DirectionDiagonalUp}, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			overlaps := make(map[point]struct{})
			overlapsPart1(c.a, c.b, overlaps)
			if len(overlaps) != c.want {
				t.Errorf("expected %d overlaps, got %d", c.want, len(overlaps))
			}
		})
	}
}
