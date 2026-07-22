# Mathrock

A small math library for Go, providing `Vector` and `Matrix` types with common linear algebra operations.

---

## Installation

```bash
go get github.com/areyoush/mathrock
```

---

## Matrix API Reference

**Import:** `"github.com/areyoush/mathrock/matrix"`

`Matrix` stores its values internally as a flat `[]float64` in row-major order.

### Constructing

```go
m, err := matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
// m:
// [1 2 3]
// [4 5 6]
```
*Note: `NewMatrix` returns an error if `len(data) != rows*cols`.*

### Methods

#### Initialization & Access

| Method | Signature | Behavior on invalid input |
| :--- | :--- | :--- |
| **At** | `func (m Matrix) At(row, col int) float64` | panics on out-of-bounds index |
| **Set** | `func (m Matrix) Set(row, col int, val float64)` | panics on out-of-bounds index |
| **Row** | `func (m Matrix) Row(row int) []float64` | panics on out-of-bounds index |
| **Column** | `func (m Matrix) Column(col int) []float64` | panics on out-of-bounds index |

#### Mathematical Operations

| Method | Signature | Behavior on invalid input |
| :--- | :--- | :--- |
| **Add** | `func (m Matrix) Add(other Matrix) Matrix` | panics on dimension mismatch |
| **Subtract** | `func (m Matrix) Subtract(other Matrix) Matrix` | panics on dimension mismatch |
| **Scale** | `func (m Matrix) Scale(scalar float64) Matrix` | never fails |
| **Matmul** | `func (m Matrix) Matmul(other Matrix) Matrix` | panics if m.cols != other.rows |
| **Transpose / T** | `func (m Matrix) Transpose() Matrix` | never fails |

#### Utility Methods

| Method | Signature | Behavior on invalid input |
| :--- | :--- | :--- |
| **Rows** | `func (m Matrix) Rows() int` | never fails |
| **Cols** | `func (m Matrix) Cols() int` | never fails |
| **Equals** | `func (m Matrix) Equals(other Matrix) bool` | returns false on dimension mismatch |
| **EqualsWithTolerance**| `func (m Matrix) EqualsWithTolerance(other Matrix, tolerance float64) bool` | returns false on dimension mismatch |
| **String** | `func (m Matrix) String() string` | never fails |

### Examples

```go
a, _ := matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
b, _ := matrix.NewMatrix(3, 2, []float64{7, 8, 9, 10, 11, 12})

a.At(1, 2)           // 6
a.Set(0, 0, 99)      // mutates a in place
a.Row(1)             // [4 5 6]
a.Column(1)          // [2 5]

a.Scale(2)
// [2 4 6]
// [8 10 12]

a.Matmul(b)
// [58 64]
// [139 154]

a.Transpose()
// [1 4]
// [2 5]
// [3 6]

a.Rows()     // 2
a.Cols()     // 3

a.Equals(a)  // true
```

*Printing a matrix directly (`fmt.Println(m)`) prints it row-by-row rather than as a flat slice, since `Matrix` implements `fmt.Stringer`.*

---

## Vector API Reference

**Import:** `"github.com/areyoush/mathrock/vector"`

`Vector` is defined as type `Vector []float64`.

### Constructing

```go
a := vector.Vector{1, 2, 3}
b := vector.Vector{4, 5, 6}
```

### Methods

#### Arithmetic & Scaling

| Method | Signature | Behavior on invalid input |
| :--- | :--- | :--- |
| **Add** | `func (v Vector) Add(other Vector) Vector` | panics on length mismatch |
| **Subtract** | `func (v Vector) Subtract(other Vector) Vector` | panics on length mismatch |
| **Multiply** | `func (v Vector) Multiply(other Vector) Vector` | panics on length mismatch |
| **Scale** | `func (v Vector) Scale(scalar float64) Vector` | never fails |

#### Spatial & Statistical Operations

| Method | Signature | Behavior on invalid input |
| :--- | :--- | :--- |
| **Dot** | `func (v Vector) Dot(other Vector) float64` | panics on length mismatch |
| **Norm** | `func (v Vector) Norm() float64` | never fails |
| **Normalize** | `func (v Vector) Normalize() (Vector, error)` | returns error for the zero vector |
| **Distance** | `func (v Vector) Distance(other Vector) float64` | panics on length mismatch |
| **Sum** | `func (v Vector) Sum() float64` | never fails |
| **Mean** | `func (v Vector) Mean() (float64, error)` | returns error for an empty vector |

#### Utility Methods

| Method | Signature | Behavior on invalid input |
| :--- | :--- | :--- |
| **Equals** | `func (v Vector) Equals(other Vector) bool` | returns false on length mismatch |
| **EqualsWithTolerance**| `func (v Vector) EqualsWithTolerance(other Vector, tolerance float64) bool` | returns false on length mismatch |

### Examples

```go
a := vector.Vector{1, 2, 3}
b := vector.Vector{4, 5, 6}

a.Dot(b)                     // 32
a.Add(b)                     // [5 7 9]
a.Subtract(b)                // [-3 -3 -3]
a.Scale(2)                   // [2 4 6]
a.Multiply(b)                // [4 10 18]
a.Norm()                     // 3.7416573867739413
a.Distance(b)                // 5.196152422706632

normalized, err := a.Normalize()
if err != nil {
    // v was the zero vector
}

mean, err := a.Mean()
if err != nil {
    // v was empty
}

a.Equals(a)                  // true
a.EqualsWithTolerance(b, 10) // true (loose tolerance)
```

---

## Status

This is an early, actively developed library (v0.1.x). The API may still change as more methods are added. Contributions and suggestions are welcome.