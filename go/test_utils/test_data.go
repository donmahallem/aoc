// Package test_utils provides helpers for the generated and hand-written Go tests.
// Test data is described by the testgen model package.
package test_utils

import (
	"maps"
	"path/filepath"
	"slices"
	"sync"
	"testing"

	"github.com/donmahallem/aoc/test/testgen/model"
)

func DefaultTestDataPath() string {
	return filepath.Join(basepath, "../../test/data.json")
}

var defaultTestData = sync.OnceValues(func() (model.TestData, error) {
	return model.Load(DefaultTestDataPath())
})

// LoadDefaultTestData returns the parsed test/data.json, parsed once per test binary.
func LoadDefaultTestData(tb testing.TB) model.TestData {
	tb.Helper()
	data, err := defaultTestData()
	if err != nil {
		tb.Fatal(err)
	}
	return data
}

// EachTestCase calls fn for every test case in ascending year, day and name order.
func EachTestCase(data model.TestData, fn func(year, day int, name string, tc model.TestCase)) {
	for _, year := range slices.Sorted(maps.Keys(data)) {
		days := data[year]
		for _, day := range slices.Sorted(maps.Keys(days)) {
			cases := days[day]
			for _, name := range slices.Sorted(maps.Keys(cases)) {
				fn(year, day, name, cases[name])
			}
		}
	}
}
