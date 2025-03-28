package day03_test

import (
	"io"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day03"
)

func TestNewMulReader(t *testing.T) {
	const sourceData = "mul(1,2)mul(1amul(2,3)"
	var data, _ = io.ReadAll(day03.NewMulReader(strings.NewReader(sourceData)))

	if i := strings.Compare(string(data), "8"); i != 0 {
		t.Errorf(`Expected %s to match %s`, string(data), "8")
	}
}

func TestPart1(t *testing.T) {
	const sourceData = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	var data = day03.Part1(strings.NewReader(sourceData))
	expected := 161
	if data != expected {
		t.Errorf(`Expected %d to match %d`, data, expected)
	}
}
