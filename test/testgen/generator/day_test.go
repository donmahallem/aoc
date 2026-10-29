package generator

import (
	"slices"
	"testing"

	"github.com/donmahallem/aoc/test/testgen/model"
)

// helper to create a pointer to a v
func partExp(v int64) *model.PartExpectation {
	return &model.PartExpectation{Result: model.Result{Kind: model.ResultInt, Int: v}, Type: model.ResultTypeInt}
}

func caseNames(cases []Case) []string {
	names := make([]string, len(cases))
	for i, c := range cases {
		names[i] = c.Name
	}
	return names
}

func TestCaseOrder(t *testing.T) {
	cases := model.TestCases{
		"z_sample": {Input: new("z")},
		"a_full":   {FullData: true, File: new("a.txt")},
		"m_sample": {Input: new("m")},
		"b_full":   {FullData: true, File: new("b.txt")},
	}
	want := []string{"m_sample", "z_sample", "a_full", "b_full"}
	if got := caseOrder(cases); !slices.Equal(got, want) {
		t.Errorf("caseOrder() = %v, want %v", got, want)
	}
}

func TestNewDay(t *testing.T) {
	cases := model.TestCases{
		"a_sample": {Input: new("x"), Part1: partExp(1), Part2: partExp(3)},
		"b_full":   {FullData: true, File: new("data/full/24/01.txt"), Part1: partExp(2)},
		"skipped":  {Input: new("y"), SkipLanguages: []string{"go"}, Part1: partExp(9)},
	}
	d := newDay(2024, 1, "go", cases, goExpected)

	if d.PaddedDay != "01" || d.YearDir != "aoc2024" {
		t.Errorf("PaddedDay = %q, YearDir = %q", d.PaddedDay, d.YearDir)
	}
	if d.Part1 == nil || !slices.Equal(caseNames(d.Part1.Cases), []string{"a_sample", "b_full"}) {
		t.Fatalf("unexpected part1 cases: %+v", d.Part1)
	}
	if d.Part2 == nil || !slices.Equal(caseNames(d.Part2.Cases), []string{"a_sample"}) {
		t.Fatalf("unexpected part2 cases: %+v", d.Part2)
	}
	if got := d.Part1.Cases[0]; got.VarName != "testData" || got.Expected != "1" {
		t.Errorf("inline case = %+v", got)
	}
	if got := d.Part1.Cases[1]; got.FilePath != "data/full/24/01.txt" || got.Expected != "2" {
		t.Errorf("file case = %+v", got)
	}
	if len(d.InlineSamples) != 1 || d.InlineSamples[0].Input != "x" {
		t.Errorf("InlineSamples = %+v", d.InlineSamples)
	}
	if !d.HasInlineCases || !d.HasFileCases || d.HasArrayExpectations {
		t.Errorf("HasInlineCases = %v, HasFileCases = %v, HasArrayExpectations = %v",
			d.HasInlineCases, d.HasFileCases, d.HasArrayExpectations)
	}
}

func TestValueKind(t *testing.T) {
	tests := map[model.ResultType]string{
		model.ResultTypeInt:     "i64",
		model.ResultTypeInt16:   "i64",
		model.ResultTypeString:  "cstr",
		model.ResultTypeInt16s:  "i16_array",
		model.ResultTypeInts:    "i64_array",
		model.ResultTypeStrings: "cstr_array",
	}
	for typ, want := range tests {
		if got := valueKind(typ); got != want {
			t.Errorf("valueKind(%q) = %q, want %q", typ, got, want)
		}
	}
}

func TestSampleVarName(t *testing.T) {
	tests := []struct {
		index int
		want  string
	}{
		{0, "testData"},
		{1, "testData2"},
		{2, "testData3"},
	}
	for _, tt := range tests {
		if got := sampleVarName(tt.index); got != tt.want {
			t.Errorf("sampleVarName(%d) = %q, want %q", tt.index, got, tt.want)
		}
	}
}
