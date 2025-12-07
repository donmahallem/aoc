package int_util

func AbsInt[T IntType](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
