package int_util

func GcdInt[T IntType](a, b T) T {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
