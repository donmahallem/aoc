package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfigResolvesPaths(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "config.json")
	content := `{"data_path": "../data.json", "languages": {"go": {"output_dir": "out", "template_dir": "tmpl"}}}`
	if err := os.WriteFile(cfgPath, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	cfg, err := loadConfig(cfgPath)
	if err != nil {
		t.Fatalf("loadConfig() error: %v", err)
	}
	if want := filepath.Join(dir, "..", "data.json"); cfg.DataPath != want {
		t.Errorf("DataPath = %q, want %q", cfg.DataPath, want)
	}
	goCfg := cfg.Languages["go"]
	if want := filepath.Join(dir, "out"); goCfg.OutputDir != want {
		t.Errorf("OutputDir = %q, want %q", goCfg.OutputDir, want)
	}
	if want := filepath.Join(dir, "tmpl"); goCfg.TemplateDir != want {
		t.Errorf("TemplateDir = %q, want %q", goCfg.TemplateDir, want)
	}
}

func TestLoadConfigErrors(t *testing.T) {
	dir := t.TempDir()

	t.Run("missing file", func(t *testing.T) {
		if _, err := loadConfig(filepath.Join(dir, "missing.json")); err == nil {
			t.Error("expected error for missing config")
		}
	})

	t.Run("invalid json", func(t *testing.T) {
		path := filepath.Join(dir, "bad.json")
		if err := os.WriteFile(path, []byte("{not json"), 0o644); err != nil {
			t.Fatal(err)
		}
		if _, err := loadConfig(path); err == nil {
			t.Error("expected error for invalid json")
		}
	})
}
