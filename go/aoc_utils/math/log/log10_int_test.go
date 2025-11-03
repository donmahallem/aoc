package log_test

import (
	"testing"

	"github.com/donmahallem/aoc/go/aoc_utils/math/log"
)

func TestLog10Int(t *testing.T) {
	res := log.Log10Int(1)
	if res != 1 {
		t.Errorf(`Expected %d to match %d`, res, 1)
	}
	res = log.Log10Int(123456)
	if res != 6 {
		t.Errorf(`Expected %d to match %d`, res, 6)
	}
	res = log.Log10Int(253000)
	if res != 6 {
		t.Errorf(`Expected %d to match %d`, res, 6)
	}
	res = log.Log10Int(13456)
	if res != 5 {
		t.Errorf(`Expected %d to match %d`, res, 5)
	}
	res = log.Log10Int(1950139)
	if res != 7 {
		t.Errorf(`Expected %d to match %d`, res, 7)
	}
}
