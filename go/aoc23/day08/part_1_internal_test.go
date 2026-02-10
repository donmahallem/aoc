package day08

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed sample1.txt
var testData string

//go:embed sample2.txt
var testData2 string

func TestParseInput(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		reader := strings.NewReader(testData)
		res, err := parseInput(reader)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
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
		res, err := parseInput(reader)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
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
