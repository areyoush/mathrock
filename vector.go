package mathrock

import (
	"fmt"
	"math"
)

// Vector represents a mathematical vector as a slice of float64 values.
type Vector []float64

// Dot returns the dot product vector of v and other. 
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

// Scale returns the product vector of v and a scalar.
func (v Vector) Scale(scalar float64) Vector {
	result := make(Vector, len(v))
	for i := range v {
		result[i] = v[i] * scalar
	}
	return result
}

// Norm returns the Euclidean length (L2 norm) of v.
func (v Vector) Norm() float64 {
	sumOfSquares, _ := v.Dot(v)
	return math.Sqrt(sumOfSquares)
}

// Normalize returns a new Vector with the same direction as v, scaled to length 1. 
// It returns an error if v is the zero vector.
func (v Vector) Normalize() (Vector, error) {
	norm := v.Norm()
	if norm == 0 {
		return nil, fmt.Errorf("mathrock: cannot normalize the zero vector")
	}
	return v.Scale(1 / norm), nil
}

// EqualsWithTolerance reports whether v and other are equal, treating elements as equal if their differene is within tolerance.
// Vectors of different lengths are never equal.
func (v Vector) EqualsWithTolerance(other Vector, tolerance float64) bool {
	if len(v) != len(other) {
		return false
	}
	for i := range v {
		if math.Abs(v[i]-other[i]) > tolerance {
			return false
		}
	}
	return true
}

// Equals reports whether v and the other are equal, using a small default tolerance to account for floating-point rounding error.
func (v Vector) Equals(other Vector) bool {
	return v.EqualsWithTolerance(other, 1e-9)
}