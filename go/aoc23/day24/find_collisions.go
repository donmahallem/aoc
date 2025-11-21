package day24

// Cramer's rule to solve for intersection point in x and y
func cramerXY(a, b floatHail) (float64, float64, bool) {
	det := a.Vx*b.Vy - a.Vy*b.Vx
	if det == 0 {
		return 0, 0, false
	}
	t := ((b.Px-a.Px)*b.Vy - (b.Py-a.Py)*b.Vx) / det
	s := ((a.Px-b.Px)*a.Vy - (a.Py-b.Py)*a.Vx) / (-det)
	return t, s, true
}

// checks for collisions between two hail stones with in the bounds of min and max in x and y
func findCollisions(hails []floatHail, min, max float64) int {
	collisionCount := 0
	for i := range len(hails) - 1 {
		h1 := hails[i]
		for j := i + 1; j < len(hails); j++ {
			h2 := hails[j]

			facT, facS, ok := cramerXY(h1, h2)
			if !ok || facT < 0 || facS < 0 {
				// parallel or passed each other "in the past"
				continue
			}
			posX := h1.Px + facT*h1.Vx
			posY := h1.Py + facT*h1.Vy
			if posX < min || posX > max || posY < min || posY > max {
				continue
			}
			collisionCount++
		}
		// check if
	}
	return collisionCount
}
