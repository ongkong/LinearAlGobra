package matrix

import (
	"testing"
)

func TestMatrixEqual(t *testing.T) {
	if !Equals(MakeZero(6, 7), MakeZero(6, 7)) {
		t.Error("fddffd")
	}
}
func TestMatrixCreation(t *testing.T) {
	var tests = []struct {
		a []int
		b [][]float64
	}{
		{[]int{0, 0}, [][]float64{}},
		{[]int{1, 1}, [][]float64{[]float64{0}}},
		{[]int{3, 3}, [][]float64{[]float64{0, 0, 0}, []float64{0, 0, 0}, []float64{0, 0, 0}}},
		{[]int{5, 3}, [][]float64{[]float64{0, 0, 0}, []float64{0, 0, 0}, []float64{0, 0, 0}, []float64{0, 0, 0}, []float64{0, 0, 0}}},
		{[]int{3, 5}, [][]float64{[]float64{0, 0, 0, 0, 0}, []float64{0, 0, 0, 0, 0}, []float64{0, 0, 0, 0, 0}}},
	}
	for _, test := range tests {
		if z, _ := Make(test.b); !Equals(MakeZero(test.a[0], test.a[1]), z) {
			t.Error("fdffdfd")
		}
	}
}
func TestMatrixMult(t *testing.T) {
	m1, _ := Make([][]float64{[]float64{4, 6, 2}, []float64{5, 2, 5}, []float64{6, 2, 1}})
	m2, _ := Make([][]float64{[]float64{5, 6, 2, 2, 4}, []float64{2, 1, 7, 9, 6}, []float64{97, 23, 21, 1, 8}})
	m3, _ := Make([][]float64{[]float64{226, 76, 92, 64, 68}, []float64{514, 147, 129, 33, 72}, []float64{131, 61, 47, 31, 44}})
	m4, _ := m1.Multiply(m2)
	if !Equals(m3, m4) {
		t.Errorf("%v", m4.data)
	}
}
