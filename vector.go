package mathrock

import "fmt"

// Vector represents a mathematical vector as a slice of float64 values.
type Vector []float64

// Dot returns the dot product of v and other. 
// It returns an error if the two vectors have different lengths.
func (v Vector) Dot(other Vector) (float64, error) {
	if len(v) != len(other) {
		return 0, fmt.Errorf("mathrock: vectors must have equal length, got %d and %d", len(v), len(other))
	}

	var sum float64
	for i := range v {
		sum += v[i] * other[i]
	}

	return sum, nil
}