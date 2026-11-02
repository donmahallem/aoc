package model

import (
	"fmt"
	"os"
	"path/filepath"
)

// TestCase provides either an inline Input or a File path, with expectations for at least one part.
type TestCase struct {
	fileCache     []byte
	baseDir       string
	Input         *string          `json:"input,omitempty"`
	File          *string          `json:"file,omitempty"`
	Part1         *PartExpectation `json:"part1,omitempty"`
	Part2         *PartExpectation `json:"part2,omitempty"`
	SkipLanguages []string         `json:"skip_languages,omitempty"`
	// FullData marks a full data case. Cases without it are samples.
	FullData bool `json:"fullData,omitempty"`
}

// IsSample reports whether the case is a sample, i.e. not marked as full data.
func (r TestCase) IsSample() bool {
	return !r.FullData
}

// use this to get the test case input as a byte slice, either from the Input field or the File field.
func (r *TestCase) Fetch() ([]byte, error) {
	if r.fileCache != nil {
		return r.fileCache, nil
	}
	if r.Input != nil {
		r.fileCache = []byte(*r.Input)
		return r.fileCache, nil
	}
	if r.File != nil {
		data, err := os.ReadFile(r.resolvedFile())
		if err != nil {
			return nil, err
		}
		r.fileCache = data
		return r.fileCache, nil
	}
	return nil, fmt.Errorf("no input or file specified")
}

// resolvedFile joins a relative File with the data file's directory; without Load, it stays relative to the working directory.
func (r *TestCase) resolvedFile() string {
	if r.baseDir == "" || filepath.IsAbs(*r.File) {
		return *r.File
	}
	return filepath.Join(r.baseDir, *r.File)
}
