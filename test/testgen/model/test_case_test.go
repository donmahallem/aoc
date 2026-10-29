package model_test

import (
	"encoding/json/v2"
	"testing"

	model "github.com/donmahallem/aoc/test/testgen/model"
)

const sampleTestCase string = "{\"input\": \"test input\", \"part1\": {\"result\": 123, \"type\": \"int16\"}, \"part2\": {\"result\": [\"a\",\"b\"], \"type\": \"[]string\"}}"

func TestTestCase(t *testing.T) {
	t.Run("sampleTestCase", func(t *testing.T) {
		var tc model.TestCase
		if err := json.Unmarshal([]byte(sampleTestCase), &tc); err != nil {
			t.Fatalf("failed to unmarshal sampleTestCase: %v", err)
		}
		if tc.Input == nil || *tc.Input != "test input" {
			t.Errorf("expected input 'test input', got %v", tc.Input)
		}
		if tc.Part1.Result.Int != 123 {
			t.Errorf("expected part1 result 123, got %v", tc.Part1.Result.Int)
		}
		if len(tc.Part2.Result.Strs) != 2 || tc.Part2.Result.Strs[0] != "a" || tc.Part2.Result.Strs[1] != "b" {
			t.Errorf("expected part2 result [\"a\",\"b\"], got %v", tc.Part2.Result.Strs)
		}
	})

	t.Run("test symmetry", func(t *testing.T) {
		var tc model.TestCase
		if err := json.Unmarshal([]byte(sampleTestCase), &tc); err != nil {
			t.Fatalf("failed to unmarshal sampleTestCase: %v", err)
		}
		data, err := json.Marshal(tc)
		if err != nil {
			t.Fatalf("failed to marshal sampleTestCase: %v", err)
		}
		var tc2 model.TestCase
		if err := json.Unmarshal(data, &tc2); err != nil {
			t.Fatalf("failed to unmarshal marshaled data: %v", err)
		}
		if tc2.Input == nil || *tc2.Input != "test input" {
			t.Errorf("expected input 'test input', got %v", tc2.Input)
		}
		if tc2.Part1.Result.Int != 123 {
			t.Errorf("expected part1 result 123, got %v", tc2.Part1.Result.Int)
		}
		if len(tc2.Part2.Result.Strs) != 2 || tc2.Part2.Result.Strs[0] != "a" || tc2.Part2.Result.Strs[1] != "b" {
			t.Errorf("expected part2 result [\"a\",\"b\"], got %v", tc2.Part2.Result.Strs)
		}
	})
}
