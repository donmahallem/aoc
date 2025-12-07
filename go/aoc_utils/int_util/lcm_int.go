package int_util

func LcmInt[T IntType](a, b T) T {
	a = AbsInt(a)
	b = AbsInt(b)
	return AbsInt(a*b) / GcdInt(a, b)
}
