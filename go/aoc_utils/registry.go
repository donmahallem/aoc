package aoc_utils

import (
	"fmt"
	"io"
)

type Registry struct {
	data map[PartSelector]PartFunc[any]
}

func NewRegistry() Registry {
	return Registry{data: make(map[PartSelector]PartFunc[any])}
}

func (v *Registry) RegisteredParts() []PartSelector {
	parts := make([]PartSelector, 0, len(v.data))
	for k := range v.data {
		parts = append(parts, k)
	}
	return parts
}

// Register accepts only a PartFunc: func(io.Reader) (any, error). This is
// intentionally strict to make signatures explicit and readable.
func (v *Registry) Register(year int, day int, part int, fn PartFunc[any]) {
	key := PartSelector{Year: year, Day: day, Part: part}
	if _, exists := v.data[key]; exists {
		// This is a sanity check to prevent accidental overwrites.
		// The registry should only have one function per part.
		panic(fmt.Sprintf("part already registered: year=%d day=%d part=%d", year, day, part))
	}
	v.data[key] = fn
}

// GetPart returns the registered PartFunc for the selection
func (v *Registry) GetPart(selection PartSelector) (PartFunc[any], bool) {
	fn, ok := v.data[selection]
	if ok {
		return fn, true
	}
	return nil, false
}

// RegisterPart registers a single part for a given day using generic adaptation.
// This allows registering parts individually if needed.
func RegisterPart[T any](v *Registry, year, day, part int, fn func(io.Reader) (T, error)) {
	if part != 1 && part != 2 {
		// Sanity check
		panic("Part must be 1 or 2")
	}
	v.Register(year, day, part, Adapt(fn))
}

// RegisterDay registers both parts for a given day using generic adaptation.
// This helper removes the need for reflection or manual casting.
func RegisterDay[T1 any, T2 any](v *Registry, year, day int, p1 func(io.Reader) (T1, error), p2 func(io.Reader) (T2, error)) {
	RegisterPart(v, year, day, 1, p1)
	RegisterPart(v, year, day, 2, p2)
}
