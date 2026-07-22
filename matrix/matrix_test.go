package mathrock

import (
	"reflect"
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

// TestMatrixRow verifies Row returns the correct slice of values for a given row and panics when row is out of bounds.
func TestMatrixRow(t *testing.T) {
	tests := []struct {
		name      string
		rows      int
		cols      int
		data      []float64
		row       int
		want      []float64
		wantPanic bool
	}{
		{
			name: "first row",
			rows: 2,
			cols: 3,
			data: []float64{1, 2, 3, 4, 5, 6},
			row:  0,
			want: []float64{1, 2, 3},
		},
		{
			name: "second row",
			rows: 2,
			cols: 3,
			data: []float64{1, 2, 3, 4, 5, 6},
			row:  1,
			want: []float64{4, 5, 6},
		},
		{
			name: "single column matrix",
			rows: 3,
			cols: 1,
			data: []float64{7, 8, 9},
			row:  1,
			want: []float64{8},
		},
		{
			name:      "row out of bounds negative",
			rows:      2,
			cols:      3,
			data:      []float64{1, 2, 3, 4, 5, 6},
			row:       -1,
			wantPanic: true,
		},
		{
			name:      "row out of bounds too large",
			rows:      2,
			cols:      3,
			data:      []float64{1, 2, 3, 4, 5, 6},
			row:       2,
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
				assertPanics(t, func() { m.Row(tt.row) })
				return
			}

			got := m.Row(tt.row)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Row(%d) = %v, want %v", tt.row, got, tt.want)
			}
		})
	}
}

// TestMatrixAdd verifies element-wise Add on same-shape matrices and confirms it panics on mismatched dimensions.
func TestMatrixAdd(t *testing.T) {
	tests := []struct {
		name       string
		aRows      int
		aCols      int
		aData      []float64
		bRows      int
		bCols      int
		bData      []float64
		want       []float64
		wantPanic  bool
	}{
		{
			name:  "same shape matrices",
			aRows: 2, aCols: 2, aData: []float64{1, 2, 3, 4},
			bRows: 2, bCols: 2, bData: []float64{5, 6, 7, 8},
			want: []float64{6, 8, 10, 12},
		},
		{
			name:  "includes negatives",
			aRows: 1, aCols: 3, aData: []float64{1, -2, 3},
			bRows: 1, bCols: 3, bData: []float64{-1, 2, -3},
			want: []float64{0, 0, 0},
		},
		{
			name:      "mismatched rows",
			aRows:     2, aCols: 2, aData: []float64{1, 2, 3, 4},
			bRows:     3, bCols: 2, bData: []float64{1, 2, 3, 4, 5, 6},
			wantPanic: true,
		},
		{
			name:      "mismatched cols",
			aRows:     2, aCols: 2, aData: []float64{1, 2, 3, 4},
			bRows:     2, bCols: 3, bData: []float64{1, 2, 3, 4, 5, 6},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, err := NewMatrix(tt.aRows, tt.aCols, tt.aData)
			if err != nil {
				t.Fatalf("NewMatrix() unexpected error = %v", err)
			}
			b, err := NewMatrix(tt.bRows, tt.bCols, tt.bData)
			if err != nil {
				t.Fatalf("NewMatrix() unexpected error = %v", err)
			}

			if tt.wantPanic {
				assertPanics(t, func() { a.Add(b) })
				return
			}

			got := a.Add(b)
			if !reflect.DeepEqual(got.data, tt.want) {
				t.Errorf("Add() = %v, want %v", got.data, tt.want)
			}
		})
	}
}

// TestMatrixSubtract verifies element-wise Subtract on same-shape matrices and confirms it panics on mismatched dimensions.
func TestMatrixSubtract(t *testing.T) {
	tests := []struct {
		name      string
		aRows     int
		aCols     int
		aData     []float64
		bRows     int
		bCols     int
		bData     []float64
		want      []float64
		wantPanic bool
	}{
		{
			name:  "same shape matrices",
			aRows: 2, aCols: 2, aData: []float64{5, 6, 7, 8},
			bRows: 2, bCols: 2, bData: []float64{1, 2, 3, 4},
			want: []float64{4, 4, 4, 4},
		},
		{
			name:  "result has negative values",
			aRows: 1, aCols: 3, aData: []float64{1, 2, 3},
			bRows: 1, bCols: 3, bData: []float64{4, 5, 6},
			want: []float64{-3, -3, -3},
		},
		{
			name:      "mismatched rows",
			aRows:     2, aCols: 2, aData: []float64{1, 2, 3, 4},
			bRows:     3, bCols: 2, bData: []float64{1, 2, 3, 4, 5, 6},
			wantPanic: true,
		},
		{
			name:      "mismatched cols",
			aRows:     2, aCols: 2, aData: []float64{1, 2, 3, 4},
			bRows:     2, bCols: 3, bData: []float64{1, 2, 3, 4, 5, 6},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, err := NewMatrix(tt.aRows, tt.aCols, tt.aData)
			if err != nil {
				t.Fatalf("NewMatrix() unexpected error = %v", err)
			}
			b, err := NewMatrix(tt.bRows, tt.bCols, tt.bData)
			if err != nil {
				t.Fatalf("NewMatrix() unexpected error = %v", err)
			}

			if tt.wantPanic {
				assertPanics(t, func() { a.Subtract(b) })
				return
			}

			got := a.Subtract(b)
			if !reflect.DeepEqual(got.data, tt.want) {
				t.Errorf("Subtract() = %v, want %v", got.data, tt.want)
			}
		})
	}
}

