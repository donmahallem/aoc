package int_util

func CountDigits[T IntType, O IntType](n T) O {
	digits := O(1)
	for n >= 10 {
		n /= 10
		digits++
	}
	return digits
}
