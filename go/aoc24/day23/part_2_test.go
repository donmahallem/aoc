package day23_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day23"
	"github.com/donmahallem/aoc/go/test_utils"
)

func TestStringifySequence(t *testing.T) {
	seq := day23.NodeHashList{day23.HashId([]byte{'a', 'b'}), day23.HashId([]byte{'z', 'u'})}
	expected := "ab,zu"
	if res := day23.StringifySequence(seq); res != expected {
		t.Errorf("String doesn't match. Expect '%s' not '%s'", expected, res)
	}
}

func TestPart2(t *testing.T) {
	t.Run("test sample 1", func(t *testing.T) {
		expected := "co,de,ka,ta"
		reader := strings.NewReader(testData)
		result := day23.Part2(reader)
		if result != expected {
			t.Errorf(`Expected number of blocks to be "%s", got "%s"`, expected, result)
		}
	})

	t.Run("test real data", func(t *testing.T) {
		result, ok := test_utils.TestFullDataForDate(t, 24, 23, day23.Part2)
		expected := "az,ed,hz,it,ld,nh,pc,td,ty,ux,wc,yg,zz"
		if !ok || result != expected {
			t.Errorf(`Expected "%s" to be "%s"`, result, expected)
		}
	})
}
func BenchmarkPart2(b *testing.B) {
	b.Run("benchmark sample data", func(b *testing.B) {

		reader := strings.NewReader(testData)
		for b.Loop() {
			day23.Part2(reader)
			reader.Seek(0, 0)
		}
	})

	b.Run("benchmark full data", func(b *testing.B) {
		test_utils.BenchmarkFullDataForDate(b, 24, 23, day23.Part2)
	})
}
