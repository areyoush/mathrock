package mathrock

import (
	"fmt"
)

// Matrix represents a 2D matrix of float64 values, stored internally as a flat slice in row-major order.
type Matrix struct {
	data []float64
	rows int
	cols int
}

// NewMatrix creates a new Matrix with a given number of rows and columns populated with data in row-major order.
// It returns an error if len(data) does not equal rows*cols.
func NewMatrix(rows, cols int, data []float64) (Matrix, error) {
	if len(data) != rows*cols {
		return Matrix{}, fmt.Errorf("mathrock: data length must equal rows*cols, got %d elements for a %dx%d matrix", len(data), rows, cols)
	}
	return Matrix{
		data: data,
		rows: rows,
		cols: cols,
	}, nil
}

// At returns the value at the given row and column.
// It panics if row or col is out of bounds for the matrix's dimensions.
func (m Matrix) At(row, col int) float64 {
	if row < 0 || row >= m.rows {
		panic(fmt.Sprintf("mathrock: row index %d out of bounds for %dx%d matrix", row, m.rows, m.cols))
	}
	if col < 0 || col >= m.cols {
		panic(fmt.Sprintf("mathrock: col index %d out of bounds for %dx%d matrix", col, m.rows, m.cols))
	}
	return m.data[row*m.cols+col]
}

// Set assigns val to the position at the given row and column.
// It panics if row or col is out of bounds for the matrix's dimensions.
func (m Matrix) Set(row, col int, val float64) {
	if row < 0 || row >= m.rows {
		panic(fmt.Sprintf("mathrock: row index %d out of bounds for %dx%d matrix", row, m.rows, m.cols))
	}
	if col < 0 || col >= m.cols {
		panic(fmt.Sprintf("mathrock: col index %d out of bounds for %dx%d matrix", col, m.rows, m.cols))
	}
	m.data[row*m.cols+col] = val	
}

// Row returns the values in the given row as a slice.
// It panics if row is out of bounds for matrix's dimenstions.
func (m Matrix) Row(row int) []float64 {
	if row < 0 || row >= m.rows {
		panic(fmt. Sprintf("mathrock: row index %d out of bounds for %dx%d matrix", row, m.rows, m.cols))
	}
	return m.data[row*m.cols : row*m.cols+m.cols]
}