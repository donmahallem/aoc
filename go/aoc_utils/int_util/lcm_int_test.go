package int_util_test

import (
	"testing"

	"github.com/donmahallem/aoc/go/aoc_utils/int_util"
)

func TestLcmInt(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"Basic LCM", 4, 6, 12},
		{"Prime numbers", 5, 7, 35},
		{"One is multiple of other", 10, 5, 10},
		{"Same numbers", 12, 12, 12},
		{"One is 1", 15, 1, 15},
		{"Negative numbers", -4, 6, 12},
		{"Both negative", -4, -6, 12},
		{"Large numbers", 21, 6, 42},
		{"Zero handling", 0, 5, 0}, // LCM(0, x) is typically 0
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if got := int_util.LcmInt(testCase.a, testCase.b); got != testCase.want {
				t.Errorf("LcmInt() = %v, want %v", got, testCase.want)
			}
		})
	}
}

func BenchmarkLcmInt(b *testing.B) {
	for b.Loop() {
		int_util.LcmInt(21, 6)
	}
}

func BenchmarkLcmInt_Large(b *testing.B) {
	for b.Loop() {
		int_util.LcmInt(123456, 789012)
	}
}
