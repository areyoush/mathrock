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

// Add returns a sum vector of v and the other.
// It returns an error if two vectors have different lengths.
func (v Vector) Add(other Vector) (Vector, error) {
	if len(v) != len(other) {
		return nil, fmt.Errorf("mathrock: vectors must have equal length got %d and %d", len(v), len(other))
	}

	result := make(Vector, len(v))
	for i := range v {
		result[i] = v[i] + other[i]
	}
	return result, nil
}

// Subtract returns a difference vector of v and the other.
// It returns an error if two vectors have different lengths.
func (v Vector) Subtract(other Vector) (Vector, error) {
	if len(v) != len(other) {
		return nil, fmt.Errorf("mathrock: vectors must have equal length got %d and %d", len(v), len(other))
	}

	result := make(Vector, len(v))
	for i := range v {
		result[i] = v[i] - other[i]
	}
	return result, nil
}

// Scale returns the product vector of v and a scalar
func (v Vector) Scale(scalar float64) Vector {
	result := make(Vector, len(v))
	for i := range v {
		result[i] = v[i] * scalar
	}
	return result
}