package day14_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day14"
)

func TestStep(t *testing.T) {
	tests := []day14.Robot{*day14.NewRobot(5, 1, 3, -3)}
	expected := *day14.NewRobot(8, 5, 3, -3)
	width, height := 11, 7
	day14.Step(tests, width, height)
	if tests[0] != expected {
		t.Errorf(`Expected %v to match %v`, tests[0], expected)
	}
}

func TestPart2(t *testing.T) {
	if result := day14.Part2(strings.NewReader(testData)); result != 1 {
		t.Errorf(`Expected %d to contain %d`, result, 1)
	}
}
func BenchmarkPart2(b *testing.B) {
	data := strings.NewReader(testData)
	for b.Loop() {
		data.Seek(0, io.SeekStart)
		day14.Part2(data)
	}
}
