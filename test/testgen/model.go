package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// TestCase matches the JSON schema: each entry has input OR file, and part1 and/or part2.
type TestCase struct {
	Input *string `json:"input,omitempty"`
	File  *string `json:"file,omitempty"`
	Name  string  `json:"name"`
	Part1 any     `json:"part1,omitempty"`
	Part2 any     `json:"part2,omitempty"`
}

func (tc TestCase) IsInline() bool {
	return tc.Input != nil
}

func (tc TestCase) IsFile() bool {
	return tc.File != nil
}

// SampleEntry represents an inline sample test case for a specific part.
type SampleEntry struct {
	Index    int
	Name     string
	Input    string
	Expected string
	Part     int
	VarName  string // Variable name to reference in generated code (e.g. testData, testData2)
	IsArray  bool   // Whether the expected value is an array (e.g. []int)
}

// InlineSample represents a unique inline sample input with its variable name.
type InlineSample struct {
	VarName string
	Input   string
}

// SampleVarName returns the variable name for an inline sample at the given index.
func SampleVarName(index int) string {
	if index == 0 {
		return "testData"
	}
	return fmt.Sprintf("testData%d", index+1)
}

// FileTestEntry represents a file-based test case for a specific part.
type FileTestEntry struct {
	Name     string
	FilePath string
	Expected string
	Part     int
	IsArray  bool // Whether the expected value is an array (e.g. []int)
}

// PartData holds the parsed test data for a single part of a day.
type PartData struct {
	Samples []SampleEntry
	Files   []FileTestEntry
}

// DayTestData holds all parsed test data for a single day.
type DayTestData struct {
	YearStr        string
	YearInt        int
	DayStr         string
	DayInt         int
	PaddedDay      string
	Part1          *PartData
	Part2          *PartData
	HasSamples     bool
	InlineSamples  []InlineSample
	HasArrayResult bool // Whether any expected value is an array
}

// BaseTemplateData is the shared portion of template context across languages.
// Language-specific generators can embed this and add their own fields.
type BaseTemplateData struct {
	Year           int
	Day            int
	PaddedDay      string
	YearPkg        string
	Part1          *PartData
	Part2          *PartData
	HasSamples     bool
	InlineSamples  []InlineSample
	HasArrayResult bool
}

type TestData map[string]map[string][]TestCase

// isIntArray checks if an any value is a JSON array of numbers.
func isIntArray(v any) bool {
	arr, ok := v.([]any)
	if !ok || len(arr) == 0 {
		return false
	}
	for _, elem := range arr {
		if _, ok := elem.(float64); !ok {
			return false
		}
	}
	return true
}

// FormatExpected returns a language-agnostic string representation of an expected value.
// Language-specific formatting (e.g. Go vs Python literals) is handled in the generators.
func FormatExpected(v any) string {
	switch val := v.(type) {
	case float64:
		if val == float64(int64(val)) {
			return strconv.FormatInt(int64(val), 10)
		}
		return strconv.FormatFloat(val, 'f', -1, 64)
	case string:
		return val
	case []any:
		parts := make([]string, len(val))
		for i, elem := range val {
			parts[i] = FormatExpected(elem)
		}
		return "[" + strings.Join(parts, ", ") + "]"
	default:
		return fmt.Sprintf("%v", val)
	}
}

// FormatExpectedGo returns a Go-literal representation (strings are quoted).
func FormatExpectedGo(v any) string {
	switch val := v.(type) {
	case float64:
		if val == float64(int64(val)) {
			return strconv.FormatInt(int64(val), 10)
		}
		return strconv.FormatFloat(val, 'f', -1, 64)
	case string:
		return fmt.Sprintf("%q", val)
	case []any:
		parts := make([]string, len(val))
		for i, elem := range val {
			parts[i] = FormatExpectedGo(elem)
		}
		return "[]int{" + strings.Join(parts, ", ") + "}"
	default:
		return fmt.Sprintf("%v", val)
	}
}

// FormatExpectedPython returns a Python-literal representation.
func FormatExpectedPython(v any) string {
	switch val := v.(type) {
	case float64:
		if val == float64(int64(val)) {
			return strconv.FormatInt(int64(val), 10)
		}
		return strconv.FormatFloat(val, 'f', -1, 64)
	case string:
		return fmt.Sprintf("%q", val)
	case []any:
		parts := make([]string, len(val))
		for i, elem := range val {
			parts[i] = FormatExpectedPython(elem)
		}
		return "[" + strings.Join(parts, ", ") + "]"
	default:
		return fmt.Sprintf("%v", val)
	}
}

