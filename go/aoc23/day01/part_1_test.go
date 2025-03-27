package day01_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc23/day01"
)

const testData string = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

func TestOutOfBoundsShouldBeInside(t *testing.T) {
	expected := 142
	reader := strings.NewReader(testData)
	if res := day01.ParseFile(reader); res != expected {
		t.Errorf(`Expected %v to match %v`, res, expected)
	}
}
