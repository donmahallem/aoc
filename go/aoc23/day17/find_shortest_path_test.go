package day17

import (
	"strings"
	"testing"
)

func Test_findShortestPath(t *testing.T) {
	t.Run("sample", func(t *testing.T) {
		f, err := parseInput(strings.NewReader(sample))
		if err != nil {
			t.Errorf("parseInput() error = %v", err)
			return
		}
		if got := findShortestPath(f, 0, 3); got != 102 {
			t.Errorf("findShortestPath() = %v, want %v", got, 102)
		}
	})
}
