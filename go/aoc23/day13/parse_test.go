package day13

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed sample.txt
var testData string

func Test_parseInput(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		const expected int = 2
		reader := strings.NewReader(testData)
		res, err := parseInput(reader)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if len(res) != expected {
			t.Errorf(`Expected %d to be %d`, len(res), expected)
		}

	})
	t.Run("test block 1", func(t *testing.T) {

		reader := strings.NewReader(testData)
		res, err := parseInput(reader)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		expectedRows := []int{
			0b11001101,
			0b10110100,
			0b100000011,
			0b100000011,
			0b010110100,
			0b011001100,
			0b010110101,
		}
		expectedColumns := []int{
			0b1001101, //0
			0b0001100, //1
			0b1110011, //2
			0b0100001, //3
			0b1010010, //4
			0b1010010, //5
			0b0100001, //6
			0b1110011, //7
			0b0001100, //8
		}
		for rowIdx := range expectedRows {
			if res[0].Rows[rowIdx] != expectedRows[rowIdx] {
				t.Errorf(`Expected row %d to be %b, got %b`, rowIdx, expectedRows[rowIdx], res[0].Rows[rowIdx])
			}
		}
		for colIdx := range expectedColumns {
			if res[0].Cols[colIdx] != expectedColumns[colIdx] {
				t.Errorf(`Expected row %d to be %b, got %b`, colIdx, expectedColumns[colIdx], res[0].Cols[colIdx])
			}
		}
	})

}
