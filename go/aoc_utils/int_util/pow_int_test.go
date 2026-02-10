package int_util

import (
	"testing"
)

func TestIntPow(t *testing.T) {
	tests := []struct {
		name     string
		base     int
		exp      int
		expected int
	}{
		{"2^0", 2, 0, 1},
		{"2^1", 2, 1, 2},
		{"2^2", 2, 2, 4},
		{"2^3", 2, 3, 8},
		{"2^10", 2, 10, 1024},
		{"3^2", 3, 2, 9},
		{"3^3", 3, 3, 27},
		{"1^100", 1, 100, 1},
		{"0^2", 0, 2, 0},
		{"0^0", 0, 0, 1}, // Implementation returns 1 for exp=0
		{"10^0", 10, 0, 1},
		{"-2^2", -2, 2, 4},
		{"-2^3", -2, 3, -8},
		{"10^5", 10, 5, 100000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IntPow(tt.base, tt.exp)
			if result != tt.expected {
				t.Errorf("IntPow(%d, %d) = %d; want %d", tt.base, tt.exp, result, tt.expected)
			}
		})
	}
}

func TestIntPow_Generics(t *testing.T) {
	t.Run("int64", func(t *testing.T) {
		got := IntPow(int64(2), int64(10))
		want := int64(1024)
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("uint", func(t *testing.T) {
		got := IntPow(uint(2), uint(10))
		want := uint(1024)
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("check exponent zero", func(t *testing.T) {
		got := IntPow(int(5), int(0))
		want := int(1)
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("int8_overflow", func(t *testing.T) {
		got := IntPow(int8(2), int8(10))
		want := int8(0) // 2^10 = 1024 = 0 mod 256
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
