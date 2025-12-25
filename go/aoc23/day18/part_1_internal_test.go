package day18

import (
	"strings"
	"testing"
)

func Test_calculateArea(t *testing.T) {
	t.Run("sample", func(t *testing.T) {
		ins, err := parseInput(strings.NewReader(sample), true)
		if err != nil {
			t.Errorf("parseInput() error = %v", err)
			return
		}
		if got := calculateArea(ins); got != 62 {
			t.Errorf("calculateArea() = %v, want %v", got, 62)
		}
	})
}
