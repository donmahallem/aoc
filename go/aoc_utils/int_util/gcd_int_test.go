package int_util_test

import (
	"testing"

	"github.com/donmahallem/aoc/go/aoc_utils/int_util"
)

func TestGcdInt(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"Basic GCD", 48, 18, 6},
		{"Prime numbers", 17, 13, 1},
		{"One is multiple of other", 54, 6, 6},
		{"Same numbers", 10, 10, 10},
		{"One is zero", 5, 0, 5},
		{"Both zero", 0, 0, 0},
		{"Large numbers", 1071, 462, 21},
		{"One is 1", 100, 1, 1},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if got := int_util.GcdInt(testCase.a, testCase.b); got != testCase.want {
				t.Errorf("GcdInt() = %v, want %v", got, testCase.want)
			}
		})
	}
}

func BenchmarkGcdInt(b *testing.B) {
	for b.Loop() {
		int_util.GcdInt(1071, 462)
	}
}

func BenchmarkGcdInt_Large(b *testing.B) {
	for b.Loop() {
		int_util.GcdInt(1234567890, 462)
	}
}
