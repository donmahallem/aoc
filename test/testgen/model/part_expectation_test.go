package model_test

import (
	"encoding/json/v2"
	"testing"

	model "github.com/donmahallem/aoc/test/testgen/model"
)

const sample1 string = "{\"result\": 6284877, \"type\": \"int\"}"
const sample2 string = "{\"result\": [\"x\",\"y\"], \"type\": \"[]string\"}"
const sample3 string = "{\"result\": [1,5], \"type\": \"[]int16\"}"

func TestPartExpectation(t *testing.T) {
	t.Run("sample1", func(t *testing.T) {
		var pe model.PartExpectation
		if err := json.Unmarshal([]byte(sample1), &pe); err != nil {
			t.Fatalf("failed to unmarshal sample1: %v", err)
		}
		if pe.Result.Kind != model.ResultInt {
			t.Errorf("expected ResultInt, got %v", pe.Result.Kind)
		}
		if pe.Result.Int != 6284877 {
			t.Errorf("expected 6284877, got %v", pe.Result.Int)
		}
	})

	t.Run("test symmetry", func(t *testing.T) {
		var pe model.PartExpectation = model.PartExpectation{
			Result: model.Result{
				Kind: model.ResultInt,
				Int:  6284877,
			},
			Type: model.ResultTypeInt,
		}
		data, err := json.Marshal(pe)
		if err != nil {
			t.Fatalf("failed to marshal sample1: %v", err)
		}
		var pe2 model.PartExpectation
		if err := json.Unmarshal(data, &pe2); err != nil {
			t.Fatalf("failed to unmarshal marshaled data: %v", err)
		}
		if pe2.Result.Kind != pe.Result.Kind {
			t.Errorf("expected ResultInt, got %v", pe2.Result.Kind)
		}
		if pe2.Result.Int != pe.Result.Int {
			t.Errorf("expected %v, got %v", pe.Result.Int, pe2.Result.Int)
		}
		if pe2.Type != pe.Type {
			t.Errorf("expected type %v, got %v", pe.Type, pe2.Type)
		}
	})

	t.Run("sample2", func(t *testing.T) {
		var pe model.PartExpectation
		if err := json.Unmarshal([]byte(sample2), &pe); err != nil {
			t.Fatalf("failed to unmarshal sample2: %v", err)
		}
		if pe.Type != model.ResultTypeStrings {
			t.Errorf("expected ResultTypeStrings, got %v", pe.Type)
		}
		if len(pe.Result.Strs) != 2 || pe.Result.Strs[0] != "x" || pe.Result.Strs[1] != "y" {
			t.Errorf("expected [\"x\",\"y\"], got %v", pe.Result.Strs)
		}
	})

	t.Run("sample3", func(t *testing.T) {
		var pe model.PartExpectation
		if err := json.Unmarshal([]byte(sample3), &pe); err != nil {
			t.Fatalf("failed to unmarshal sample3: %v", err)
		}
		if pe.Type != model.ResultTypeInt16s {
			t.Errorf("expected ResultTypeInt16s, got %v", pe.Type)
		}
		if len(pe.Result.Ints) != 2 || pe.Result.Ints[0] != 1 || pe.Result.Ints[1] != 5 {
			t.Errorf("expected [1,5], got %v", pe.Result.Ints)
		}
	})
}
