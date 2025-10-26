package math

// Interval represents a closed interval [Min, Max]
type Interval[A IntType] struct {
	Min A
	Max A
}

// Size returns the size of the interval
func (iv Interval[A]) Size() A {
	if iv.Max < iv.Min {
		return 0
	}
	return iv.Max - iv.Min + 1
}

// Contains returns true if the interval contains the given value
func (iv Interval[A]) Contains(value A) bool {
	return value >= iv.Min && value <= iv.Max
}

// Overlaps returns true if two intervals overlap
func (iv Interval[A]) Overlaps(other Interval[A]) bool {
	return iv.Min <= other.Max && other.Min <= iv.Max
}

// Intersection returns the intersection of two intervals and a boolean indicating
func (iv Interval[A]) Intersection(other Interval[A]) (Interval[A], bool) {
	if !iv.Overlaps(other) {
		return Interval[A]{}, false
	}
	return Interval[A]{
		Min: max(iv.Min, other.Min),
		Max: min(iv.Max, other.Max),
	}, true
}

// Fix ensures that Min is less than or equal to Max
func (iv *Interval[A]) Fix() {
	if iv.Max < iv.Min {
		iv.Min, iv.Max = iv.Max, iv.Min
	}
}
