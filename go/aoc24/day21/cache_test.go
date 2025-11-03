package day21_test

import (
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day21"
)

func TestHashId(t *testing.T) {
	p1 := day21.Point{X: 1, Y: 1}
	p2 := day21.Point{X: 3, Y: 3}
	depth := uint8(5)
	test := day21.HashId(&p1, &p2, depth)
	if test != 1375 {
		t.Errorf(`Expected %d to match 253`, test)
	}

	test2 := day21.HashId2(&p1, &p2, depth)
	if test != test2 {
		t.Errorf(`Expected %d to match %d`, test, test2)
	}
}

func BenchmarkHashId(b *testing.B) {
	p1 := day21.Point{X: 1, Y: 1}
	p2 := day21.Point{X: 3, Y: 3}
	depth := uint8(5)
	for b.Loop() {
		day21.HashId(&p1, &p2, depth)
	}
}

func BenchmarkHashId2(b *testing.B) {
	p1 := day21.Point{X: 1, Y: 1}
	p2 := day21.Point{X: 3, Y: 3}
	depth := uint8(5)
	for b.Loop() {
		day21.HashId2(&p1, &p2, depth)
	}
}
