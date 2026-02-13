package day23_test

import (
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day23"
)

func TestStringifySequence(t *testing.T) {
	seq := day23.NodeHashList{day23.HashId([]byte{'a', 'b'}), day23.HashId([]byte{'z', 'u'})}
	expected := "ab,zu"
	if res := day23.StringifySequence(seq); res != expected {
		t.Errorf("String doesn't match. Expect '%s' not '%s'", expected, res)
	}
}
