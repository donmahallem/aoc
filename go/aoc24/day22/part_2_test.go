package day22_test

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/aoc24/day22"
)

func TestPart2(t *testing.T) {
	result := day22.Part2(strings.NewReader(testData))
	if result != 23 {
		t.Errorf(`Expected %d to be 23`, result)
	}
}
func TestEncodeSequence(t *testing.T) {
	testData := []uint32{1, 2, 3, 4}
	result := day22.EncodeSequence(&testData)
	if result != 7642 {
		t.Errorf(`Expected %d to be 23`, result)
	}
}

func TestCreatePatternDifferences(t *testing.T) {
	seed := uint32(123)
	patterns := day22.CreatePatternDifferences(&seed, 10)
	expected := []int8{-3, 6, -1, -1, 0, 2, -2, 0, -2}
	if !slices.Equal(*patterns, expected) {
		t.Errorf(`Expected %v to be %v`, patterns, expected)
	}
}

func TestCreatePatterns(t *testing.T) {
	seed := uint32(123)
	patterns := day22.CreatePatterns(&seed, 10)
	expected := []int8{3, 0, 6, 5, 4, 4, 6, 4, 4, 2}
	if !slices.Equal(*patterns, expected) {
		t.Errorf(`Expected %v to be %v`, patterns, expected)
	}
}
func TestCreatePatterns2(t *testing.T) {
	seed := uint32(123)
	cache := make(day22.CacheMap)
	day22.CreatePatterns2(&seed, 10, &cache)
	expectedValues := [][2]uint32{{53969, 4},
		{113174, 4},
		{65182, 6},
		{85394, 2}}
	if len(cache) != 6 {
		t.Errorf(`Expected %v to be 22`, cache)
	}
	for _, item := range expectedValues {
		if _, ok := cache[item[0]]; !ok {
			t.Errorf(`Expected %v to contain %d and be %d`, cache, item[0], item[1])
		}
	}
}
func BenchmarkPart2(b *testing.B) {
	for b.Loop() {
		day22.Part2(strings.NewReader(testData))
	}
}

var testDepths []int = []int{10, 100, 1000, 10000}

func BenchmarkCreatePatterns(b *testing.B) {
	seed := uint32(123)
	for _, testDepth := range testDepths {
		b.Run(fmt.Sprintf("test depth %d", testDepth), func(b *testing.B) {
			for b.Loop() {
				day22.CreatePatterns(&seed, testDepth)
			}
		})
	}
}
func BenchmarkCreatePatterns2(b *testing.B) {
	cache := make(day22.CacheMap)
	seed := uint32(123)
	for _, testDepth := range testDepths {
		b.Run(fmt.Sprintf("test depth %d", testDepth), func(b *testing.B) {
			for b.Loop() {
				day22.CreatePatterns2(&seed, testDepth, &cache)
			}
		})
	}
}

func BenchmarkCreatePatternDifferences(b *testing.B) {
	seed := uint32(123)
	for _, testDepth := range testDepths {
		b.Run(fmt.Sprintf("test depth %d", testDepth), func(b *testing.B) {
			for b.Loop() {
				day22.CreatePatternDifferences(&seed, testDepth)
			}
		})
	}
}
