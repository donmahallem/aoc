package day21

import (
	"testing"
)

func TestWalkDirectional(t *testing.T) {
	t.Run("A to 0 depth 1", func(t *testing.T) {
		start := directionalPosition_Up
		end := directionalPosition_A
		cache := make(cache)
		result := WalkDirectional(&start, &end, 0, 1, &cache)
		if result != 2 {
			t.Errorf(`Expected %d to match 2`, result)
		}
	})
	t.Run("A to 0 depth 2", func(t *testing.T) {
		start := directionalPosition_Up
		end := directionalPosition_A
		cache := make(cache)
		result := WalkDirectional(&start, &end, 0, 2, &cache)
		if result != 4 {
			t.Errorf(`Expected %d to match 2`, result)
		}
	})
	t.Run("A to 0 depth 3", func(t *testing.T) {
		start := directionalPosition_Up
		end := directionalPosition_A
		cache := make(cache)
		result := WalkDirectional(&start, &end, 0, 3, &cache)
		if result != 10 {
			t.Errorf(`Expected %d to match 2`, result)
		}
	})
	t.Run("A to UP depth 1", func(t *testing.T) {
		start := directionalPosition_A
		end := directionalPosition_Up
		cache := make(cache)
		result := WalkDirectional(&start, &end, 0, 1, &cache)
		if result != 2 {
			t.Errorf(`Expected %d to match 2`, result)
		}
	})
	t.Run("A to UP depth 2", func(t *testing.T) {
		start := directionalPosition_A
		end := directionalPosition_Up
		cache := make(cache)
		result := WalkDirectional(&start, &end, 0, 2, &cache)
		if result != 8 {
			t.Errorf(`Expected %d to match 8`, result)
		}
	})
	t.Run("A to UP depth 3", func(t *testing.T) {
		start := directionalPosition_A
		end := directionalPosition_Up
		cache := make(cache)
		result := WalkDirectional(&start, &end, 0, 3, &cache)
		if result != 18 {
			t.Errorf(`Expected %d to match 18`, result)
		}
	})
}
