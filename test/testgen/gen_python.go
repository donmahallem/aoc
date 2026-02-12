package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

func init() {
	RegisterGenerator("python", &PythonGenerator{})
}

type PythonGenerator struct{}

// PythonTestFileData is the template context for Python test file generation.
// It embeds shared fields and adds Python-specific class names.
type PythonTestFileData struct {
	BaseTemplateData
	Part1Class string // "Test2024Day01Part01"
	Part2Class string // "Test2024Day01Part02"
}

// TemplateName returns the template filename for Python tests.
func (g PythonGenerator) TemplateName() string { return "test_generated.py.tmpl" }

// OutputFilename returns the output test filename for Python.
func (g PythonGenerator) OutputFilename() string { return "test_generated.py" }

// FormatExpected formats expected values for Python literals.
func (g PythonGenerator) FormatExpected(v interface{}) string { return FormatExpectedPython(v) }

// FuncMap returns nil (no custom functions) for Python templates.
func (g PythonGenerator) FuncMap() template.FuncMap { return nil }

// GetTemplateData returns the Python template data.
func (g PythonGenerator) GetTemplateData(dd DayTestData, yearPkg string) interface{} {
	fullYear := 2000 + dd.YearInt
	return PythonTestFileData{
		BaseTemplateData: BuildBaseTemplateData(dd, yearPkg),
		Part1Class:       fmt.Sprintf("Test%dDay%sPart01", fullYear, dd.PaddedDay),
		Part2Class:       fmt.Sprintf("Test%dDay%sPart02", fullYear, dd.PaddedDay),
	}
}

// GetOutputPath returns the Python output directory for a day.
func (g PythonGenerator) GetOutputPath(baseDir, yearPkg string, dd DayTestData) string {
	return filepath.Join(baseDir, yearPkg, fmt.Sprintf("day%s", dd.PaddedDay))
}

// PrepareOutput uses default no-op.
func (g PythonGenerator) PrepareOutput(outPath string, dd DayTestData, opts GenerationOptions) error {
	return nil
}

// FormatContent runs yapf on the generated Python content to match repo formatting.
func (g PythonGenerator) FormatContent(content, targetPath string) (string, error) {
	cmd := exec.Command("python", "-m", "yapf")
	cmd.Stdin = strings.NewReader(content)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("yapf failed for %s: %v: %s", targetPath, err, stderr.String())
	}
	return stdout.String(), nil
}

// Compile-time check
var _ Generator = (*PythonGenerator)(nil)
