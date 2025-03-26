package maina

import (
	"fmt"
)

// Function to check if three points form a corner
func isCorner(x1, y1, x2, y2, x3, y3 int) bool {
	// Calculate the slopes of the two segments
	dx1, dy1 := x2-x1, y2-y1
	dx2, dy2 := x3-x2, y3-y2
	return !(dx1*dy2 == dx2*dy1) // Not collinear
}

func countCorners(grid [][]int) int {
	corners := 0
	n := len(grid)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if isCorner(grid[i][0], grid[i][1], grid[j][0], grid[j][1], grid[k][0], grid[k][1]) {
					corners++
				}
			}
		}
	}
	return corners
}

func main() {
	// Example grid coordinates
	grid := [][]int{
		{0, 0},
		{1, 0},
		{1, 1},
		{0, 1},
	}

	fmt.Println("Number of corners:", countCorners(grid))
}
