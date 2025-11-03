package day19_test

import (
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc23/day19"
)

var testData string = `px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}`

func TestParseInput(t *testing.T) {
	t.Run("testData1", func(t *testing.T) {
		const expected int = 11
		reader := strings.NewReader(testData)
		if res := day19.ParseInput(reader); len(res.Workflows) != expected {
			t.Errorf(`Expected %d to be %d`, len(res.Workflows), expected)
		}

	})

}

func TestPart1(t *testing.T) {
	t.Run("test block 2", func(t *testing.T) {

		reader := strings.NewReader(testData)
		res := day19.Part1(reader)
		if res != 19114 {
			t.Errorf(`Expected number of blocks to be 19114, got %d`, res)
		}
	})
}

func BenchmarkPart1(b *testing.B) {

	reader := strings.NewReader(testData)
	for b.Loop() {
		day19.Part1(reader)
		reader.Seek(0, 0)
	}
}
