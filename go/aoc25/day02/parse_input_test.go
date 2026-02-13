package day02

import (
	"strings"
	"testing"
)

func Test_parseInputGen(t *testing.T) {
	t.Run("test parseInputGen", func(t *testing.T) {
		reader := strings.NewReader("11-22,95-115,\n996-1012")
		var result []intInterval
		parseInputGen(reader)(func(data intInterval) bool {
			result = append(result, data)
			return true
		})
		if len(result) == 0 {
			t.Fatalf("expected to parse at least one interval")
		}

	})
}
