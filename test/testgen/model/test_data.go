package model

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// TestData maps year -> day -> named test cases, as described by test_schema.json.
// JSON keys are zero-padded numbers ("24", "01").
type TestData map[int]map[int]TestCases

type rawTestData = map[string]map[string]TestCases

// Load reads and parses a data file.
func Load(path string) (TestData, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read data file: %w", err)
	}
	var data TestData
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, fmt.Errorf("failed to parse data file: %w", err)
	}
	base, err := filepath.Abs(filepath.Dir(path))
	if err != nil {
		return nil, fmt.Errorf("failed to resolve data file directory: %w", err)
	}
	for _, days := range data {
		for _, cases := range days {
			for name, tc := range cases {
				tc.baseDir = base
				cases[name] = tc
			}
		}
	}
	return data, nil
}

func (r *TestData) UnmarshalJSONFrom(dec *jsontext.Decoder) error {
	var raw rawTestData
	if err := json.UnmarshalDecode(dec, &raw); err != nil {
		return err
	}
	data := make(TestData, len(raw))
	for yearKey, days := range raw {
		year, err := strconv.Atoi(yearKey)
		if err != nil {
			return fmt.Errorf("invalid year key %q: %w", yearKey, err)
		}
		data[year] = make(map[int]TestCases, len(days))
		for dayKey, cases := range days {
			day, err := strconv.Atoi(dayKey)
			if err != nil {
				return fmt.Errorf("invalid day key %q: %w", dayKey, err)
			}
			data[year][day] = cases
		}
	}
	*r = data
	return nil
}

func (r TestData) MarshalJSONTo(enc *jsontext.Encoder) error {
	if r == nil {
		return enc.WriteToken(jsontext.Null)
	}
	raw := make(rawTestData, len(r))
	for year, days := range r {
		rawDays := make(map[string]TestCases, len(days))
		for day, cases := range days {
			rawDays[fmt.Sprintf("%02d", day)] = cases
		}
		raw[fmt.Sprintf("%02d", year)] = rawDays
	}
	return json.MarshalEncode(enc, raw, json.Deterministic(true))
}
