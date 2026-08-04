package matrix

import (
	"reflect"
	"testing"

	"github.com/areyoush/mathrock/vector"
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

// TestMatrixColumn verifies Column returns the correct slice of values for a given column and panics when col is out of bounds.
func TestMatrixColumn(t *testing.T) {
	tests := []struct {
		name      string
		rows      int
		cols      int
		data      []float64
		col       int
		want      []float64
		wantPanic bool
	}{
		{
			name: "first column",
			rows: 2, cols: 3,
			data: []float64{1, 2, 3, 4, 5, 6},
			col:  0,
			want: []float64{1, 4},
		},
		{
			name: "last column",
			rows: 2, cols: 3,
			data: []float64{1, 2, 3, 4, 5, 6},
			col:  2,
			want: []float64{3, 6},
		},
		{
			name: "single row matrix",
			rows: 1, cols: 3,
			data: []float64{7, 8, 9},
			col:  1,
			want: []float64{8},
		},
		{
			name:      "col out of bounds negative",
			rows:      2, cols: 3,
			data:      []float64{1, 2, 3, 4, 5, 6},
			col:       -1,
			wantPanic: true,
		},
		{
			name:      "col out of bounds too large",
			rows:      2, cols: 3,
			data:      []float64{1, 2, 3, 4, 5, 6},
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
				assertPanics(t, func() { m.Column(tt.col) })
				return
			}

			got := m.Column(tt.col)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Column(%d) = %v, want %v", tt.col, got, tt.want)
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

// TestMatrixMatmul verifies matrix multiplication produces the correct product and panics when m's columns don't match other's rows.
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

// TestMatrixTranspose verifies Transpose correctly swaps rows and columns.
func TestMatrixTranspose(t *testing.T) {
	tests := []struct {
		name     string
		rows     int
		cols     int
		data     []float64
		wantRows int
		wantCols int
		want     []float64
	}{
		{
			name:     "2x3 matrix",
			rows:     2, cols: 3,
			data:     []float64{1, 2, 3, 4, 5, 6},
			wantRows: 3, wantCols: 2,
			want:     []float64{1, 4, 2, 5, 3, 6},
		},
		{
			name:     "square matrix",
			rows:     2, cols: 2,
			data:     []float64{1, 2, 3, 4},
			wantRows: 2, wantCols: 2,
			want:     []float64{1, 3, 2, 4},
		},
		{
			name:     "single row",
			rows:     1, cols: 3,
			data:     []float64{1, 2, 3},
			wantRows: 3, wantCols: 1,
			want:     []float64{1, 2, 3},
		},
		{
			name:     "single column",
			rows:     3, cols: 1,
			data:     []float64{1, 2, 3},
			wantRows: 1, wantCols: 3,
			want:     []float64{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := NewMatrix(tt.rows, tt.cols, tt.data)
			if err != nil {
				t.Fatalf("NewMatrix() unexpected error = %v", err)
			}

			got := m.Transpose()
			if got.rows != tt.wantRows || got.cols != tt.wantCols {
				t.Errorf("Transpose() shape = %dx%d, want %dx%d", got.rows, got.cols, tt.wantRows, tt.wantCols)
			}
			if !reflect.DeepEqual(got.data, tt.want) {
				t.Errorf("Transpose() = %v, want %v", got.data, tt.want)
			}
		})
	}
}

// TestMatrixT verifies T is a working alias for Transpose.
func TestMatrixT(t *testing.T) {
	m, err := NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
	if err != nil {
		t.Fatalf("NewMatrix() unexpected error = %v", err)
	}

	got := m.T()
	want := m.Transpose()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("T() = %v, want %v", got, want)
	}
}

// TestMatrixEqualsWithTolerance verifies equality checks at, within, and outside a given tolerance, including mismatched dimensions.
func TestMatrixEqualsWithTolerance(t *testing.T) {
	tests := []struct {
		name      string
		aRows     int
		aCols     int
		aData     []float64
		bRows     int
		bCols     int
		bData     []float64
		tolerance float64
		want      bool
	}{
		{
			name:      "identical matrices",
			aRows:     2, aCols: 2, aData: []float64{1, 2, 3, 4},
			bRows:     2, bCols: 2, bData: []float64{1, 2, 3, 4},
			tolerance: 1e-9,
			want:      true,
		},
		{
			name:      "different values",
			aRows:     2, aCols: 2, aData: []float64{1, 2, 3, 4},
			bRows:     2, bCols: 2, bData: []float64{1, 2, 3, 5},
			tolerance: 1e-9,
			want:      false,
		},
		{
			name:      "within tolerance",
			aRows:     1, aCols: 3, aData: []float64{1, 2, 3},
			bRows:     1, bCols: 3, bData: []float64{1.0001, 2, 3},
			tolerance: 0.001,
			want:      true,
		},
		{
			name:      "outside tolerance",
			aRows:     1, aCols: 3, aData: []float64{1, 2, 3},
			bRows:     1, bCols: 3, bData: []float64{1.1, 2, 3},
			tolerance: 0.001,
			want:      false,
		},
		{
			name:      "mismatched rows",
			aRows:     2, aCols: 2, aData: []float64{1, 2, 3, 4},
			bRows:     4, bCols: 1, bData: []float64{1, 2, 3, 4},
			tolerance: 1e-9,
			want:      false,
		},
		{
			name:      "mismatched cols",
			aRows:     2, aCols: 2, aData: []float64{1, 2, 3, 4},
			bRows:     2, bCols: 3, bData: []float64{1, 2, 3, 4, 5, 6},
			tolerance: 1e-9,
			want:      false,
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

			got := a.EqualsWithTolerance(b, tt.tolerance)
			if got != tt.want {
				t.Errorf("EqualsWithTolerance() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestMatrixEquals verifies Equals using the default tolerance, including floating-point rounding noise and mismatched dimensions.
func TestMatrixEquals(t *testing.T) {
	tests := []struct {
		name  string
		aRows int
		aCols int
		aData []float64
		bRows int
		bCols int
		bData []float64
		want  bool
	}{
		{
			name:  "identical matrices",
			aRows: 2, aCols: 2, aData: []float64{1, 2, 3, 4},
			bRows: 2, bCols: 2, bData: []float64{1, 2, 3, 4},
			want:  true,
		},
		{
			name:  "different values",
			aRows: 2, aCols: 2, aData: []float64{1, 2, 3, 4},
			bRows: 2, bCols: 2, bData: []float64{1, 2, 3, 5},
			want:  false,
		},
		{
			name:  "floating point noise",
			aRows: 1, aCols: 2, aData: []float64{0.6, 0.8},
			bRows: 1, bCols: 2, bData: []float64{0.6000000000000001, 0.8},
			want:  true,
		},
		{
			name:  "mismatched dimensions",
			aRows: 2, aCols: 2, aData: []float64{1, 2, 3, 4},
			bRows: 4, bCols: 1, bData: []float64{1, 2, 3, 4},
			want:  false,
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

			got := a.Equals(b)
			if got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestMatrixRows verifies Rows returns the correct row count.
func TestMatrixRows(t *testing.T) {
	tests := []struct {
		name string
		rows int
		cols int
		data []float64
		want int
	}{
		{
			name: "2x3 matrix",
			rows: 2, cols: 3,
			data: []float64{1, 2, 3, 4, 5, 6},
			want: 2,
		},
		{
			name: "single row",
			rows: 1, cols: 4,
			data: []float64{1, 2, 3, 4},
			want: 1,
		},
		{
			name: "square matrix",
			rows: 3, cols: 3,
			data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := NewMatrix(tt.rows, tt.cols, tt.data)
			if err != nil {
				t.Fatalf("NewMatrix() unexpected error = %v", err)
			}

			got := m.Rows()
			if got != tt.want {
				t.Errorf("Rows() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestMatrixCols verifies Cols returns the correct column count.
func TestMatrixCols(t *testing.T) {
	tests := []struct {
		name string
		rows int
		cols int
		data []float64
		want int
	}{
		{
			name: "2x3 matrix",
			rows: 2, cols: 3,
			data: []float64{1, 2, 3, 4, 5, 6},
			want: 3,
		},
		{
			name: "single column",
			rows: 4, cols: 1,
			data: []float64{1, 2, 3, 4},
			want: 1,
		},
		{
			name: "square matrix",
			rows: 3, cols: 3,
			data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := NewMatrix(tt.rows, tt.cols, tt.data)
			if err != nil {
				t.Fatalf("NewMatrix() unexpected error = %v", err)
			}

			got := m.Cols()
			if got != tt.want {
				t.Errorf("Cols() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestIdentity verifies Identity produces a matrix with 1s on the diagonal and 0s elsewhere, and panics for non-positive n.
func TestIdentity(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		want      []float64
		wantPanic bool
	}{
		{
			name: "3x3 identity",
			n:    3,
			want: []float64{
				1, 0, 0,
				0, 1, 0,
				0, 0, 1,
			},
		},
		{
			name: "1x1 identity",
			n:    1,
			want: []float64{1},
		},
		{
			name: "2x2 identity",
			n:    2,
			want: []float64{1, 0, 0, 1},
		},
		{
			name:      "zero size panics",
			n:         0,
			wantPanic: true,
		},
		{
			name:      "negative size panics",
			n:         -2,
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assertPanics(t, func() { Identity(tt.n) })
				return
			}

			got := Identity(tt.n)
			if got.rows != tt.n || got.cols != tt.n {
				t.Errorf("Identity(%d) shape = %dx%d, want %dx%d", tt.n, got.rows, got.cols, tt.n, tt.n)
			}
			if !reflect.DeepEqual(got.data, tt.want) {
				t.Errorf("Identity(%d) = %v, want %v", tt.n, got.data, tt.want)
			}
		})
	}
}

// TestZeros verifies Zeros produces a matrix of the correct shape filled with zeros, and panics for non-positive dimensions.
func TestZeros(t *testing.T) {
	tests := []struct {
		name      string
		rows      int
		cols      int
		wantPanic bool
	}{
		{
			name: "2x3 zeros",
			rows: 2,
			cols: 3,
		},
		{
			name: "1x1 zeros",
			rows: 1,
			cols: 1,
		},
		{
			name: "square zeros",
			rows: 3,
			cols: 3,
		},
		{
			name:      "zero rows panics",
			rows:      0,
			cols:      3,
			wantPanic: true,
		},
		{
			name:      "zero cols panics",
			rows:      3,
			cols:      0,
			wantPanic: true,
		},
		{
			name:      "negative rows panics",
			rows:      -2,
			cols:      3,
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assertPanics(t, func() { Zeros(tt.rows, tt.cols) })
				return
			}

			got := Zeros(tt.rows, tt.cols)
			if got.rows != tt.rows || got.cols != tt.cols {
				t.Errorf("Zeros(%d, %d) shape = %dx%d, want %dx%d", tt.rows, tt.cols, got.rows, got.cols, tt.rows, tt.cols)
			}
			for i, v := range got.data {
				if v != 0 {
					t.Errorf("Zeros(%d, %d) data[%d] = %v, want 0", tt.rows, tt.cols, i, v)
				}
			}
		})
	}
}

// TestOnes verifies Ones produces a matrix of the correct shape filled with ones, and panics for non-positive dimensions.
func TestOnes(t *testing.T) {
	tests := []struct {
		name      string
		rows      int
		cols      int
		wantPanic bool
	}{
		{
			name: "2x3 ones",
			rows: 2,
			cols: 3,
		},
		{
			name: "1x1 ones",
			rows: 1,
			cols: 1,
		},
		{
			name: "square ones",
			rows: 3,
			cols: 3,
		},
		{
			name:      "zero rows panics",
			rows:      0,
			cols:      3,
			wantPanic: true,
		},
		{
			name:      "zero cols panics",
			rows:      3,
			cols:      0,
			wantPanic: true,
		},
		{
			name:      "negative cols panics",
			rows:      3,
			cols:      -1,
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assertPanics(t, func() { Ones(tt.rows, tt.cols) })
				return
			}

			got := Ones(tt.rows, tt.cols)
			if got.rows != tt.rows || got.cols != tt.cols {
				t.Errorf("Ones(%d, %d) shape = %dx%d, want %dx%d", tt.rows, tt.cols, got.rows, got.cols, tt.rows, tt.cols)
			}
			for i, v := range got.data {
				if v != 1 {
					t.Errorf("Ones(%d, %d) data[%d] = %v, want 1", tt.rows, tt.cols, i, v)
				}
			}
		})
	}
}

// TestMatrixTrace verifies Trace returns the sum of diagonal elements and panics on a non-square matrix.
func TestMatrixTrace(t *testing.T) {
	tests := []struct {
		name      string
		rows      int
		cols      int
		data      []float64
		want      float64
		wantPanic bool
	}{
		{
			name: "3x3 matrix",
			rows: 3, cols: 3,
			data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 15, // 1 + 5 + 9
		},
		{
			name: "2x2 matrix",
			rows: 2, cols: 2,
			data: []float64{1, 2, 3, 4},
			want: 5, // 1 + 4
		},
		{
			name: "1x1 matrix",
			rows: 1, cols: 1,
			data: []float64{7},
			want: 7,
		},
		{
			name: "includes negatives",
			rows: 2, cols: 2,
			data: []float64{-1, 2, 3, -4},
			want: -5,
		},
		{
			name:      "non-square matrix panics",
			rows:      2, cols: 3,
			data:      []float64{1, 2, 3, 4, 5, 6},
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
				assertPanics(t, func() { m.Trace() })
				return
			}

			got := m.Trace()
			if got != tt.want {
				t.Errorf("Trace() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestMatrixDiagonal verifies Diagonal returns the correct diagonal values as a Vector and panics on a non-square matrix.
func TestMatrixDiagonal(t *testing.T) {
	tests := []struct {
		name      string
		rows      int
		cols      int
		data      []float64
		want      vector.Vector
		wantPanic bool
	}{
		{
			name: "3x3 matrix",
			rows: 3, cols: 3,
			data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: vector.Vector{1, 5, 9},
		},
		{
			name: "2x2 matrix",
			rows: 2, cols: 2,
			data: []float64{1, 2, 3, 4},
			want: vector.Vector{1, 4},
		},
		{
			name: "1x1 matrix",
			rows: 1, cols: 1,
			data: []float64{7},
			want: vector.Vector{7},
		},
		{
			name: "includes negatives",
			rows: 2, cols: 2,
			data: []float64{-1, 2, 3, -4},
			want: vector.Vector{-1, -4},
		},
		{
			name:      "non-square matrix panics",
			rows:      2, cols: 3,
			data:      []float64{1, 2, 3, 4, 5, 6},
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
				assertPanics(t, func() { m.Diagonal() })
				return
			}

			got := m.Diagonal()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Diagonal() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestMatrixIsSquare verifies IsSquare correctly identifies square and non-square matrices.
func TestMatrixIsSquare(t *testing.T) {
	tests := []struct {
		name string
		rows int
		cols int
		data []float64
		want bool
	}{
		{
			name: "square matrix",
			rows: 3, cols: 3,
			data: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: true,
		},
		{
			name: "1x1 matrix",
			rows: 1, cols: 1,
			data: []float64{5},
			want: true,
		},
		{
			name: "more rows than cols",
			rows: 3, cols: 2,
			data: []float64{1, 2, 3, 4, 5, 6},
			want: false,
		},
		{
			name: "more cols than rows",
			rows: 2, cols: 3,
			data: []float64{1, 2, 3, 4, 5, 6},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := NewMatrix(tt.rows, tt.cols, tt.data)
			if err != nil {
				t.Fatalf("NewMatrix() unexpected error = %v", err)
			}

			got := m.IsSquare()
			if got != tt.want {
				t.Errorf("IsSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}