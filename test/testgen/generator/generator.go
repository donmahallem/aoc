// Package generator renders per-day test files from model.TestData.
package generator

import (
	"bytes"
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"text/template"

	"github.com/donmahallem/aoc/test/testgen/model"
)

// Language describes how tests are rendered for one target language.
type Language struct {
	Template   string                                      // template file name inside the template dir
	OutputFile func(d *DayInfo) string                     // output path relative to the output dir
	SourceFile func(d *DayInfo) string                     // header checked under SourceDir when set
	Context    func(d *DayInfo) any                        // template data embedding the day
	Funcs      template.FuncMap                            // extra template helpers
	Expected   func(model.Result, model.ResultType) string // formats an expected value as a literal
	Format     func([]byte) ([]byte, error)                // optional post-processing of rendered output
}

// Languages holds all supported languages keyed by config name.
var Languages = map[string]Language{
	"go":  Go,
	"c":   C,
	"cpp": Cpp,
}

// Options configures a single generation run.
type Options struct {
	TemplateDir       string
	OutputDir         string
	SourceDir         string // when set, only days with a header at SourceDir/HeaderPath are generated
	CreateMissingDirs bool
	DryRun            bool // only log actions
	CheckOnly         bool // fail if generated files are out of date; no writes
}

// Generate renders the tests of every day in data for the given language.
func Generate(name string, lang Language, data model.TestData, opts Options) error {
	tmpl, err := template.New(lang.Template).Funcs(lang.Funcs).ParseFiles(filepath.Join(opts.TemplateDir, lang.Template))
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	for _, year := range slices.Sorted(maps.Keys(data)) {
		for _, dayNum := range slices.Sorted(maps.Keys(data[year])) {
			day := newDay(year, dayNum, name, data[year][dayNum], lang.Expected)
			if opts.SourceDir != "" && lang.SourceFile != nil && !exists(filepath.Join(opts.SourceDir, lang.SourceFile(day))) {
				continue
			}

			target := filepath.Join(opts.OutputDir, lang.OutputFile(day))
			outDir := filepath.Dir(target)
			info, err := os.Stat(outDir)
			switch {
			case os.IsNotExist(err):
				if opts.CheckOnly || !opts.CreateMissingDirs {
					fmt.Printf("  skipping %02d/%02d (%s missing)\n", year, dayNum, outDir)
					continue
				}
			case err != nil:
				return err
			case !info.IsDir():
				return fmt.Errorf("output path %s exists and is not a directory", outDir)
			}
			if opts.DryRun {
				fmt.Printf("[dry-run] would generate %02d/%02d -> %s\n", year, dayNum, target)
				continue
			}

			var buf bytes.Buffer
			if err := tmpl.Execute(&buf, lang.Context(day)); err != nil {
				return fmt.Errorf("failed to execute template for %s: %w", target, err)
			}
			content := buf.Bytes()
			if lang.Format != nil {
				if content, err = lang.Format(content); err != nil {
					return fmt.Errorf("failed to format %s: %w", target, err)
				}
			}

			if opts.CheckOnly {
				existing, err := os.ReadFile(target)
				if err != nil {
					return fmt.Errorf("check failed: missing %s", target)
				}
				if !bytes.Equal(content, existing) {
					return fmt.Errorf("check failed: %s is out of date", target)
				}
				fmt.Printf("  check ok %s\n", target)
				continue
			}

			if err := os.MkdirAll(outDir, 0o755); err != nil {
				return err
			}
			if err := os.WriteFile(target, content, 0o644); err != nil {
				return err
			}
			fmt.Printf("  wrote %s\n", target)
		}
	}
	return nil
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// join formats each item with f and joins them with ", ".
func join[T any](items []T, f func(T) string) string {
	parts := make([]string, len(items))
	for i, v := range items {
		parts[i] = f(v)
	}
	return strings.Join(parts, ", ")
}
