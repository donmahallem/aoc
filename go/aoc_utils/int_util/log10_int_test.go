package int_util_test

import (
	"testing"

	"github.com/donmahallem/aoc/go/aoc_utils/int_util"
)

func TestLog10Int(t *testing.T) {
	res := int_util.Log10Int(1)
	if res != 0 {
		t.Errorf(`Expected %d to match %d`, res, 0)
	}
	res = int_util.Log10Int(123456)
	if res != 5 {
		t.Errorf(`Expected %d to match %d`, res, 5)
	}
	res = int_util.Log10Int(253000)
	if res != 5 {
		t.Errorf(`Expected %d to match %d`, res, 5)
	}
	res = int_util.Log10Int(13456)
	if res != 4 {
		t.Errorf(`Expected %d to match %d`, res, 4)
	}
	res = int_util.Log10Int(1950139)
	if res != 6 {
		t.Errorf(`Expected %d to match %d`, res, 6)
	}
}
