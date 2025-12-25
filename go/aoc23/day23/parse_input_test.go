package day23

import (
	_ "embed"
	"fmt"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/go/test_utils"
)

func Test_parseInput3(t *testing.T) {
	t.Run("respect slope", func(t *testing.T) {
		t.Run("test sample", func(t *testing.T) {
			reader := strings.NewReader(testData)
			result, w, h, err := parseInput(reader, true)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if w != 23 || h != 23 {
				t.Errorf(`Expected width and height to be 23, got %d and %d`, w, h)
			}
			testPoints := map[[2]int]cell{
				{1, 0}:   dirExitBottom | dirEntryBottom,
				{1, 1}:   dirExitRight | dirExitTop | dirEntryRight | dirEntryTop,
				{2, 1}:   dirExitLeft | dirExitRight | dirEntryLeft | dirEntryRight,
				{9, 3}:   dirExitRight | dirExitBottom | dirEntryBottom,
				{10, 3}:  dirExitRight | dirEntryLeft | dirEntryRight,
				{11, 3}:  dirExitRight | dirExitBottom | dirEntryLeft | dirExitLeft,
				{3, 3}:   dirExitRight | dirExitBottom | dirEntryRight,
				{3, 4}:   dirExitBottom | dirEntryTop | dirEntryBottom,
				{3, 5}:   dirExitBottom | dirExitRight | dirEntryTop | dirExitTop,
				{4, 5}:   dirExitRight | dirEntryLeft | dirEntryRight,
				{13, 19}: dirEntryLeft | dirExitLeft | dirExitRight | dirEntryTop | dirExitTop,
			}
			for idx, expectedVal := range testPoints {
				pos := idx[1]*w + idx[0]
				if result[pos] != expectedVal {
					t.Errorf(`Expected index %d (%d,%d) to be %08b, got %08b`, idx, idx[0], idx[1], expectedVal, result[pos])
				}
			}
		})
	})

	t.Run("don't respect slope", func(t *testing.T) {

		t.Run("test sample", func(t *testing.T) {
			reader := strings.NewReader(testData)
			result, w, h, err := parseInput(reader, false)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if w != 23 || h != 23 {
				t.Errorf(`Expected width and height to be 23, got %d and %d`, w, h)
			}
			testPoints := map[[2]int]cell{
				{1, 0}:  dirExitBottom | dirEntryBottom,
				{1, 1}:  dirExitRight | dirExitTop | dirEntryRight | dirEntryTop,
				{2, 1}:  dirExitLeft | dirExitRight | dirEntryLeft | dirEntryRight,
				{9, 3}:  dirExitRight | dirExitBottom | dirEntryBottom | dirEntryRight,
				{10, 3}: dirExitRight | dirEntryLeft | dirEntryRight | dirExitLeft,
				{11, 3}: dirExitRight | dirExitBottom | dirEntryLeft | dirExitLeft | dirEntryRight | dirEntryBottom,
				{3, 3}:  dirExitRight | dirExitBottom | dirEntryRight | dirEntryBottom,
				{3, 4}:  dirExitBottom | dirEntryTop | dirEntryBottom | dirExitTop,
				{3, 5}:  dirExitBottom | dirExitRight | dirEntryTop | dirExitTop | dirEntryBottom | dirEntryRight,
				{4, 5}:  dirExitRight | dirEntryLeft | dirEntryRight | dirExitLeft,
			}
			for idx, expectedVal := range testPoints {
				pos := idx[1]*w + idx[0]
				if result[pos] != expectedVal {
					t.Errorf(`Expected index %d (%d,%d) to be %08b, got %08b`, idx, idx[0], idx[1], expectedVal, result[pos])
				}
			}
		})
		t.Run("test cross sample", func(t *testing.T) {
			testData := "#.#\n^.<\n#^#"
			expected := []cell{
				0, dirExitBottom | dirEntryBottom, 0,
				dirEntryRight | dirExitRight, (dirEntryAll | dirExitAll), dirExitLeft | dirEntryLeft,
				0, dirExitTop | dirEntryTop, 0,
			}
			reader := strings.NewReader(testData)
			result, w, h, err := parseInput(reader, false)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if w != 3 || h != 3 {
				t.Errorf(`Expected width and height to be 3, got %d and %d`, w, h)
			}
			for idx, val := range expected {
				if result[idx] != val {
					t.Errorf(`Expected index %d (%d,%d) to be %08b, got %08b`, idx, idx%w, idx/w, val, result[idx])
				}
			}
		})
		t.Run("test sample 2", func(t *testing.T) {
			reader := strings.NewReader(testData)
			result, w, h, err := parseInput(reader, false)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if w != 23 || h != 23 {
				t.Errorf(`Expected width and height to be 23, got %d and %d`, w, h)
			}
			testPoints := map[int]cell{
				1:         dirExitBottom | dirEntryBottom,
				24:        dirExitRight | dirExitTop | dirEntryRight | dirEntryTop,
				25:        dirExitLeft | dirExitRight | dirEntryLeft | dirEntryRight,
				3*23 + 9:  dirExitRight | dirExitBottom | dirEntryBottom | dirEntryRight,
				3*23 + 10: dirExitRight | dirExitLeft | dirEntryLeft | dirEntryRight,
				3*23 + 11: dirExitRight | dirExitBottom | dirExitLeft | dirEntryLeft | dirEntryRight | dirEntryBottom,
			}
			for idx, val := range testPoints {
				if result[idx] != val {
					t.Errorf(`Expected index %d (%d,%d) to be %08b, got %08b`, idx, idx%w, idx/w, val, result[idx])
				}
			}
		})
	})
}

func Benchmark_parseInput(b *testing.B) {
	testData := []string{
		testData,
	}

	full_data, ok := test_utils.GetTestData(23, 23)
	if ok {
		testData = append(testData, full_data)
	}
	b.Run("respect slope", func(b *testing.B) {
		for _, data := range testData {
			b.Run(fmt.Sprintf("data size %d", len(data)), func(b *testing.B) {
				reader := strings.NewReader(data)
				for b.Loop() {
					reader.Seek(0, 0)
					parseInput(reader, true)
				}
			})
		}
	})

	b.Run("don't respect slope", func(b *testing.B) {
		for _, data := range testData {
			b.Run(fmt.Sprintf("data size %d", len(data)), func(b *testing.B) {
				reader := strings.NewReader(data)
				for b.Loop() {
					reader.Seek(0, 0)
					parseInput(reader, false)
				}
			})
		}
	})
}
