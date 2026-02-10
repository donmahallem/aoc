package day16

import (
	"strings"
	"testing"
)

var testData string = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

func TestParseInput(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		const expected int = 100
		reader := strings.NewReader(testData)
		res, err := parseInput(reader)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if len(res.Cells) != expected {
			t.Errorf(`Expected %d to be %d`, len(res.Cells), expected)
		}

	})

}
