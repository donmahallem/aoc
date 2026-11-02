package generator

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/donmahallem/aoc/test/testgen/model"
)

// DayInfo holds the template data shared by all languages for one day.
type DayInfo struct {
	// Puzzle year, e.g. 2024.
	YearNumber int
	// Puzzle day, 1 to 25.
	DayNumber int
	// Day zero-padded to two digits, e.g. "01".
	PaddedDay string
	// Year directory name, e.g. "aoc24".
	YearDir string
	// Cases of part 1, nil when none apply.
	Part1 *Part
	// Cases of part 2, nil when none apply.
	Part2 *Part
	// Inline inputs referenced by the cases.
	InlineSamples []Sample
	// True when any case uses an inline input.
	HasInlineCases bool
	// True when any case reads its input from a file.
	HasFileCases bool
	// True when any expected result is an array.
	HasArrayExpectations bool
}

// Part holds the test cases of one part. Sample cases come first, then full data cases,
// each group sorted by name.
type Part struct {
	Cases []Case
}

// Case is a single expectation for one part.
type Case struct {
	Name      string
	VarName   string // inline input variable; empty when the input is read from FilePath
	FilePath  string // input file; empty when the input is inline
	Expected  string
	IsArray   bool
	ValueKind string // C result kind: i64, cstr, i64_array, i16_array or cstr_array
}

// Sample is an inline input shared by the parts that reference it.
type Sample struct {
	VarName string
	Input   string
}

// newDay collects the cases not skipped for lang, sample cases first, then full data cases.
func newDay(year, day int, lang string, cases model.TestCases, expected func(model.Result, model.ResultType) string) *DayInfo {
	d := &DayInfo{
		YearNumber: year,
		DayNumber:  day,
		PaddedDay:  fmt.Sprintf("%02d", day),
		YearDir:    fmt.Sprintf("aoc%d", year),
	}

	var parts [2]Part
	for _, name := range caseOrder(cases) {
		tc := cases[name]
		if slices.Contains(tc.SkipLanguages, lang) {
			continue
		}
		sample := Sample{VarName: sampleVarName(len(d.InlineSamples))}
		used := false
		for i, pe := range []*model.PartExpectation{tc.Part1, tc.Part2} {
			if pe == nil || slices.Contains(pe.SkipLanguages, lang) {
				continue
			}
			c := Case{
				Name:      name,
				Expected:  expected(pe.Result, pe.Type),
				IsArray:   pe.Type.IsArray(),
				ValueKind: valueKind(pe.Type),
			}
			d.HasArrayExpectations = d.HasArrayExpectations || c.IsArray
			if tc.Input != nil {
				c.VarName = sample.VarName
				d.HasInlineCases = true
			} else {
				c.FilePath = *tc.File
				d.HasFileCases = true
			}
			parts[i].Cases = append(parts[i].Cases, c)
			used = true
		}
		if used && tc.Input != nil {
			sample.Input = *tc.Input
			d.InlineSamples = append(d.InlineSamples, sample)
		}
	}

	if len(parts[0].Cases) > 0 {
		d.Part1 = &parts[0]
	}
	if len(parts[1].Cases) > 0 {
		d.Part2 = &parts[1]
	}
	return d
}

// caseOrder returns the case names with sample cases first, then full data cases.
// Each group is sorted by name.
func caseOrder(cases model.TestCases) []string {
	names := slices.Sorted(maps.Keys(cases))
	slices.SortStableFunc(names, func(a, b string) int {
		sa, sb := cases[a].IsSample(), cases[b].IsSample()
		switch {
		case sa && !sb:
			return -1
		case !sa && sb:
			return 1
		}
		return strings.Compare(a, b)
	})
	return names
}

func dayPkg(d *DayInfo) string {
	return "day" + d.PaddedDay
}

func headerPath(d *DayInfo) string {
	return fmt.Sprintf("%s/%s/%s.h", d.YearDir, dayPkg(d), dayPkg(d))
}

func sampleVarName(index int) string {
	if index == 0 {
		return "testData"
	}
	return fmt.Sprintf("testData%d", index+1)
}

func valueKind(t model.ResultType) string {
	kind := "i64"
	switch t.Elem() {
	case model.ResultTypeString:
		kind = "cstr"
	case model.ResultTypeInt16:
		if t.IsArray() {
			kind = "i16"
		}
	}
	if t.IsArray() {
		kind += "_array"
	}
	return kind
}
