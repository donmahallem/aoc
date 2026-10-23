package test_utils

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// ResultKind identifies which field of a Result holds the value.
type ResultKind int

const (
	ResultUnset ResultKind = iota
	ResultInt
	ResultInts
	ResultString
	ResultStrings
)

// Result is an expected answer: an int, []int, string, or []string.
type Result struct {
	Kind ResultKind
	Int  int64
	Ints []int64
	Str  string
	Strs []string
}

// implements json.UnmarshalerFrom.
func (r *Result) UnmarshalJSONFrom(dec *jsontext.Decoder) error {
	switch dec.PeekKind() {
	case '0':
		r.Kind = ResultInt
		return json.UnmarshalDecode(dec, &r.Int)
	case '"':
		r.Kind = ResultString
		return json.UnmarshalDecode(dec, &r.Str)
	case '[':
		raw, err := dec.ReadValue()
		if err != nil {
			return err
		}
		var ints []int64
		if json.Unmarshal(raw, &ints) == nil {
			r.Kind, r.Ints = ResultInts, ints
			return nil
		}
		var strs []string
		if json.Unmarshal(raw, &strs) == nil {
			r.Kind, r.Strs = ResultStrings, strs
			return nil
		}
		return errors.New("result array must contain only integers or only strings")
	default:
		return fmt.Errorf("result must be an integer, string, or array; got %v", dec.PeekKind())
	}
}

func (r *Result) MarshalJSONTo(dec *jsontext.Encoder) error {
	switch r.Kind {
	case ResultInt:
		return json.MarshalEncode(dec, r.Int)
	case ResultInts:
		return json.MarshalEncode(dec, r.Ints)
	case ResultString:
		return json.MarshalEncode(dec, r.Str)
	case ResultStrings:
		return json.MarshalEncode(dec, r.Strs)
	default:
		return json.MarshalEncode(dec, nil)
	}
}

type PartExpectation struct {
	Result        Result   `json:"result"`
	Type          string   `json:"type"`
	SkipLanguages []string `json:"skip_languages"`
}

type TestCase struct {
	Name          string           `json:"name"`
	Input         *string          `json:"input"`
	File          *string          `json:"file"`
	Part1         *PartExpectation `json:"part1"`
	Part2         *PartExpectation `json:"part2"`
	SkipLanguages []string         `json:"skip_languages"`
}

// TestData year, day, to the list of test cases.
type TestData map[string]map[string][]TestCase

func DefaultTestDataPath() string {
	return filepath.Join(basepath, "../../test/data.json")
}

func LoadTestData(path string) (TestData, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return parseTestData(raw, path)
}

func parseTestData(raw []byte, name string) (TestData, error) {
	var data TestData
	if err := json.Unmarshal(raw, &data, json.RejectUnknownMembers(true)); err != nil {
		return nil, fmt.Errorf("parse %s: %w", name, err)
	}
	if err := validateTestData(data); err != nil {
		return nil, err
	}
	return data, nil
}

func validateTestData(data TestData) error {
	for year, days := range data {
		for day, list := range days {
			for i, tc := range list {
				where := fmt.Sprintf("%s/%s[%d] (%s)", year, day, i, tc.Name)
				if (tc.Input == nil) == (tc.File == nil) {
					return fmt.Errorf("%s: exactly one of input or file is required", where)
				}
				if tc.Part1 == nil && tc.Part2 == nil {
					return fmt.Errorf("%s: part1 or part2 is required", where)
				}
				for _, p := range []*PartExpectation{tc.Part1, tc.Part2} {
					if p != nil && p.Result.Kind == ResultUnset {
						return fmt.Errorf("%s: result is required", where)
					}
				}
			}
		}
	}
	return nil
}
