package mathrock

import (
	"testing"
)

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

// TestNewMatrix verifies matrix construction succeeds when data length matches rows*cols, and returns an error when it doesn't.
func TestNewMatrix(t *testing.T) {
	tests := []struct {
		name    string
		rows    int
		cols    int
		data    []float64
		wantErr bool
	}{
		{
			name: "valid 2x3 matrix",
			rows: 2,
			cols: 3,
			data: []float64{1, 2, 3, 4, 5, 6},
		},
		{
			name: "valid 1x1 matrix",
			rows: 1,
			cols: 1,
			data: []float64{42},
		},
		{
			name:    "data length too short",
			rows:    2,
			cols:    3,
			data:    []float64{1, 2, 3, 4, 5},
			wantErr: true,
		},
		{
			name:    "data length too long",
			rows:    2,
			cols:    2,
			data:    []float64{1, 2, 3, 4, 5},
			wantErr: true,
		},
		{
			name: "zero rows and cols with empty data",
			rows: 0,
			cols: 0,
			data: []float64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := NewMatrix(tt.rows, tt.cols, tt.data)

			if (err != nil) != tt.wantErr {
				t.Fatalf("NewMatrix() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			if m.rows != tt.rows || m.cols != tt.cols {
				t.Errorf("NewMatrix() rows/cols = %d/%d, want %d/%d", m.rows, m.cols, tt.rows, tt.cols)
			}
		})
	}
}

// TestMatrixAt verifies At returns the correct value for valid indices and panics when row or col is out of bounds.
func TestMatrixAt(t *testing.T) {
	tests := []struct {
		name      string
		rows      int
		cols      int
		data      []float64
		row       int
		col       int
		want      float64
		wantPanic bool
	}{
		{
			name: "top-left element",
			rows: 2,
			cols: 3,
			data: []float64{1, 2, 3, 4, 5, 6},
			row:  0,
			col:  0,
			want: 1,
		},
		{
			name: "bottom-right element",
			rows: 2,
			cols: 3,
			data: []float64{1, 2, 3, 4, 5, 6},
			row:  1,
			col:  2,
			want: 6,
		},
		{
			name: "middle element",
			rows: 2,
			cols: 3,
			data: []float64{1, 2, 3, 4, 5, 6},
			row:  1,
			col:  0,
			want: 4,
		},
		{
			name:      "row out of bounds negative",
			rows:      2,
			cols:      3,
			data:      []float64{1, 2, 3, 4, 5, 6},
			row:       -1,
			col:       0,
			wantPanic: true,
		},
		{
			name:      "row out of bounds too large",
			rows:      2,
			cols:      3,
			data:      []float64{1, 2, 3, 4, 5, 6},
			row:       2,
			col:       0,
			wantPanic: true,
		},
		{
			name:      "col out of bounds negative",
			rows:      2,
			cols:      3,
			data:      []float64{1, 2, 3, 4, 5, 6},
			row:       0,
			col:       -1,
			wantPanic: true,
		},
		{
			name:      "col out of bounds too large",
			rows:      2,
			cols:      3,
			data:      []float64{1, 2, 3, 4, 5, 6},
			row:       0,
			col:       3,
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := NewMatrix(tt.rows, tt.cols, tt.data)
			if err != nil {
				t.Fatalf("NewMatrix() unexpected error = %v", err)
			}

			if tt.wantPanic {
				assertPanics(t, func() { m.At(tt.row, tt.col) })
				return
			}

			got := m.At(tt.row, tt.col)
			if got != tt.want {
				t.Errorf("At(%d, %d) = %v, want %v", tt.row, tt.col, got, tt.want)
			}
		})
	}
}

// TestMatrixSet verifies Set correctly writes a value at a given position (checked via At) and panics when row or col is out of bounds.
func TestMatrixSet(t *testing.T) {
	tests := []struct {
		name      string
		rows      int
		cols      int
		data      []float64
		row       int
		col       int
		val       float64
		wantPanic bool
	}{
		{
			name: "set top-left element",
			rows: 2,
			cols: 3,
			data: []float64{1, 2, 3, 4, 5, 6},
			row:  0,
			col:  0,
			val:  99,
		},
		{
			name: "set bottom-right element",
			rows: 2,
			cols: 3,
			data: []float64{1, 2, 3, 4, 5, 6},
			row:  1,
			col:  2,
			val:  -5,
		},
		{
			name:      "row out of bounds",
			rows:      2,
			cols:      3,
			data:      []float64{1, 2, 3, 4, 5, 6},
			row:       5,
			col:       0,
			val:       1,
			wantPanic: true,
		},
		{
			name:      "col out of bounds",
			rows:      2,
			cols:      3,
			data:      []float64{1, 2, 3, 4, 5, 6},
			row:       0,
			col:       5,
			val:       1,
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := NewMatrix(tt.rows, tt.cols, tt.data)
			if err != nil {
				t.Fatalf("NewMatrix() unexpected error = %v", err)
			}

			if tt.wantPanic {
				assertPanics(t, func() { m.Set(tt.row, tt.col, tt.val) })
				return
			}

			m.Set(tt.row, tt.col, tt.val)

			got := m.At(tt.row, tt.col)
			if got != tt.val {
				t.Errorf("after Set(%d, %d, %v), At() = %v, want %v", tt.row, tt.col, tt.val, got, tt.val)
			}
		})
	}
}