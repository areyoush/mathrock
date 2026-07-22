package mathrock

import (
	"fmt"
	"strings"
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

// Add returns a new Matrix that is the element-wise sum of m and the other.
// It panics if the two matrices do not have the same dimensions.
func (m Matrix) Add(other Matrix) Matrix {
	if m.rows != other.rows || m.cols != other.cols {
		panic(fmt.Sprintf("mathrock: matrices must have the same dimensions, got %dx%d and %dx%d", m.rows, m.cols, other.rows, other.cols))
	}

	result := make([]float64, len(m.data))
	for i := range m.data {
		result[i] = m.data[i] + other.data[i]
	}
	return Matrix{data: result, rows: m.rows, cols: m.cols}
}

// Subtract returns a new Matrix that is the element-wise difference of m and the other.
// It panics if the two matrices do not have the same dimensions.
func (m Matrix) Subtract(other Matrix) Matrix {
	if m.rows != other.rows || m.cols != other.cols {
		panic(fmt.Sprintf("mathrock: matrices must have the same dimensions, got %dx%d and %dx%d", m.rows, m.cols, other.rows, other.cols))
	}

	result := make([]float64, len(m.data))
	for i := range m.data {
		result[i] = m.data[i] - other.data[i]
	}
	return Matrix{data: result, rows: m.rows, cols: m.cols}
}

// Scale returns a new Matrix with every element multiplied by scalar.
func (m Matrix) Scale(scalar float64) Matrix {
	result := make([]float64, len(m.data))
	for i := range m.data {
		result[i] = m.data[i] * scalar
	}
	return Matrix{data: result, rows: m.rows, cols: m.cols}
}

// Matmul (Naive) returns a new Matrix that is the product of m and the other (m * other).
// It panics if m's column count does not match other's row count.
func (m Matrix) Matmul(other Matrix) Matrix {
	if m.cols != other.rows {
		panic(fmt.Sprintf("mathrock: cannot multiply %d%d matrix by %d%d matrix, m.cols must equal other.rows", m.rows, m.cols, other.rows, other.cols))
	}

	result := make([]float64, m.rows*other.cols)

	for i := 0; i < m.rows; i++ {
		for k := 0; k < m.cols; k++ {
			mik := m.data[i*m.cols+k]
			if mik == 0 {
				continue
			}
			for j := 0; j < other.cols; j++ {
				result[i*other.cols+j] += mik * other.data[k*other.cols+j]
			}
		}
	}
	
	return Matrix{data: result, rows: m.rows, cols: other.cols}
}

// Transpose returns a new Matrix with rows and columns swapped.
func (m Matrix) Transpose() Matrix {
	result := make([]float64, len(m.data))

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			result[j*m.rows+i] = m.data[i*m.cols+j]
		}
	}
	return Matrix{data: result, rows: m.cols, cols: m.rows}
}

// T is a shorthand alias for Transpose.
func (m Matrix) T() Matrix {
	return m.Transpose()
}

// String returns a human-readable, row-by-row representation of the matrix.
func (m Matrix) String() string {
	var sb strings.Builder
	for r := 0; r < m.rows; r++ {
		sb.WriteString(fmt.Sprint(m.Row(r)))
		sb.WriteString("\n")
	}
	return sb.String()
}