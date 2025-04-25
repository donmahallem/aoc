package day22_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day22"
)

const testData string = `1
10
100
2024`

var testCases []int = []int{
	123,
	15887950,
	16495136,
	527345,
	704524,
	1553684,
	12683156,
	11100544,
	12249484,
	7753432,
	5908254}

func TestStep(t *testing.T) {
	for idx := range len(testCases) - 1 {
		t.Run(fmt.Sprintf("Expect %d to become %d", testCases[idx], testCases[idx+1]), func(t *testing.T) {
			points := day22.Step(uint32(testCases[idx]))
			if points != uint32(testCases[idx+1]) {
				t.Errorf(`Expected %d to match %d`, points, testCases[idx+1])
			}
		})
	}
}

func TestPart1(t *testing.T) {
	points := day22.Part1(strings.NewReader(testData))

	if points != 37327623 {
		t.Errorf(`Expected %d to match 37327623`, points)
	}
}

func BenchmarkStep(b *testing.B) {
	for idx := range len(testCases) - 1 {
		b.Run(fmt.Sprintf("Expect %d to become %d", testCases[idx], testCases[idx+1]), func(b *testing.B) {
			testValue := uint32(testCases[idx])
			for b.Loop() {
				day22.Step(testValue)
			}
		})
	}
}

func BenchmarkAddUpSecrets(b *testing.B) {
	secrets := day22.ParseInput(strings.NewReader(testData))
	for b.Loop() {
		day22.AddUpSecrets(secrets)
	}
}

func BenchmarkPart1(b *testing.B) {
	for b.Loop() {
		day22.Part1(strings.NewReader(testData))
	}
}
