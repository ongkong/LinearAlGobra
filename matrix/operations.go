package matrix

import (
	"errors"
	"fmt"
)

// Scale creates a new matrix that is the result of scaling the reciever matrix by a float64.
func (a matrix) Scale(c float64) matrix {
	result := MakeZero(a.lenRow, a.lenCol)
	for row := 0; row < a.lenRow; row++ {
		for col := 0; col < a.lenCol; col++ {
			result.data[row][col] = c * a.data[row][col]
		}
	}
	return result
}

// Multiply performs matrix multiplication and returns the resulting matrix and an error if one exists.
// Order matters, meaning which matrix is the receiever and which is the argument affects the result, or error.
// If either matrix is an empty matrix, the result will be an empty matrix as well.
func (a matrix) Multiply(b matrix) (matrix, error) {
	if a.lenRow == 0 || a.lenCol == 0 || b.lenRow == 0 || b.lenCol == 0 {
		// if either arguments are empty matrices.
		return MakeZero(0, 0), nil
	}
	if a.lenCol != b.lenRow {
		// matrices must be dimensionally compatible
		return MakeZero(0, 0), errors.New(
			fmt.Sprintf("Two Matrices are incompatible: %d x %d and %d x %d multiplication",
				a.lenRow, a.lenCol, b.lenRow, b.lenCol))
	}
	result := MakeZero(a.lenRow, b.lenCol) // creates result matrix with proper dimensions
	// matrix multiplcation algorithm
	for row := 0; row < result.lenRow; row++ {
		for col := 0; col < result.lenCol; col++ {
			for i := 0; i < a.lenCol; i++ {
				result.data[row][col] += a.data[row][i] * b.data[i][col] // similar to dot product
			}
		}
	}
	return result, nil
}

// Transpose returns a transposed receiver matrix.
func (a matrix) Transpose() matrix {
	result := MakeZero(a.lenCol,a.lenRow)
	for row := 0; row < result.lenRow; row++{
		for col := 0; col < result.lenCol; col++{
			result.data[row][col] = a.data[col][row]
		}
	}
	return result
}

func (a matrix) RowEchelon() matrix {
	
}
