package mathrock

import (
	"fmt"
	"math"
)

// Vector represents a mathematical vector as a slice of float64 values.
type Vector []float64

// Dot returns the dot product of v and other.
// It panics if the two vectors have different lengths.
func (v Vector) Dot(other Vector) float64 {
	if len(v) != len(other) {
		panic(fmt.Sprintf("mathrock: vectors must have equal length, got %d and %d", len(v), len(other)))
	}
	var sum float64
	for i := range v {
		sum += v[i] * other[i]
	}
	return sum
}

// Add returns a sum vector of v and the other.
// It panics if two vectors have different lengths.
func (v Vector) Add(other Vector) Vector {
	if len(v) != len(other) {
		panic(fmt.Sprintf("mathrock: vectors must have equal length, got %d and %d", len(v), len(other)))
	}
	result := make(Vector, len(v))
	for i := range v {
		result[i] = v[i] + other[i]
	}
	return result
}

// Subtract returns a difference vector of v and the other.
// It panics if two vectors have different lengths.
func (v Vector) Subtract(other Vector) Vector {
	if len(v) != len(other) {
		panic(fmt.Sprintf("mathrock: vectors must have equal length, got %d and %d", len(v), len(other)))
	}
	result := make(Vector, len(v))
	for i := range v {
		result[i] = v[i] - other[i]
	}
	return result
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
	return math.Sqrt(v.Dot(v))
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

// EqualsWithTolerance reports whether v and other are equal, treating elements as equal if their difference is within tolerance.
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

// Sum returns the sum of all values in v.
func (v Vector) Sum() float64 {
	var sum float64
	for i := range v {
		sum += v[i]
	}
	return sum
}

// Mean returns the average of all elements in v.
// It returns an error if v is empty.
func (v Vector) Mean() (float64, error) {
	if len(v) == 0 {
		return 0, fmt.Errorf("mathrock: cannot compute mean of an empty vector")
	}
	return v.Sum() / float64(len(v)), nil
}

// Multiply returns a new Vector that is the element-wise product of v and other.
// It panics if the two vectors have different lengths.
func (v Vector) Multiply(other Vector) Vector {
	if len(v) != len(other) {
		panic(fmt.Sprintf("mathrock: vectors must have equal length, got %d and %d", len(v), len(other)))
	}
	result := make(Vector, len(v))
	for i := range v {
		result[i] = v[i] * other[i]
	}
	return result
}

// Distance returns the Euclidean distance between v and other.
// It panics if the two vectors have different lengths.
func (v Vector) Distance(other Vector) float64 {
	return v.Subtract(other).Norm()
}