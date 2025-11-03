package day23_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day23"
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
	b.Run("HashId", func(b *testing.B) {
		testKey := []byte{15, 125}
		for b.Loop() {
			day23.HashId(testKey)
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
func TestParseInputMap(t *testing.T) {
	points := day23.ParseInputMap(strings.NewReader(testData))
	var keyTa []byte = []byte{'t', 'a'}
	var hashTa day23.NodeHash = day23.HashId(keyTa)
	if res := len((*points)[hashTa]); res != 0 {
		t.Errorf(`Expected %v %d to have length %d not 0`, keyTa, hashTa, res)
	}
	keyTa = []byte{'c', 'o'}
	hashTa = day23.HashId(keyTa)
	if res := len((*points)[hashTa]); res != 4 {
		t.Errorf(`Expected %v %d to have length %d not 4`, keyTa, hashTa, res)
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
