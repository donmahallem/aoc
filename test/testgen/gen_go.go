package main

import (
	"fmt"
	"go/format"
	"path/filepath"
	"strconv"
	"strings"
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
func (g GoGenerator) FormatExpected(v interface{}, typeHint *string) string {
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
			parts[i] = g.FormatExpected(elem, typeHint)
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

// ShouldGenerate always returns true for Go (all days have implementations).
func (g GoGenerator) ShouldGenerate(dd DayTestData, yearPkg, sourceDir string) bool { return true }

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
