package mathrock

import (
	"math"
	"reflect"
	"testing"
)

func vectorsApproxEqual(a, b Vector, tolerance float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.Abs(a[i]-b[i]) > tolerance {
			return false
		}
	}
	return true
}

// assertPanics runs fn and fails the test if it does not panic.
func assertPanics(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic, but function did not panic")
		}
	}()
	fn()
}

func TestVectorDot(t *testing.T) {
	tests := []struct {
		name      string
		a         Vector
		b         Vector
		want      float64
		wantPanic bool
	}{
		{
			name: "equal length vectors",
			a:    Vector{1, 2, 3},
			b:    Vector{4, 5, 6},
			want: 32,
		},
		{
			name:      "mismatched length vectors",
			a:         Vector{1, 2},
			b:         Vector{1, 2, 3},
			wantPanic: true,
		},
		{
			name: "empty vectors",
			a:    Vector{},
			b:    Vector{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assertPanics(t, func() { tt.a.Dot(tt.b) })
				return
			}

			got := tt.a.Dot(tt.b)
			if got != tt.want {
				t.Errorf("Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVectorAdd(t *testing.T) {
	tests := []struct {
		name      string
		a         Vector
		b         Vector
		want      Vector
		wantPanic bool
	}{
		{
			name: "equal length vectors",
			a:    Vector{1, 2, 3},
			b:    Vector{4, 5, 6},
			want: Vector{5, 7, 9},
		},
		{
			name:      "mismatched length vectors",
			a:         Vector{1, 2},
			b:         Vector{1, 2, 3},
			wantPanic: true,
		},
		{
			name: "empty vectors",
			a:    Vector{},
			b:    Vector{},
			want: Vector{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assertPanics(t, func() { tt.a.Add(tt.b) })
				return
			}

			got := tt.a.Add(tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVectorSubtract(t *testing.T) {
	tests := []struct {
		name      string
		a         Vector
		b         Vector
		want      Vector
		wantPanic bool
	}{
		{
			name: "equal length vectors",
			a:    Vector{4, 5, 6},
			b:    Vector{1, 2, 3},
			want: Vector{3, 3, 3},
		},
		{
			name:      "mismatched length vectors",
			a:         Vector{1, 2},
			b:         Vector{1, 2, 3},
			wantPanic: true,
		},
		{
			name: "empty vectors",
			a:    Vector{},
			b:    Vector{},
			want: Vector{},
		},
		{
			name: "result has negative values",
			a:    Vector{1, 2, 3},
			b:    Vector{4, 5, 6},
			want: Vector{-3, -3, -3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assertPanics(t, func() { tt.a.Subtract(tt.b) })
				return
			}

			got := tt.a.Subtract(tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVectorScale(t *testing.T) {
	tests := []struct {
		name   string
		v      Vector
		scalar float64
		want   Vector
	}{
		{
			name:   "positive scalar",
			v:      Vector{1, 2, 3},
			scalar: 2,
			want:   Vector{2, 4, 6},
		},
		{
			name:   "negative scalar",
			v:      Vector{1, 2, 3},
			scalar: -1,
			want:   Vector{-1, -2, -3},
		},
		{
			name:   "zero scalar",
			v:      Vector{1, 2, 3},
			scalar: 0,
			want:   Vector{0, 0, 0},
		},
		{
			name:   "empty vector",
			v:      Vector{},
			scalar: 5,
			want:   Vector{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.v.Scale(tt.scalar)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVectorNorm(t *testing.T) {
	tests := []struct {
		name string
		v    Vector
		want float64
	}{
		{
			name: "3-4-5 triangle",
			v:    Vector{3, 4},
			want: 5,
		},
		{
			name: "zero vector",
			v:    Vector{0, 0, 0},
			want: 0,
		},
		{
			name: "single element",
			v:    Vector{7},
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.v.Norm()
			if got != tt.want {
				t.Errorf("Norm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVectorNormalize(t *testing.T) {
	tests := []struct {
		name    string
		v       Vector
		want    Vector
		wantErr bool
	}{
		{
			name:    "3-4-5 triangle",
			v:       Vector{3, 4},
			want:    Vector{0.6, 0.8},
			wantErr: false,
		},
		{
			name:    "zero vector",
			v:       Vector{0, 0, 0},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Normalize()

			if (err != nil) != tt.wantErr {
				t.Fatalf("Normalize() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			if !vectorsApproxEqual(got, tt.want, 1e-9) {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVectorEqualsWithTolerance(t *testing.T) {
	tests := []struct {
		name      string
		a         Vector
		b         Vector
		tolerance float64
		want      bool
	}{
		{
			name:      "identical vectors",
			a:         Vector{1, 2, 3},
			b:         Vector{1, 2, 3},
			tolerance: 1e-9,
			want:      true,
		},
		{
			name:      "different vectors",
			a:         Vector{1, 2, 3},
			b:         Vector{1, 2, 4},
			tolerance: 1e-9,
			want:      false,
		},
		{
			name:      "within tolerance",
			a:         Vector{1, 2, 3},
			b:         Vector{1.0001, 2, 3},
			tolerance: 0.001,
			want:      true,
		},
		{
			name:      "outside tolerance",
			a:         Vector{1, 2, 3},
			b:         Vector{1.1, 2, 3},
			tolerance: 0.001,
			want:      false,
		},
		{
			name:      "mismatched lengths",
			a:         Vector{1, 2},
			b:         Vector{1, 2, 3},
			tolerance: 1e-9,
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.EqualsWithTolerance(tt.b, tt.tolerance)
			if got != tt.want {
				t.Errorf("EqualsWithTolerance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVectorEquals(t *testing.T) {
	tests := []struct {
		name string
		a    Vector
		b    Vector
		want bool
	}{
		{
			name: "identical vectors",
			a:    Vector{1, 2, 3},
			b:    Vector{1, 2, 3},
			want: true,
		},
		{
			name: "different vectors",
			a:    Vector{1, 2, 3},
			b:    Vector{1, 2, 4},
			want: false,
		},
		{
			name: "floating point noise",
			a:    Vector{0.6, 0.8},
			b:    Vector{0.6000000000000001, 0.8},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.Equals(tt.b)
			if got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVectorSum(t *testing.T) {
	tests := []struct {
		name string
		v    Vector
		want float64
	}{
		{
			name: "positive numbers",
			v:    Vector{1, 2, 3},
			want: 6,
		},
		{
			name: "includes negatives",
			v:    Vector{1, -2, 3},
			want: 2,
		},
		{
			name: "empty vector",
			v:    Vector{},
			want: 0,
		},
		{
			name: "single element",
			v:    Vector{5},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.v.Sum()
			if got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVectorMean(t *testing.T) {
	tests := []struct {
		name    string
		v       Vector
		want    float64
		wantErr bool
	}{
		{
			name:    "positive numbers",
			v:       Vector{1, 2, 3},
			want:    2,
			wantErr: false,
		},
		{
			name:    "includes negatives",
			v:       Vector{1, -1, 4},
			want:    1.333333333333333,
			wantErr: false,
		},
		{
			name:    "empty vector",
			v:       Vector{},
			want:    0,
			wantErr: true,
		},
		{
			name:    "single element",
			v:       Vector{7},
			want:    7,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Mean()

			if (err != nil) != tt.wantErr {
				t.Fatalf("Mean() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("Mean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVectorMultiply(t *testing.T) {
	tests := []struct {
		name      string
		a         Vector
		b         Vector
		want      Vector
		wantPanic bool
	}{
		{
			name: "equal length vectors",
			a:    Vector{1, 2, 3},
			b:    Vector{4, 5, 6},
			want: Vector{4, 10, 18},
		},
		{
			name:      "mismatched length vectors",
			a:         Vector{1, 2},
			b:         Vector{1, 2, 3},
			wantPanic: true,
		},
		{
			name: "empty vectors",
			a:    Vector{},
			b:    Vector{},
			want: Vector{},
		},
		{
			name: "includes negative and zero",
			a:    Vector{-2, 0, 3},
			b:    Vector{5, 9, -1},
			want: Vector{-10, 0, -3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assertPanics(t, func() { tt.a.Multiply(tt.b) })
				return
			}

			got := tt.a.Multiply(tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVectorDistance(t *testing.T) {
	tests := []struct {
		name      string
		a         Vector
		b         Vector
		want      float64
		wantPanic bool
	}{
		{
			name: "3-4-5 triangle",
			a:    Vector{0, 0},
			b:    Vector{3, 4},
			want: 5,
		},
		{
			name: "same vector",
			a:    Vector{1, 2, 3},
			b:    Vector{1, 2, 3},
			want: 0,
		},
		{
			name:      "mismatched lengths",
			a:         Vector{1, 2},
			b:         Vector{1, 2, 3},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assertPanics(t, func() { tt.a.Distance(tt.b) })
				return
			}

			got := tt.a.Distance(tt.b)
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}