package day24

type hail[A int64 | float64] struct {
	Px, Py, Pz A
	Vx, Vy, Vz A
}

type floatHail = hail[float64]
type intHail = hail[int64]
