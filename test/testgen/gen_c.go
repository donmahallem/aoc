package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func init() {
	RegisterGenerator("c", &CGenerator{})
}

type CGenerator struct {
	lastPaddedDay string
}

// Template context for C test file generation.
type CTestFileData struct {
	BaseTemplateData
	HeaderPath string
}

func (g *CGenerator) TemplateName() string { return "generated_test.c.tmpl" }

// OutputFilename is day-specific (e.g. day01_generated.c) to avoid per-day subdirectories.
func (g *CGenerator) OutputFilename() string {
	if g.lastPaddedDay == "" {
		return "generated_test.c"
	}
	return fmt.Sprintf("day%s_generated.c", g.lastPaddedDay)
}

func (g *CGenerator) FuncMap() template.FuncMap {
	return template.FuncMap{
		"cEscapeString": func(s string) string {
			// produce a quoted C string literal with escapes
			b := strings.Builder{}
			b.WriteString("\"")
			for _, r := range s {
				switch r {
				case '\\':
					b.WriteString("\\\\")
				case '"':
					b.WriteString("\\\"")
				case '\n':
					b.WriteString("\\n")
				case '\r':
					b.WriteString("\\r")
				case '\t':
					b.WriteString("\\t")
				default:
					b.WriteRune(r)
				}
			}
			b.WriteString("\"")
			return b.String()
		},
		"join": strings.Join,
	}
}

func (g *CGenerator) FormatExpected(v interface{}, typeHint *string) string {
	return FormatExpectedC(v, typeHint)
}

func (g *CGenerator) GetTemplateData(dd DayTestData, yearPkg string) interface{} {
	// remember padded day so OutputFilename can include it
	g.lastPaddedDay = dd.PaddedDay
	// header path keeps same structure as C source
	dayPkg := fmt.Sprintf("day%s", dd.PaddedDay)
	headerPath := fmt.Sprintf("%s/%s/%s.h", yearPkg, dayPkg, dayPkg)
	return CTestFileData{
		BaseTemplateData: BuildBaseTemplateData(dd, yearPkg),
		HeaderPath:       headerPath,
	}
}

func (g *CGenerator) GetOutputPath(baseDir, yearPkg string, dd DayTestData) string {
	return filepath.Join(baseDir, yearPkg)
}

func (g CGenerator) PrepareOutput(outPath string, dd DayTestData, opts GenerationOptions) error {
	return nil
}

// Only generate tests when header for the day exists in source_dir
func (g CGenerator) ShouldGenerate(dd DayTestData, yearPkg, sourceDir string) bool {
	if sourceDir == "" {
		return true
	}
	dayPkg := fmt.Sprintf("day%s", dd.PaddedDay)
	header := filepath.Join(sourceDir, yearPkg, dayPkg, dayPkg+".h")
	_, err := os.Stat(header)
	return err == nil
}

// keep content unchanged
func (g CGenerator) FormatContent(content, targetPath string) (string, error) { return content, nil }

var _ Generator = (*CGenerator)(nil)