// LoadData reads and parses the data.json file.
func LoadData(path string) (*TestData, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read data file: %w", err)
	}

	var data TestData
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, fmt.Errorf("failed to parse data file: %w", err)
	}
	return &data, nil
}

// ValidateData performs structural validation on the loaded test data before generation.
func ValidateData(data TestData) error {
	for year, days := range data {
		for day, tests := range days {
			for i, tc := range tests {
				hasInput := tc.Input != nil && *tc.Input != ""
				hasFile := tc.File != nil && *tc.File != ""
				if hasInput == hasFile {
					return fmt.Errorf("%s/%s test[%d] '%s': specify exactly one of input or file", year, day, i, tc.Name)
				}

				hasResult := tc.Part1 != nil || tc.Part2 != nil
				if !hasResult {
					return fmt.Errorf("%s/%s test[%d] '%s': requires part1 and/or part2 expected values", year, day, i, tc.Name)
				}
			}
		}
	}
	return nil
}

// ParseDayData converts raw test cases into a structured DayTestData.
func ParseDayData(yearStr, dayStr string, testCases []TestCase, fmtExpected func(any) string) DayTestData {
	yearInt, _ := strconv.Atoi(yearStr)
	dayInt, _ := strconv.Atoi(dayStr)
	paddedDay := fmt.Sprintf("%02d", dayInt)

	var p1, p2 PartData
	hasPart1, hasPart2 := false, false
	hasArrayResult := false
	sampleIdx := 0
	var inlineSamples []InlineSample

	for _, testCase := range testCases {
		varName := ""
		if testCase.IsInline() {
			varName = SampleVarName(sampleIdx)
			inlineSamples = append(inlineSamples, InlineSample{
				VarName: varName,
				Input:   *testCase.Input,
			})
		}
		if testCase.Part1 != nil {
			hasPart1 = true
			p1IsArray := isIntArray(testCase.Part1)
			if p1IsArray {
				hasArrayResult = true
			}
			if testCase.IsInline() {
				p1.Samples = append(p1.Samples, SampleEntry{
					Index:    sampleIdx,
					Name:     testCase.Name,
					Input:    *testCase.Input,
					Expected: fmtExpected(testCase.Part1),
					Part:     1,
					VarName:  varName,
					IsArray:  p1IsArray,
				})
			} else if testCase.IsFile() {
				p1.Files = append(p1.Files, FileTestEntry{
					Name:     testCase.Name,
					FilePath: *testCase.File,
					Expected: fmtExpected(testCase.Part1),
					Part:     1,
					IsArray:  p1IsArray,
				})
			}
		}
		if testCase.Part2 != nil {
			hasPart2 = true
			p2IsArray := isIntArray(testCase.Part2)
			if p2IsArray {
				hasArrayResult = true
			}
			if testCase.IsInline() {
				p2.Samples = append(p2.Samples, SampleEntry{
					Index:    sampleIdx,
					Name:     testCase.Name,
					Input:    *testCase.Input,
					Expected: fmtExpected(testCase.Part2),
					Part:     2,
					VarName:  varName,
					IsArray:  p2IsArray,
				})
			} else if testCase.IsFile() {
				p2.Files = append(p2.Files, FileTestEntry{
					Name:     testCase.Name,
					FilePath: *testCase.File,
					Expected: fmtExpected(testCase.Part2),
					Part:     2,
					IsArray:  p2IsArray,
				})
			}
		}
		if testCase.IsInline() {
			sampleIdx++
		}
	}

	hasSamples := len(p1.Samples) > 0 || len(p2.Samples) > 0

	dd := DayTestData{
		YearStr:        yearStr,
		YearInt:        yearInt,
		DayStr:         dayStr,
		DayInt:         dayInt,
		PaddedDay:      paddedDay,
		HasSamples:     hasSamples,
		InlineSamples:  inlineSamples,
		HasArrayResult: hasArrayResult,
	}
	if hasPart1 {
		dd.Part1 = &p1
	}
	if hasPart2 {
		dd.Part2 = &p2
	}
	return dd
}

// BuildBaseTemplateData constructs the reusable template context portion.
func BuildBaseTemplateData(dd DayTestData, yearPkg string) BaseTemplateData {
	return BaseTemplateData{
		Year:           dd.YearInt,
		Day:            dd.DayInt,
		PaddedDay:      dd.PaddedDay,
		YearPkg:        yearPkg,
		Part1:          dd.Part1,
		Part2:          dd.Part2,
		HasSamples:     dd.HasSamples,
		InlineSamples:  dd.InlineSamples,
		HasArrayResult: dd.HasArrayResult,
	}
}

func SortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
