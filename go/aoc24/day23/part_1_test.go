package day23_test

import (
	"math/rand/v2"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day23"
)

const testData string = `kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`

func BenchmarkLookup(b *testing.B) {
	test_2 := [][2]byte{}
	for range 10000 {
		test_2 = append(test_2, [2]byte{byte(rand.UintN(255)), byte(rand.UintN(255))})
	}

	b.Run("LookupA", func(b *testing.B) {
		test := make(map[[2]byte]bool)
		for _, entry := range test_2 {
			test[entry] = true
		}
		for b.Loop() {
			day23.LookupA(test, test_2[0])
		}
	})
	b.Run("LookupB", func(b *testing.B) {
		test := make(map[uint]bool)
		test_keys := make([]uint, 0)
		for _, entry := range test_2 {
			test_key := uint(entry[0]*255 + entry[1])
			test_keys = append(test_keys, test_key)
			test[test_key] = true
		}
		for b.Loop() {
			day23.LookupB(test, test_keys[0])
		}
	})
	b.Run("HashId", func(b *testing.B) {
		testKey := []byte{15, 125}
		for b.Loop() {
			day23.HashId(testKey)
		}
	})
	b.Run("t", func(b *testing.B) {
		for b.Loop() {
			day23.ParseInput(strings.NewReader(testData))
		}
	})
	b.Run("t2", func(b *testing.B) {
		for b.Loop() {
			day23.ParseInputMap(strings.NewReader(testData))
		}
	})
}

func TestFindTriplets(t *testing.T) {
	points := day23.ParseInputMap(strings.NewReader(testData))
	data := day23.FindTriplets(points)
	if data != 7 {
		t.Errorf(`Expected %d to match 7`, data)
	}
}

func TestPart1(t *testing.T) {
	points := day23.Part1(strings.NewReader(testData))
	if points != 7 {
		t.Errorf(`Expected %d to match 7`, points)
	}
}

func TestStartsWithT(t *testing.T) {
	testId := day23.HashId([]byte{'b', 'a'})
	if day23.StartsWithT(&testId) {
		t.Errorf(`Expected %d to not start with t. %s`, testId, string(*day23.UnhashId(&testId)))
	}
	testId = day23.HashId([]byte{'t', 'a'})
	if !day23.StartsWithT(&testId) {
		t.Errorf(`Expected %d to not start with t. %s`, testId, string(*day23.UnhashId(&testId)))
	}
}

func BenchmarkPart1(b *testing.B) {
	for b.Loop() {
		day23.Part1(strings.NewReader(testData))
	}
}
