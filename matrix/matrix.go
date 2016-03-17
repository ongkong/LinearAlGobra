// Package matrix defines the matrix struct, creation of matrices, and operations.
// Throughout the documentation of this package, the word matrix will be used to mean either matrices or vectors depending on the situation.
// A vector is a matrix with 1 as the lenRow or lenCol attribute depending on either its a column or row vector.
package matrix

import (
	"errors"
	"github.com/ongkong/internal/vector"
	"math"
)

// Matrix type that is used for operations.
// A matrix is represented as a collection of row vectors.
// A matrix can also respresent a row or column vector.
// For example:
// Matrix{1, 4, []vector.Vector{vector.Vector{4,5,5,6}} is the row vector [4,5,5,6]
// Matrix{3, 1, []vector.Vector{vector.Vector{4},vector.Vector{5},vector.Vector{6}} is column vector [4,5,6]
type matrix struct {
	lenRow int
	lenCol int
	data   []vector.Vector
}

// MakeZero creates a zero Matrix with given dimension.
func MakeZero(lenrow, lencol int) matrix {
	m := matrix{lenRow: lenrow, lenCol: lencol}
	m.data = make([]vector.Vector, lenrow)
	for i := range m.data {
		m.data[i] = vector.MakeZero(lencol)
	}
	return m
}

// MakeFilled creates a matrix of the given dimension filled with the given value.
func MakeFilled(lenrow, lencol int, value float64) matrix {
	result := MakeZero(lenrow, lencol)
	for row := 0; row < result.lenRow; row++ {
		vector.Memset(result.data[row], value)
	}
}

// Make creates a Matrix with the given 2x2 float64 slice
// All the slices must be of the same length.
func Make(args [][]float64) (matrix, error) {
	numRow := len(args)
	if numRow == 0 {
		return MakeZero(0, 0), nil // if no slices are provided, creates an an empty matrix
	}
	numCol := len(args[0]) // every slice has to be this length
	for i := range args {
		if len(args[i]) != numCol { // check if slices have equal length
			return MakeZero(0, 0), errors.New("All rows of the matrix must have the same number of columns")
		}
	}
	m := MakeZero(numRow, numCol)       // create zero matrix with proper dimensions
	for row := 0; row < numRow; row++ { // populate the zero matrix with elements of slices
		for col := 0; col < numCol; col++ {
			m.data[row][col] = args[row][col]
		}
	}
	return m, nil
}

// MakeIdentity creates an identity matrix given the size.
// It will be a square matrix.
func MakeIdentity(size int) matrix {
	m := MakeZero(size, size)
	for i := range m.data {
		m.data[i][i] = 1
	}
	return m
}

// Equals checks if the two matrices are equal, returning true or false accordingly.
// The absolute value of the difference between two corresponding elements have to be less than a set error.
// The dimensions must match up.
func Equals(a matrix, b matrix) bool {
	const e float64 = 1E-6 // allowable error
	if a.lenRow == 0 && b.lenRow == 0 {
		// two empty matrices are equal
		return true
	}
	if a.lenRow == b.lenRow && a.lenCol == b.lenCol {
		// checks dimension equality
		for row := 0; row < a.lenRow; row++ {
			for col := 0; col < a.lenCol; col++ {
				// checks that each element are equal within an error
				if math.Abs(a.data[row][col]-b.data[row][col]) > e {
					return false
				}
			}
		}
		return true
	}
	return false
}
