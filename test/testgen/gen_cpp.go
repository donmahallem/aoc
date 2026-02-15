package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

func init() {
	RegisterGenerator("cpp", &CppGenerator{})
}

type CppGenerator struct{}

// CppTestFileData is the template context for C++ test file generation.
type CppTestFileData struct {
	BaseTemplateData
	Namespace  string // "aoc24::day01" – used in code (function calls)
	TestSuite  string // "Aoc24Day01" – used in TEST() macro (no :: allowed)
	HeaderPath string // "aoc24/day01/day01.h"
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
			parts[i] = g.FormatExpected(elem, typeHint)
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

// GetTemplateData returns the C++ template data.
func (g CppGenerator) GetTemplateData(dd DayTestData, yearPkg string) interface{} {
	dayPkg := fmt.Sprintf("day%s", dd.PaddedDay)
	namespace := fmt.Sprintf("%s::%s", yearPkg, dayPkg)
	testSuite := fmt.Sprintf("Aoc%02dDay%s", 2000+dd.YearInt-2000, dd.PaddedDay)
	headerPath := fmt.Sprintf("%s/%s/%s.h", yearPkg, dayPkg, dayPkg)

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
		TestSuite:        testSuite,
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

// FormatContent returns the content as-is; the template already produces consistent output.
func (g CppGenerator) FormatContent(content, targetPath string) (string, error) {
	return content, nil
}

// cppEscapeString escapes a string for use as a C++ raw string literal content.
// Since we use R"(...)" delimiters, the only thing that would break is the
// sequence )". We use a custom delimiter R"testgen(...)testgen" to avoid this.
func cppEscapeString(s string) string {
	return s
}

// Compile-time check
var _ Generator = (*CppGenerator)(nil)
