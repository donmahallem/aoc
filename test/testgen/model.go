package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// PartSpec groups result, type hint, and per-language skips for a part.
type PartSpec struct {
	Result any      `json:"result"`
	Type   *string  `json:"type,omitempty"`
	Skip   []string `json:"skip_languages,omitempty"`
}

// TestCase matches the JSON schema: each entry has input OR file, and part1 and/or part2.
type TestCase struct {
	Input *string `json:"input,omitempty"`
	File  *string `json:"file,omitempty"`
	Name  string  `json:"name"`

	Part1 *PartSpec `json:"part1,omitempty"`
	Part2 *PartSpec `json:"part2,omitempty"`

	SkipLangs []string `json:"skip_languages,omitempty"`
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

type arrayKind int

const (
	arrayKindNone arrayKind = iota
	arrayKindInt
	arrayKindInt16
	arrayKindString
	arrayKindInvalid
)

func classifyArray(v any) arrayKind {
	arr, ok := v.([]any)
	if !ok {
		return arrayKindNone
	}
	if len(arr) == 0 {
		return arrayKindInvalid
	}
	kind := arrayKindNone
	for _, elem := range arr {
		switch elem.(type) {
		case float64:
			if kind == arrayKindString {
				return arrayKindInvalid
			}
			kind = arrayKindInt
		case string:
			if kind == arrayKindInt {
				return arrayKindInvalid
			}
			kind = arrayKindString
		default:
			return arrayKindInvalid
		}
	}
	return kind
}

// resolveArrayKind returns the array kind, applying an optional type hint.
// Hints may widen numeric arrays to int16 or enforce string arrays; mismatches are invalid.
func resolveArrayKind(v any, typeHint *string) arrayKind {
	base := classifyArray(v)
	if typeHint == nil {
		return base
	}
	switch *typeHint {
	case "int":
		if base == arrayKindNone {
			return arrayKindInvalid
		}
		if base == arrayKindString {
			return arrayKindInvalid
		}
		return arrayKindInt
	case "int16":
		if base == arrayKindNone {
			return arrayKindInvalid
		}
		if base == arrayKindString {
			return arrayKindInvalid
		}
		return arrayKindInt16
	case "string":
		if base == arrayKindNone {
			return arrayKindInvalid
		}
		if base != arrayKindString {
			return arrayKindInvalid
		}
		return arrayKindString
	default:
		return arrayKindInvalid
	}
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
// typeHint allows choosing between int and int16 array element types (and scalar casts).
func FormatExpectedGo(v any, typeHint *string) string {
	switch val := v.(type) {
	case float64:
		if val == float64(int64(val)) {
			if typeHint != nil && *typeHint == "int16" {
				return fmt.Sprintf("int16(%d)", int64(val))
			}
			return strconv.FormatInt(int64(val), 10)
		}
		return strconv.FormatFloat(val, 'f', -1, 64)
	case string:
		return fmt.Sprintf("%q", val)
	case []any:
		kind := resolveArrayKind(val, typeHint)
		parts := make([]string, len(val))
		for i, elem := range val {
			parts[i] = FormatExpectedGo(elem, typeHint)
		}
		switch kind {
		case arrayKindInt:
			return "[]int{" + strings.Join(parts, ", ") + "}"
		case arrayKindInt16:
			return "[]int16{" + strings.Join(parts, ", ") + "}"
		case arrayKindString:
			return "[]string{" + strings.Join(parts, ", ") + "}"
		default:
			panic(fmt.Sprintf("unsupported array contents: %#v", val))
		}
	default:
		return fmt.Sprintf("%v", val)
	}
}

// FormatExpectedCpp returns a C++-literal representation.
// typeHint allows choosing between int and int16_t array element types.
func FormatExpectedCpp(v any, typeHint *string) string {
	switch val := v.(type) {
	case float64:
		if val == float64(int64(val)) {
			iv := int64(val)
			if typeHint != nil && *typeHint == "int16" {
				return fmt.Sprintf("static_cast<int16_t>(%d)", iv)
			}
			// Use LL suffix for values that don't fit in 32-bit int
			if iv > 2147483647 || iv < -2147483648 {
				return fmt.Sprintf("%dLL", iv)
			}
			return strconv.FormatInt(iv, 10)
		}
		return strconv.FormatFloat(val, 'f', -1, 64)
	case string:
		return fmt.Sprintf("std::string(%q)", val)
	case []any:
		kind := resolveArrayKind(val, typeHint)
		parts := make([]string, len(val))
		for i, elem := range val {
			parts[i] = FormatExpectedCpp(elem, typeHint)
		}
		switch kind {
		case arrayKindInt:
			return "std::vector<int>{" + strings.Join(parts, ", ") + "}"
		case arrayKindInt16:
			return "std::vector<int16_t>{" + strings.Join(parts, ", ") + "}"
		case arrayKindString:
			return "std::vector<std::string>{" + strings.Join(parts, ", ") + "}"
		default:
			panic(fmt.Sprintf("unsupported array contents: %#v", val))
		}
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

				for partName, part := range map[string]*PartSpec{"part1": tc.Part1, "part2": tc.Part2} {
					if part == nil {
						continue
					}
					if part.Result == nil {
						return fmt.Errorf("%s/%s test[%d] '%s': %s.result is required", year, day, i, tc.Name, partName)
					}
					if part.Type != nil {
						switch *part.Type {
						case "int", "int16", "string":
						default:
							return fmt.Errorf("%s/%s test[%d] '%s': %s.type must be one of int, int16, string", year, day, i, tc.Name, partName)
						}
					}
					if kind := resolveArrayKind(part.Result, part.Type); kind == arrayKindInvalid {
						return fmt.Errorf("%s/%s test[%d] '%s': %s arrays must be non-empty and contain only integers or only strings; optional type hint may be int, int16, or string", year, day, i, tc.Name, partName)
					}
					if err := validateLangList(part.Skip, partName+".skip_languages", year, day, i, tc.Name); err != nil {
						return err
					}
				}

				if err := validateLangList(tc.SkipLangs, "skip_languages", year, day, i, tc.Name); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// ParseDayData converts raw test cases into a structured DayTestData.
func ParseDayData(yearStr, dayStr string, lang string, testCases []TestCase, fmtExpected func(any, *string) string) DayTestData {
	yearInt, _ := strconv.Atoi(yearStr)
	dayInt, _ := strconv.Atoi(dayStr)
	paddedDay := fmt.Sprintf("%02d", dayInt)

	var p1, p2 PartData
	hasPart1, hasPart2 := false, false
	hasArrayResult := false
	sampleIdx := 0
	var inlineSamples []InlineSample

	for _, testCase := range testCases {
		overallSkip := shouldSkip(testCase.SkipLangs, lang)
		p1Allowed := testCase.Part1 != nil && !overallSkip && !shouldSkip(testCase.Part1.Skip, lang)
		p2Allowed := testCase.Part2 != nil && !overallSkip && !shouldSkip(testCase.Part2.Skip, lang)
		if !p1Allowed && !p2Allowed {
			continue
		}
		varName := ""
		if testCase.IsInline() {
			varName = SampleVarName(sampleIdx)
			inlineSamples = append(inlineSamples, InlineSample{
				VarName: varName,
				Input:   *testCase.Input,
			})
		}
		if p1Allowed {
			hasPart1 = true
			p1Kind := resolveArrayKind(testCase.Part1.Result, testCase.Part1.Type)
			p1IsArray := p1Kind == arrayKindInt || p1Kind == arrayKindInt16 || p1Kind == arrayKindString
			if p1IsArray {
				hasArrayResult = true
			}
			if testCase.IsInline() {
				p1.Samples = append(p1.Samples, SampleEntry{
					Index:    sampleIdx,
					Name:     testCase.Name,
					Input:    *testCase.Input,
					Expected: fmtExpected(testCase.Part1.Result, testCase.Part1.Type),
					Part:     1,
					VarName:  varName,
					IsArray:  p1IsArray,
				})
			} else if testCase.IsFile() {
				p1.Files = append(p1.Files, FileTestEntry{
					Name:     testCase.Name,
					FilePath: *testCase.File,
					Expected: fmtExpected(testCase.Part1.Result, testCase.Part1.Type),
					Part:     1,
					IsArray:  p1IsArray,
				})
			}
		}
		if p2Allowed {
			hasPart2 = true
			p2Kind := resolveArrayKind(testCase.Part2.Result, testCase.Part2.Type)
			p2IsArray := p2Kind == arrayKindInt || p2Kind == arrayKindInt16 || p2Kind == arrayKindString
			if p2IsArray {
				hasArrayResult = true
			}
			if testCase.IsInline() {
				p2.Samples = append(p2.Samples, SampleEntry{
					Index:    sampleIdx,
					Name:     testCase.Name,
					Input:    *testCase.Input,
					Expected: fmtExpected(testCase.Part2.Result, testCase.Part2.Type),
					Part:     2,
					VarName:  varName,
					IsArray:  p2IsArray,
				})
			} else if testCase.IsFile() {
				p2.Files = append(p2.Files, FileTestEntry{
					Name:     testCase.Name,
					FilePath: *testCase.File,
					Expected: fmtExpected(testCase.Part2.Result, testCase.Part2.Type),
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

func shouldSkip(skipList []string, lang string) bool {
	for _, s := range skipList {
		if s == lang {
			return true
		}
	}
	return false
}

func validateLangList(list []string, field, year, day string, idx int, name string) error {
	for _, lang := range list {
		if lang == "" {
			return fmt.Errorf("%s/%s test[%d] '%s': %s entries must be non-empty", year, day, idx, name, field)
		}
		switch lang {
		case "go", "python", "cpp":
		default:
			return fmt.Errorf("%s/%s test[%d] '%s': %s entries must be one of go, python, cpp", year, day, idx, name, field)
		}
	}
	return nil
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
