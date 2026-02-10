package day22_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/aoc24/day22"
)

const testData2 string = `1
2
3
2024`

func TestPart2(t *testing.T) {
	result, err := day22.Part2(strings.NewReader(testData2))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
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

func TestCreatePatterns(t *testing.T) {
	t.Run("test simple", func(t *testing.T) {
		seed := uint32(123)
		cache := make(day22.CacheMap)
		day22.CreatePatterns(&seed, 10, &cache)
		expectedValues := [][2]uint32{{53969, 4},
			{113174, 4},
			{65182, 6},
			{85394, 2},
			{85394, 2}}
		if len(cache) != 6 {
			t.Errorf(`Expected %v to be 22`, cache)
		}
		for _, item := range expectedValues {
			if _, ok := cache[item[0]]; !ok {
				t.Errorf(`Expected %v to contain %d and be %d`, cache, item[0], item[1])
			}
		}
		day22.CreatePatterns(&seed, 10, &cache)
		if len(cache) != 6 {
			t.Errorf(`Expected %v to be 22`, cache)
		}
		for _, item := range expectedValues {
			if _, ok := cache[item[0]]; !ok {
				t.Errorf(`Expected %v to contain %d and be %d`, cache, item[0], item[1]*2)
			}
		}
	})
	t.Run("test samples", func(t *testing.T) {
		seed := uint32(1)
		cache := make(day22.CacheMap)
		day22.CreatePatterns(&seed, 2000, &cache)
		lookup := uint32(59027) //sequence -2,1,-1,3
		if result := cache[lookup]; result.Value != 7 {
			t.Errorf(`Expected %d to be 7 at %d`, result, lookup)
		}
		seed = 2
		cache = make(day22.CacheMap)
		day22.CreatePatterns(&seed, 2000, &cache)
		if result := cache[lookup]; result.Value != 7 {
			t.Errorf(`Expected %d to be 7 at %d`, result, lookup)
		}
		seed = 3
		cache = make(day22.CacheMap)
		day22.CreatePatterns(&seed, 2000, &cache)
		if result := cache[lookup]; result.Value != 0 {
			t.Errorf(`Expected %d to be 0 at %d`, result, lookup)
		}
		seed = 2024
		cache = make(day22.CacheMap)
		day22.CreatePatterns(&seed, 2000, &cache)
		if result := cache[lookup]; result.Value != 9 {
			t.Errorf(`Expected %d to be 9 at %d`, result, lookup)
		}
	})
	t.Run("test combine", func(t *testing.T) {
		seed := uint32(1)
		cache := make(day22.CacheMap)
		day22.CreatePatterns(&seed, 2000, &cache)
		lookup := uint32(59027) //sequence -2,1,-1,3
		if result := cache[lookup]; result.Value != 7 {
			t.Errorf(`Expected %d to be 7 at %d`, result, lookup)
		}
		seed = 2
		day22.CreatePatterns(&seed, 2000, &cache)
		if result := cache[lookup]; result.Value != 14 {
			t.Errorf(`Expected %d to be 14 at %d`, result, lookup)
		}
		seed = 3
		day22.CreatePatterns(&seed, 2000, &cache)
		if result := cache[lookup]; result.Value != 14 {
			t.Errorf(`Expected %d to be 14 at %d`, result, lookup)
		}
		seed = 2024
		day22.CreatePatterns(&seed, 2000, &cache)
		if result := cache[lookup]; result.Value != 23 {
			t.Errorf(`Expected %d to be 23 at %d`, result, lookup)
		}
	})
}
func BenchmarkPart2(b *testing.B) {
	for b.Loop() {
		day22.Part2(strings.NewReader(testData2))
	}
}

var testDepths []int = []int{10, 100, 1000, 10000}

func BenchmarkCreatePatterns(b *testing.B) {
	cache := make(day22.CacheMap)
	seed := uint32(123)
	for _, testDepth := range testDepths {
		b.Run(fmt.Sprintf("test depth %d", testDepth), func(b *testing.B) {
			for b.Loop() {
				day22.CreatePatterns(&seed, testDepth, &cache)
			}
		})
	}
}
