package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

func init() {
	RegisterGenerator("cpp", &CppGenerator{})
}

type CppGenerator struct{}

// CppTestFileData is the template context for C++ test file generation.
type CppTestFileData struct {
	BaseTemplateData
	Namespace  string // "Aoc24Day01"
	HeaderPath string // "../../../src/aoc24/day01/day01.h"
	HasFiles   bool   // whether any file-based test cases exist
}

// TemplateName returns the template filename for C++ tests.
func (g CppGenerator) TemplateName() string { return "generated_test.cpp.tmpl" }

// OutputFilename returns the output test filename for C++.
func (g CppGenerator) OutputFilename() string { return "generated_test.cpp" }

// FuncMap provides C++-specific template helpers.
func (g CppGenerator) FuncMap() template.FuncMap {
	return template.FuncMap{
		"cppEscapeString": cppEscapeString,
	}
}

// FormatExpected formats expected values for C++ literals.
func (g CppGenerator) FormatExpected(v interface{}, typeHint *string) string {
	return FormatExpectedCpp(v, typeHint)
}

// GetTemplateData returns the C++ template data.
func (g CppGenerator) GetTemplateData(dd DayTestData, yearPkg string) interface{} {
	dayPkg := fmt.Sprintf("day%s", dd.PaddedDay)
	namespace := fmt.Sprintf("Aoc%02dDay%s", 2000+dd.YearInt-2000, dd.PaddedDay)
	headerPath := fmt.Sprintf("../../../src/%s/%s/%s.h", yearPkg, dayPkg, dayPkg)

	hasFiles := false
	if dd.Part1 != nil && len(dd.Part1.Files) > 0 {
		hasFiles = true
	}
	if dd.Part2 != nil && len(dd.Part2.Files) > 0 {
		hasFiles = true
	}

	return CppTestFileData{
		BaseTemplateData: BuildBaseTemplateData(dd, yearPkg),
		Namespace:        namespace,
		HeaderPath:       headerPath,
		HasFiles:         hasFiles,
	}
}

// GetOutputPath returns the C++ output directory for a day.
func (g CppGenerator) GetOutputPath(baseDir, yearPkg string, dd DayTestData) string {
	dayPkg := fmt.Sprintf("day%s", dd.PaddedDay)
	return filepath.Join(baseDir, yearPkg, dayPkg)
}

// PrepareOutput uses default no-op.
func (g CppGenerator) PrepareOutput(outPath string, dd DayTestData, opts GenerationOptions) error {
	return nil
}

// ShouldGenerate returns true only when the day's header file exists in source_dir.
// This prevents generating tests for days that have no C++ implementation yet.
func (g CppGenerator) ShouldGenerate(dd DayTestData, yearPkg, sourceDir string) bool {
	if sourceDir == "" {
		return true
	}
	dayPkg := fmt.Sprintf("day%s", dd.PaddedDay)
	header := filepath.Join(sourceDir, yearPkg, dayPkg, dayPkg+".h")
	_, err := os.Stat(header)
	return err == nil
}

// FormatContent runs clang-format on the generated C++ content for consistent style.
// Falls back to unformatted content if clang-format is not available.
func (g CppGenerator) FormatContent(content, targetPath string) (string, error) {
	cmd := exec.Command("clang-format", "--assume-filename="+targetPath)
	cmd.Stdin = bytes.NewReader([]byte(content))
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		// clang-format not available; return unformatted content
		return content, nil
	}
	return stdout.String(), nil
}

// cppEscapeString escapes a string for use as a C++ raw string literal content.
// Since we use R"(...)" delimiters, the only thing that would break is the
// sequence )". We use a custom delimiter R"testgen(...)testgen" to avoid this.
func cppEscapeString(s string) string {
	return s
}

// Compile-time check
var _ Generator = (*CppGenerator)(nil)
