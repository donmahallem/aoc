package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// GenerationOptions holds cross-language options.
type GenerationOptions struct {
	DryRun            bool // when true, only logs actions without writing files
	CheckOnly         bool // when true, fails if generated content would differ; no writes
	CreateMissingDirs bool // when true, creates output directories as needed
}

// Generator defines the interface for language-specific test generators.
// Implementors can embed BaseGenerator to pick up sensible defaults.
type Generator interface {
	TemplateName() string
	OutputFilename() string
	FuncMap() template.FuncMap
	FormatExpected(v interface{}, typeHint *string) string
	GetTemplateData(dd DayTestData, yearPkg string) interface{}
	GetOutputPath(outDir, yearPkg string, dd DayTestData) string
	PrepareOutput(outPath string, dd DayTestData, opts GenerationOptions) error
	FormatContent(content, targetPath string) (string, error)
	// ShouldGenerate returns whether a test file should be generated for the given day.
	// Generators can use this to skip days with no source implementation.
	// sourceDir is the resolved source_dir from config (empty if not set).
	ShouldGenerate(dd DayTestData, yearPkg, sourceDir string) bool
}

// BaseGenerator provides default no-op implementations for optional pieces.
type BaseGenerator struct{}

func (BaseGenerator) FuncMap() template.FuncMap                                  { return nil }
func (BaseGenerator) PrepareOutput(string, DayTestData, GenerationOptions) error { return nil }
func (BaseGenerator) FormatContent(content, _ string) (string, error)            { return content, nil }
func (BaseGenerator) ShouldGenerate(DayTestData, string, string) bool            { return true }

// Generators registry
var generators = make(map[string]Generator)

func RegisterGenerator(name string, g Generator) {
	generators[name] = g
}

// GenerateGeneric implements the common logic for generating tests using a Generator.
func GenerateGeneric(gen Generator, lang string, data TestData, tmplDir, outDir, sourceDir string, opts GenerationOptions) error {
	tmpl := template.New(gen.TemplateName())
	if fm := gen.FuncMap(); fm != nil {
		tmpl = tmpl.Funcs(fm)
	}

	tmpl, err := tmpl.ParseFiles(filepath.Join(tmplDir, gen.TemplateName()))
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	for _, yearStr := range SortedKeys(data) {
		days := data[yearStr]
		yearPkg := fmt.Sprintf("aoc%s", yearStr)

		for _, dayStr := range SortedKeys(days) {
			cases := days[dayStr]
			dd := ParseDayData(yearStr, dayStr, lang, cases, gen.FormatExpected)

			if !gen.ShouldGenerate(dd, yearPkg, sourceDir) {
				continue
			}

			outPath := gen.GetOutputPath(outDir, yearPkg, dd)
			if opts.DryRun {
				fmt.Printf("[dry-run] would generate %s/%s -> %s\n", yearStr, dayStr, filepath.Join(outPath, gen.OutputFilename()))
				continue
			}

			td := gen.GetTemplateData(dd, yearPkg)
			content, err := renderTemplate(tmpl, td, filepath.Join(outPath, gen.OutputFilename()))
			if err != nil {
				return err
			}
			content, err = gen.FormatContent(content, filepath.Join(outPath, gen.OutputFilename()))
			if err != nil {
				return err
			}

			targetPath := filepath.Join(outPath, gen.OutputFilename())

			if opts.CheckOnly {
				existing, err := os.ReadFile(targetPath)
				if err != nil {
					return fmt.Errorf("check failed: missing %s", targetPath)
				}
				if content != string(existing) {
					return fmt.Errorf("check failed: %s is out of date", targetPath)
				}
				fmt.Printf("  check ok %s\n", targetPath)
				continue
			}

			if err := ensureDir(outPath, opts.CreateMissingDirs); err != nil {
				return err
			}
			if err := gen.PrepareOutput(outPath, dd, opts); err != nil {
				return fmt.Errorf("failed to prepare output for %s/%s: %w", yearStr, dayStr, err)
			}

			if err := writeTemplateContent(content, targetPath); err != nil {
				return err
			}

			fmt.Printf("Generated tests for year %s day %s\n", yearStr, dayStr)
		}
	}
	return nil
}

func renderTemplate(tmpl *template.Template, data interface{}, path string) (string, error) {
	var buf strings.Builder
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template for %s: %w", path, err)
	}
	return buf.String(), nil
}

func writeTemplateContent(content, path string) error {
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write %s: %w", path, err)
	}
	fmt.Printf("  wrote %s\n", path)
	return nil
}

func ensureDir(path string, allowCreate bool) error {
	info, err := os.Stat(path)
	if err == nil {
		if info.IsDir() {
			return nil
		}
		return fmt.Errorf("output path %s exists and is not a directory", path)
	}
	if !os.IsNotExist(err) {
		return fmt.Errorf("failed to stat %s: %w", path, err)
	}
	if !allowCreate {
		return fmt.Errorf("output directory %s does not exist (set create_missing_dirs to true to create)", path)
	}
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", path, err)
	}
	return nil
}
