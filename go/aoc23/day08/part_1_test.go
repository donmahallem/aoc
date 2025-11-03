package day08_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day08"
)

const testData string = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

const testData2 string = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

func TestParseInput(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		reader := strings.NewReader(testData)
		res := day08.ParseInput(reader)
		if len(res.Instructions) != 2 {
			t.Errorf(`Expected 2 instructions not %d`, len(res.Instructions))
		}
		if res.Instructions[0] != true || res.Instructions[1] != false {
			t.Errorf(`Expected the instructions to be true,false`)
		}
		if res.Start == nil {
			t.Errorf(`Start should not be nil`)
		}
		if res.End == nil {
			t.Errorf(`End should not be nil`)
		}
		if res.Start.Right.Left != res.End {
			t.Errorf(`End should be accesible by traversing Right->Left from Start`)
		}
		if res.End != res.End.Left || res.End != res.End.Right {
			t.Error("Expected res.End to be equal to res.End.Left and Right")
		}
	})
	t.Run("testData2", func(t *testing.T) {
		reader := strings.NewReader(testData2)
		res := day08.ParseInput(reader)
		if len(res.Instructions) != 3 {
			t.Errorf(`Expected 2 instructions not %d`, len(res.Instructions))
		}
		if res.Instructions[0] != false || res.Instructions[1] != false || res.Instructions[2] != true {
			t.Errorf(`Expected the instructions to be false,false,true`)
		}
		if res.Start == nil {
			t.Errorf(`Start should not be nil`)
		}
		if res.End == nil {
			t.Errorf(`End should not be nil`)
		}
		if res.Start.Left.Left.Right.Left.Left.Right != res.End {
			t.Errorf(`End should be accesible by traversing Left.Left.Right.Left.Left.Right from Start`)
		}
		if res.End != res.End.Left || res.End != res.End.Right {
			t.Error("Expected res.End to be equal to res.End.Left and Right")
		}
	})
}

func TestPart1(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		const expected int = 2
		reader := strings.NewReader(testData)
		if res := day08.Part1(reader); res != expected {
			t.Errorf(`Expected %d to be %d`, res, expected)
		}
	})
	t.Run("testData2", func(t *testing.T) {
		const expected int = 6
		reader := strings.NewReader(testData2)
		if res := day08.Part1(reader); res != expected {
			t.Errorf(`Expected %d to be %d`, res, expected)
		}
	})
}

func BenchmarkPart1(b *testing.B) {
	reader := strings.NewReader(testData)
	for b.Loop() {
		reader.Seek(0, io.SeekStart)
		day08.Part1(reader)
	}
}
