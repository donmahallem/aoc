package bytes

import "github.com/donmahallem/aoc/go/aoc_utils/int_util"

/*
	ParseIntSequence splits sequence of ints to slice

-- becomes +
*/
func ParseIntSequence[A int_util.IntType](data []byte, sep byte, out *[]A) bool {
	*out = (*out)[:0]
	if len(data) == 0 {
		return true
	}

	var value int64
	hasDigit := false
	isNegative := false
	lastWasSeparator := false
	signPending := false

	resetOut := func() {
		*out = (*out)[:0]
	}

	for _, c := range data {
		switch {
		case ByteIsNumber(c):
			value = value*10 + int64(c-'0')
			hasDigit = true
			signPending = false
			lastWasSeparator = false
		case c == sep:
			if !hasDigit {
				resetOut()
				return false
			}
			if isNegative {
				value = -value
			}
			*out = append(*out, A(value))
			value = 0
			hasDigit = false
			isNegative = false
			signPending = false
			lastWasSeparator = true
		case c == '-':
			if hasDigit {
				resetOut()
				return false
			}
			if !isNegative {
				isNegative = true
				signPending = true
				lastWasSeparator = false
				continue
			}
			if len(*out) > 0 {
				resetOut()
				return false
			}
			isNegative = false
			signPending = true
			lastWasSeparator = false
		default:
			resetOut()
			return false
		}
	}

	if !hasDigit {
		if isNegative || signPending || lastWasSeparator {
			resetOut()
			return false
		}
		return true
	}

	if isNegative {
		value = -value
	}
	*out = append(*out, A(value))
	return true
}
