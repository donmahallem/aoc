package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/donmahallem/aoc/test/testgen/model"
)

func TestGenerate(t *testing.T) {
	tmplDir := t.TempDir()
	outDir := t.TempDir()
	if err := os.WriteFile(filepath.Join(tmplDir, "t.tmpl"), []byte("{{.YearNumber}}/{{.PaddedDay}}"), 0o644); err != nil {
		t.Fatal(err)
	}
	lang := Language{
		Template:   "t.tmpl",
		OutputFile: func(d *DayInfo) string { return filepath.Join(d.YearDir, dayPkg(d)+".txt") },
		Context:    func(d *DayInfo) any { return d },
		Expected:   goExpected,
	}
	data := model.TestData{2024: {1: {"a": {Input: new("x"), Part1: partExp(1)}}}}
	opts := Options{TemplateDir: tmplDir, OutputDir: outDir, CreateMissingDirs: true}
	target := filepath.Join(outDir, "aoc2024", "day01.txt")

	t.Run("writes output", func(t *testing.T) {
		if err := Generate("go", lang, data, opts); err != nil {
			t.Fatalf("Generate() error: %v", err)
		}
		got, err := os.ReadFile(target)
		if err != nil {
			t.Fatalf("failed to read output: %v", err)
		}
		if string(got) != "2024/01" {
			t.Errorf("output = %q, want %q", got, "2024/01")
		}
	})

	t.Run("check passes when up to date", func(t *testing.T) {
		checkOpts := opts
		checkOpts.CheckOnly = true
		if err := Generate("go", lang, data, checkOpts); err != nil {
			t.Errorf("Generate() in check mode error: %v", err)
		}
	})

	t.Run("check fails when stale", func(t *testing.T) {
		if err := os.WriteFile(target, []byte("stale"), 0o644); err != nil {
			t.Fatal(err)
		}
		checkOpts := opts
		checkOpts.CheckOnly = true
		err := Generate("go", lang, data, checkOpts)
		if err == nil || !strings.Contains(err.Error(), "out of date") {
			t.Errorf("Generate() error = %v, want out of date", err)
		}
	})
}
