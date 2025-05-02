package day21_test

import (
	"testing"

	"github.com/donmahallem/aoc/aoc24/day21"
)

func TestWalkDirectional(t *testing.T) {
	t.Run("A to 0 depth 1", func(t *testing.T) {
		start := day21.DIRECTIONAL_UP
		end := day21.DIRECTIONAL_A
		cache := make(day21.Cache)
		result := day21.WalkDirectional(&start, &end, 0, 1, &cache)
		if result != 2 {
			t.Errorf(`Expected %d to match 2`, result)
		}
	})
	t.Run("A to 0 depth 2", func(t *testing.T) {
		start := day21.DIRECTIONAL_UP
		end := day21.DIRECTIONAL_A
		cache := make(day21.Cache)
		result := day21.WalkDirectional(&start, &end, 0, 2, &cache)
		if result != 4 {
			t.Errorf(`Expected %d to match 2`, result)
		}
	})
	t.Run("A to 0 depth 3", func(t *testing.T) {
		start := day21.DIRECTIONAL_UP
		end := day21.DIRECTIONAL_A
		cache := make(day21.Cache)
		result := day21.WalkDirectional(&start, &end, 0, 3, &cache)
		if result != 10 {
			t.Errorf(`Expected %d to match 2`, result)
		}
	})
	t.Run("A to UP depth 1", func(t *testing.T) {
		start := day21.DIRECTIONAL_A
		end := day21.DIRECTIONAL_UP
		cache := make(day21.Cache)
		result := day21.WalkDirectional(&start, &end, 0, 1, &cache)
		if result != 2 {
			t.Errorf(`Expected %d to match 2`, result)
		}
	})
	t.Run("A to UP depth 2", func(t *testing.T) {
		start := day21.DIRECTIONAL_A
		end := day21.DIRECTIONAL_UP
		cache := make(day21.Cache)
		result := day21.WalkDirectional(&start, &end, 0, 2, &cache)
		if result != 8 {
			t.Errorf(`Expected %d to match 8`, result)
		}
	})
	t.Run("A to UP depth 3", func(t *testing.T) {
		start := day21.DIRECTIONAL_A
		end := day21.DIRECTIONAL_UP
		cache := make(day21.Cache)
		result := day21.WalkDirectional(&start, &end, 0, 3, &cache)
		if result != 18 {
			t.Errorf(`Expected %d to match 18`, result)
		}
	})
}
