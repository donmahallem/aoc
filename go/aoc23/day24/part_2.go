package day24

import (
	"fmt"
	"io"
	"math"
)

type solverMatrix = [6][7]float64
type matrixA = [6][6]float64
type vectorB = [6]float64

func solveLinear(A *matrixA, b *vectorB, s *solverMatrix) (*floatHail, error) {
	n := len(A)
	// Augment A with b
	for i := range 6 {
		for j := range 6 {
			s[i][j] = A[i][j]
		}
		s[i][n] = b[i]
	}

	for i := range n {
		// Pivot
		pivotRow := i
		for k := i + 1; k < n; k++ {
			if math.Abs(s[k][i]) > math.Abs(s[pivotRow][i]) {
				pivotRow = k
			}
		}
		s[i], s[pivotRow] = s[pivotRow], s[i]

		if math.Abs(s[i][i]) < 1e-9 {
			return nil, fmt.Errorf("singular matrix")
		}

		// Normalize row i
		div := s[i][i]
		for j := i; j <= n; j++ {
			s[i][j] /= div
		}

		// Eliminate other rows
		for k := range n {
			if k != i {
				factor := s[k][i]
				for j := i; j <= n; j++ {
					s[k][j] -= factor * s[i][j]
				}
			}
		}
	}

	return &floatHail{
		Px: s[0][n],
		Py: s[1][n],
		Pz: s[2][n],
		Vx: s[3][n],
		Vy: s[4][n],
		Vz: s[5][n],
	}, nil
}

func buildPairLinearSystem(A *matrixA, b *vectorB, hailStones []floatHail, idxA, idxB, idxC int) {

	fillBlock := func(rowOffset, i, j int) {

		dvx := hailStones[j].Vx - hailStones[i].Vx
		dvy := hailStones[j].Vy - hailStones[i].Vy
		dvz := hailStones[j].Vz - hailStones[i].Vz
		drx := hailStones[i].Px - hailStones[j].Px
		dry := hailStones[i].Py - hailStones[j].Py
		drz := hailStones[i].Pz - hailStones[j].Pz

		// Matrix A block
		// Row 0
		A[rowOffset+0][0] = 0
		A[rowOffset+0][1] = -dvz
		A[rowOffset+0][2] = dvy
		A[rowOffset+0][3] = 0
		A[rowOffset+0][4] = -drz
		A[rowOffset+0][5] = dry

		// Row 1
		A[rowOffset+1][0] = dvz
		A[rowOffset+1][1] = 0
		A[rowOffset+1][2] = -dvx
		A[rowOffset+1][3] = drz
		A[rowOffset+1][4] = 0
		A[rowOffset+1][5] = -drx

		// Row 2
		A[rowOffset+2][0] = -dvy
		A[rowOffset+2][1] = dvx
		A[rowOffset+2][2] = 0
		A[rowOffset+2][3] = -dry
		A[rowOffset+2][4] = drx
		A[rowOffset+2][5] = 0

		// Vector b block: (pi x vi) - (pj x vj)
		cix := hailStones[i].Py*hailStones[i].Vz - hailStones[i].Pz*hailStones[i].Vy
		ciy := hailStones[i].Pz*hailStones[i].Vx - hailStones[i].Px*hailStones[i].Vz
		ciz := hailStones[i].Px*hailStones[i].Vy - hailStones[i].Py*hailStones[i].Vx

		cjx := hailStones[j].Py*hailStones[j].Vz - hailStones[j].Pz*hailStones[j].Vy
		cjy := hailStones[j].Pz*hailStones[j].Vx - hailStones[j].Px*hailStones[j].Vz
		cjz := hailStones[j].Px*hailStones[j].Vy - hailStones[j].Py*hailStones[j].Vx

		b[rowOffset+0] = cix - cjx
		b[rowOffset+1] = ciy - cjy
		b[rowOffset+2] = ciz - cjz
	}

	fillBlock(0, idxB, idxA)
	fillBlock(3, idxA, idxC)
}

func Part2(in io.Reader) int64 {
	inp := parseInput[float64](in)
	if len(inp) < 3 {
		return -1
	}
	A := matrixA{}
	S := solverMatrix{}
	for i := range 6 {
		A[i] = [6]float64{}
		S[i] = [7]float64{}
	}
	b := vectorB{}

	var res *floatHail
	var err error

	aAddress := &A
	bAddress := &b
	sAddress := &S
	// brute for non linear system by trying all triplets
	for i := 0; i < len(inp)-2; i++ {
		buildPairLinearSystem(aAddress, bAddress, inp, i, i+1, i+2)
		res, err = solveLinear(aAddress, bAddress, sAddress)
		if err == nil {
			break
		}
	}

	if res == nil {
		fmt.Println("Could not find non-singular set of hailstones")
		return -1
	}

	// Round the result to nearest int64
	rx := int64(math.Round(res.Px))
	ry := int64(math.Round(res.Py))
	rz := int64(math.Round(res.Pz))

	return rx + ry + rz
}