// TestMatrixScale verifies Scale multiplies every element by the given scalar.
func TestMatrixScale(t *testing.T) {
	tests := []struct {
		name   string
		rows   int
		cols   int
		data   []float64
		scalar float64
		want   []float64
	}{
		{
			name:   "positive scalar",
			rows:   2, cols: 2,
			data:   []float64{1, 2, 3, 4},
			scalar: 2,
			want:   []float64{2, 4, 6, 8},
		},
		{
			name:   "negative scalar",
			rows:   1, cols: 3,
			data:   []float64{1, 2, 3},
			scalar: -1,
			want:   []float64{-1, -2, -3},
		},
		{
			name:   "zero scalar",
			rows:   2, cols: 2,
			data:   []float64{1, 2, 3, 4},
			scalar: 0,
			want:   []float64{0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := NewMatrix(tt.rows, tt.cols, tt.data)
			if err != nil {
				t.Fatalf("NewMatrix() unexpected error = %v", err)
			}

			got := m.Scale(tt.scalar)
			if !reflect.DeepEqual(got.data, tt.want) {
				t.Errorf("Scale() = %v, want %v", got.data, tt.want)
			}
		})
	}
}

// TestMatrixMatmul verifies matrix multiplication produces the correct
// product and panics when m's columns don't match other's rows.
func TestMatrixMatmul(t *testing.T) {
	tests := []struct {
		name      string
		aRows     int
		aCols     int
		aData     []float64
		bRows     int
		bCols     int
		bData     []float64
		wantRows  int
		wantCols  int
		want      []float64
		wantPanic bool
	}{
		{
			name:     "2x3 times 3x2",
			aRows:    2, aCols: 3, aData: []float64{1, 2, 3, 4, 5, 6},
			bRows:    3, bCols: 2, bData: []float64{7, 8, 9, 10, 11, 12},
			wantRows: 2, wantCols: 2,
			want:     []float64{58, 64, 139, 154},
		},
		{
			name:     "identity matrix",
			aRows:    2, aCols: 2, aData: []float64{1, 2, 3, 4},
			bRows:    2, bCols: 2, bData: []float64{1, 0, 0, 1},
			wantRows: 2, wantCols: 2,
			want:     []float64{1, 2, 3, 4},
		},
		{
			name:     "row vector times column vector",
			aRows:    1, aCols: 3, aData: []float64{1, 2, 3},
			bRows:    3, bCols: 1, bData: []float64{4, 5, 6},
			wantRows: 1, wantCols: 1,
			want:     []float64{32},
		},
		{
			name:     "matrix with zero elements",
			aRows:    2, aCols: 2, aData: []float64{0, 0, 0, 0},
			bRows:    2, bCols: 2, bData: []float64{1, 2, 3, 4},
			wantRows: 2, wantCols: 2,
			want:     []float64{0, 0, 0, 0},
		},
		{
			name:      "mismatched dimensions",
			aRows:     2, aCols: 3, aData: []float64{1, 2, 3, 4, 5, 6},
			bRows:     2, bCols: 2, bData: []float64{1, 2, 3, 4},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, err := NewMatrix(tt.aRows, tt.aCols, tt.aData)
			if err != nil {
				t.Fatalf("NewMatrix() unexpected error = %v", err)
			}
			b, err := NewMatrix(tt.bRows, tt.bCols, tt.bData)
			if err != nil {
				t.Fatalf("NewMatrix() unexpected error = %v", err)
			}

			if tt.wantPanic {
				assertPanics(t, func() { a.Matmul(b) })
				return
			}

			got := a.Matmul(b)
			if got.rows != tt.wantRows || got.cols != tt.wantCols {
				t.Errorf("Matmul() shape = %dx%d, want %dx%d", got.rows, got.cols, tt.wantRows, tt.wantCols)
			}
			if !reflect.DeepEqual(got.data, tt.want) {
				t.Errorf("Matmul() = %v, want %v", got.data, tt.want)
			}
		})
	}
}