package model_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/donmahallem/aoc/test/testgen/model"
)

func TestTestCaseFetch(t *testing.T) {
	t.Run("inline input", func(t *testing.T) {
		input := "inline"
		tc := model.TestCase{Input: &input}
		got, err := tc.Fetch()
		if err != nil || string(got) != "inline" {
			t.Errorf("Fetch() = %q, %v; want %q, nil", got, err, "inline")
		}
	})

	t.Run("file input", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "in.txt")
		if err := os.WriteFile(path, []byte("from file"), 0o644); err != nil {
			t.Fatal(err)
		}
		tc := model.TestCase{File: &path}
		got, err := tc.Fetch()
		if err != nil || string(got) != "from file" {
			t.Errorf("Fetch() = %q, %v; want %q, nil", got, err, "from file")
		}
	})

	t.Run("missing file", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "missing.txt")
		tc := model.TestCase{File: &path}
		if _, err := tc.Fetch(); err == nil {
			t.Error("expected error for missing file")
		}
	})

	t.Run("file relative to loaded data file", func(t *testing.T) {
		dir := t.TempDir()
		if err := os.WriteFile(filepath.Join(dir, "in.txt"), []byte("relative"), 0o644); err != nil {
			t.Fatal(err)
		}
		dataPath := filepath.Join(dir, "data.json")
		content := `{"2024": {"01": {"a": {"file": "in.txt", "fullData": true}}}}`
		if err := os.WriteFile(dataPath, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
		data, err := model.Load(dataPath)
		if err != nil {
			t.Fatalf("Load() error: %v", err)
		}
		tc := data[2024][1]["a"]
		got, err := tc.Fetch()
		if err != nil || string(got) != "relative" {
			t.Errorf("Fetch() = %q, %v; want %q, nil", got, err, "relative")
		}
	})

	t.Run("no input or file", func(t *testing.T) {
		var tc model.TestCase
		if _, err := tc.Fetch(); err == nil {
			t.Error("expected error when neither input nor file is set")
		}
	})
}
