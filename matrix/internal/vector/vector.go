// Package vector defines the Vector struct and creation.
package vector

// Vector type that is used to build a Matrix.
// The default initialization is an empty vector.
type Vector []float64

// MakeZero creates a zero vector with given length.
func MakeZero(numrow int) Vector {
	return make(Vector, numrow)
}

