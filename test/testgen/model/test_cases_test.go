package model_test

import (
	"encoding/json/v2"
	"testing"

	model "github.com/donmahallem/aoc/test/testgen/model"
)

const keyedCases string = `{
  "sample1": {
    "input": "1abc2",
    "part1": { "result": 142, "type": "int" }
  },
  "full_data": {
    "file": "data/full/23/01.txt",
    "part1": { "result": [1, 2], "type": "[]int16" },
    "part2": { "result": ["a", "b"], "type": "[]array" }
  }
}`

func TestTestCases(t *testing.T) {
	t.Run("keyed object sets names from keys", func(t *testing.T) {
		var cases model.TestCases
		if err := json.Unmarshal([]byte(keyedCases), &cases); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}
		if len(cases) != 2 {
			t.Fatalf("expected 2 cases, got %d", len(cases))
		}
		if cases["full_data"].Part1.Type != model.ResultTypeInt16s {
			t.Errorf("expected ResultTypeInt16s, got %v", cases["full_data"].Part1.Type)
		}
	})

	t.Run("round trip", func(t *testing.T) {
		var cases model.TestCases
		if err := json.Unmarshal([]byte(keyedCases), &cases); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}
		data, err := json.Marshal(cases)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}
		var again model.TestCases
		if err := json.Unmarshal(data, &again); err != nil {
			t.Fatalf("failed to unmarshal marshaled data: %v", err)
		}
	})

	t.Run("rejects array form", func(t *testing.T) {
		var cases model.TestCases
		if err := json.Unmarshal([]byte(`[{"name": "x"}]`), &cases); err == nil {
			t.Error("expected error for array form")
		}
	})
}

func TestResultTypeNames(t *testing.T) {
	names := []string{
		"int", "int16", "int32", "int64",
		"uint", "uint16", "uint32", "uint64",
		"string",
		"[]int", "[]int16", "[]int32", "[]int64",
		"[]uint", "[]uint16", "[]uint32", "[]uint64",
		"[]string", "[]array",
	}
	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			var pe model.PartExpectation
			input := `{"result": 1, "type": "` + name + `"}`
			if err := json.Unmarshal([]byte(input), &pe); err != nil {
				t.Fatalf("failed to unmarshal type %q: %v", name, err)
			}
			data, err := json.Marshal(pe)
			if err != nil {
				t.Fatalf("failed to marshal type %q: %v", name, err)
			}
			var pe2 model.PartExpectation
			if err := json.Unmarshal(data, &pe2); err != nil || pe2.Type != pe.Type {
				t.Errorf("type %q did not round trip: %s", name, data)
			}
		})
	}

	t.Run("unknown type", func(t *testing.T) {
		var pe model.PartExpectation
		if err := json.Unmarshal([]byte(`{"result": 1, "type": "float"}`), &pe); err == nil {
			t.Error("expected error for unknown type")
		}
	})
}
