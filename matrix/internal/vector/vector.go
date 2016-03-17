// Package vector defines the Vector struct and creation.
package vector

// Vector type that is used to build a Matrix.
// The default initialization is an empty vector.
type Vector []float64

// MakeZero creates a zero vector with given length.
func MakeZero(numrow int) Vector {
	return make(Vector, numrow)
}

// MemSet fills a vector with the given value.
// It mimics the behavior of std::memset.
// Credits to icza of stackoverflow.
func MemSet(v []Vector, value float64){
	if len(v) == 0 {
		return
	}
	v[0] = value
	for i := 1; i < len(v); i *= 2{
		copy(v[i:],v[:i])
	}
}