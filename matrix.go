package mathrock

import (
	"fmt"
)

type Matrix struct {
	data	[]float64
	rows	int
	cols	int
}

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

func (m Matrix) At(row, col int) (float64, error) {
	if row < 0 || row >= m.rows {
		return 0, fmt.Errorf("mathrock: row index %d out of bounds for %dx%d matrix", row, m.rows, m.cols)
	}
	if col < 0 || col >= m.cols {
		return 0, fmt.Errorf("mathrock: col index %d out of bounds for %dx%d matrix", col, m.rows, m.cols)
	}

	return m.data[row*m.cols+col], nil	
}

