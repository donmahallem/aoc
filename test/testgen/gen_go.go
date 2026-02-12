package main

import (
	"fmt"
	"go/format"
	"path/filepath"
	"text/template"
)

func init() {
	RegisterGenerator("go", &GoGenerator{})
}

type GoGenerator struct{}

// GoTestFileData is the template context for Go test file generation.
// It embeds shared fields and adds Go-specific package naming.
type GoTestFileData struct {
	BaseTemplateData
	GoPackage string // "day01"
}

// TemplateName returns the template filename for Go tests.
func (g GoGenerator) TemplateName() string { return "generated_test.go.tmpl" }

// OutputFilename returns the output test filename for Go.
func (g GoGenerator) OutputFilename() string { return "generated_test.go" }

// FuncMap provides Go-specific template helpers.
func (g GoGenerator) FuncMap() template.FuncMap {
	return template.FuncMap{
		"backtick": func() string { return "`" },
	}
}

// FormatExpected formats expected values for Go literals.
func (g GoGenerator) FormatExpected(v interface{}) string { return FormatExpectedGo(v) }

// GetTemplateData returns the Go template data.
func (g GoGenerator) GetTemplateData(dd DayTestData, yearPkg string) interface{} {
	return GoTestFileData{
		BaseTemplateData: BuildBaseTemplateData(dd, yearPkg),
		GoPackage:        fmt.Sprintf("day%s", dd.PaddedDay),
	}
}

// GetOutputPath returns the Go output directory for a day.
func (g GoGenerator) GetOutputPath(baseDir, yearPkg string, dd DayTestData) string {
	goPackage := fmt.Sprintf("day%s", dd.PaddedDay)
	return filepath.Join(baseDir, yearPkg, goPackage)
}

// PrepareOutput uses default no-op from BaseGenerator.
func (g GoGenerator) PrepareOutput(outPath string, dd DayTestData, opts GenerationOptions) error {
	return nil
}

// FormatContent formats generated Go source using go/format for consistent style.
func (g GoGenerator) FormatContent(content, targetPath string) (string, error) {
	formatted, err := format.Source([]byte(content))
	if err != nil {
		return "", fmt.Errorf("gofmt failed for %s: %w", targetPath, err)
	}
	return string(formatted), nil
}

// Compile-time check
var _ Generator = (*GoGenerator)(nil)
