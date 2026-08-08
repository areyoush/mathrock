<p align="center">
  <img src="assets/logo.jpeg" width="800" alt="mathrock logo">
</p>

<p align="center">
  <a href="https://pkg.go.dev/github.com/areyoush/mathrock"><img src="https://pkg.go.dev/badge/github.com/areyoush/mathrock.svg" alt="Go Reference"></a>
  <a href="https://github.com/areyoush/mathrock/blob/main/LICENSE"><img src="https://img.shields.io/github/license/areyoush/mathrock" alt="License"></a>
</p>

<h1 align="left">mathrock</h1>

<p align="left">A small math library for Go, providing <code>Vector</code> and <code>Matrix</code> types with common linear algebra operations.</p>

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
| **Abs** | `a.Abs()` | never fails |
| **Apply** | `a.Apply(fn)` | never fails |

#### Spatial & Statistical Operations

| Method | Call | Behavior on invalid input |
| :--- | :--- | :--- |
| **Dot** | `a.Dot(b)` | panics on length mismatch |
| **Norm** | `a.Norm()` | never fails |
| **Normalize** | `a.Normalize()` | returns error for the zero vector |
| **Distance** | `a.Distance(b)` | panics on length mismatch |
| **Sum** | `a.Sum()` | never fails |
| **Mean** | `a.Mean()` | returns error for an empty vector |
| **Max** | `a.Max()` | panics on an empty vector |
| **Min** | `a.Min()` | panics on an empty vector |

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
a.Max()                      // 3
a.Min()                      // 1

vector.Vector{-1, 2, -3}.Abs() // [1 2 3]

a.Apply(func(x float64) float64 {
    return x * x
})
// [1 4 9]

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

matrix.Identity(3)
// [1 0 0]
// [0 1 0]
// [0 0 1]

matrix.Zeros(2, 3)
// [0 0 0]
// [0 0 0]

matrix.Ones(2, 2)
// [1 1]
// [1 1]
```
*`NewMatrix` returns an error if `len(data) != rows*cols`. `Identity`, `Zeros`, and `Ones` panic on non-positive dimensions.*

### Methods

#### Initialization & Access

| Method | Call | Behavior on invalid input |
| :--- | :--- | :--- |
| **At** | `m.At(1, 2)` | panics on out-of-bounds index |
| **Set** | `m.Set(0, 0, 99)` | panics on out-of-bounds index |
| **Row** | `m.Row(1)` | panics on out-of-bounds index |
| **Column** | `m.Column(1)` | panics on out-of-bounds index |
| **Diagonal** | `m.Diagonal()` | panics if `m` is not square |

#### Mathematical Operations

| Method | Call | Behavior on invalid input |
| :--- | :--- | :--- |
| **Add** | `m.Add(other)` | panics on dimension mismatch |
| **Subtract** | `m.Subtract(other)` | panics on dimension mismatch |
| **Scale** | `m.Scale(2)` | never fails |
| **Matmul** | `m.Matmul(other)` | panics if `m.cols != other.rows` |
| **Transpose / T** | `m.Transpose()` / `m.T()` | never fails |
| **Trace** | `m.Trace()` | panics if `m` is not square |
| **Apply** | `m.Apply(fn)` | never fails |

#### Utility Methods

| Method | Call | Behavior on invalid input |
| :--- | :--- | :--- |
| **Rows** | `m.Rows()` | never fails |
| **Cols** | `m.Cols()` | never fails |
| **IsSquare** | `m.IsSquare()` | never fails |
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

a.Rows()      // 2
a.Cols()      // 3
a.IsSquare()  // false

square, _ := matrix.NewMatrix(2, 2, []float64{1, 2, 3, 4})
square.Trace()     // 5
square.Diagonal()  // [1 4]

a.Apply(func(x float64) float64 {
    return x * 2
})
// [2 4 6]
// [8 10 12]

a.Equals(a)  // true
```

*Printing a matrix directly (`fmt.Println(m)`) prints it row-by-row rather than as a flat slice, since `Matrix` implements `fmt.Stringer`.*

---

## Status

This is an early, actively developed library (`v0.1.x`). The API may still change as more methods are added. 

## Contribution

Contributions and Suggestions are welcome. Start a discussion regarding your proposed change before raising an issue or a PR. Let's talk.
