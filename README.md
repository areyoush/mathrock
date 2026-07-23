# mathrock

A small math library for Go, providing `Vector` and `Matrix` types with common linear algebra operations.

---

## Installation

```bash
go get github.com/areyoush/mathrock
```

---

## Vector

**Import:** `"github.com/areyoush/mathrock/vector"`

`Vector` is defined as `type Vector []float64`.

### Constructing

```go
a := vector.Vector{1, 2, 3}
b := vector.Vector{4, 5, 6}
```

### Methods

#### Arithmetic & Scaling

| Method | Call | Behavior on invalid input |
| :--- | :--- | :--- |
| **Add** | `a.Add(b)` | panics on length mismatch |
| **Subtract** | `a.Subtract(b)` | panics on length mismatch |
| **Multiply** | `a.Multiply(b)` | panics on length mismatch |
| **Scale** | `a.Scale(2)` | never fails |

#### Spatial & Statistical Operations

| Method | Call | Behavior on invalid input |
| :--- | :--- | :--- |
| **Dot** | `a.Dot(b)` | panics on length mismatch |
| **Norm** | `a.Norm()` | never fails |
| **Normalize** | `a.Normalize()` | returns error for the zero vector |
| **Distance** | `a.Distance(b)` | panics on length mismatch |
| **Sum** | `a.Sum()` | never fails |
| **Mean** | `a.Mean()` | returns error for an empty vector |

#### Utility Methods

| Method | Call | Behavior on invalid input |
| :--- | :--- | :--- |
| **Equals** | `a.Equals(b)` | returns `false` on length mismatch |
| **EqualsWithTolerance** | `a.EqualsWithTolerance(b, 1e-6)` | returns `false` on length mismatch |

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

## Matrix

**Import:** `"github.com/areyoush/mathrock/matrix"`

`Matrix` stores its values internally as a flat `[]float64` in row-major order.

### Constructing

```go
m, err := matrix.NewMatrix(2, 3, []float64{1, 2, 3, 4, 5, 6})
// m:
// [1 2 3]
// [4 5 6]
```
*`NewMatrix` returns an error if `len(data) != rows*cols`.*

### Methods

#### Initialization & Access

| Method | Call | Behavior on invalid input |
| :--- | :--- | :--- |
| **At** | `m.At(1, 2)` | panics on out-of-bounds index |
| **Set** | `m.Set(0, 0, 99)` | panics on out-of-bounds index |
| **Row** | `m.Row(1)` | panics on out-of-bounds index |
| **Column** | `m.Column(1)` | panics on out-of-bounds index |

#### Mathematical Operations

| Method | Call | Behavior on invalid input |
| :--- | :--- | :--- |
| **Add** | `m.Add(other)` | panics on dimension mismatch |
| **Subtract** | `m.Subtract(other)` | panics on dimension mismatch |
| **Scale** | `m.Scale(2)` | never fails |
| **Matmul** | `m.Matmul(other)` | panics if `m.cols != other.rows` |
| **Transpose / T** | `m.Transpose()` / `m.T()` | never fails |

#### Utility Methods

| Method | Call | Behavior on invalid input |
| :--- | :--- | :--- |
| **Rows** | `m.Rows()` | never fails |
| **Cols** | `m.Cols()` | never fails |
| **Equals** | `m.Equals(other)` | returns `false` on dimension mismatch |
| **EqualsWithTolerance** | `m.EqualsWithTolerance(other, 1e-6)` | returns `false` on dimension mismatch |
| **String** | `fmt.Println(m)` | never fails |

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

## Status

This is an early, actively developed library (`v0.1.x`). The API may still change as more methods are added. Contributions and suggestions are welcome.
