package matrix

import (
	"errors"
	"goml/validation"
)

func Create2D(rows, cols int) [][]float64 {
	res := make([][]float64, rows)
	for i := range res {
		res[i] = make([]float64, cols)
	}
	return res
}

// r(j) • l(k) -> w(j,k)
func DotToCreateWeights(right, left []float64) [][]float64 {
	w := Create2D(len(right), len(left))

	for j, a := range right {
		for k, b := range left {
			w[j][k] = a * b
		}
	}

	return w
}

// w(j,k) • a(k) -> z(j)
func DotWeightsAndActivations(w [][]float64, a []float64) ([]float64, error) {
	var res = make([]float64, len(w))

	for i := range w {
		if !validation.IsEqualDimensions1D(w[i], a) {
			return []float64{}, errors.New("weight column size must match activation row size")
		}

		sum := float64(0)
		for j := range w[i] {
			sum += w[i][j] * a[j]
		}
		res[i] = sum
	}

	return res, nil
}

func Transpose(m [][]float64) [][]float64 {
	res := make([][]float64, len(m[0]))
	for i := range res {
		res[i] = make([]float64, len(m))
		for j := range res[i] {
			res[i][j] = m[j][i]
		}
	}

	return res
}

func Hadamard1D(a, b []float64) ([]float64, error) {
	if !validation.IsEqualDimensions1D(a, b) {
		return []float64{}, errors.New("matrices must have identical dimensions when calculating the hadamard product")
	}

	res := make([]float64, len(a))
	for i := range a {
		res[i] = a[i] * b[i]
	}

	return res, nil
}

func Hadamard2D(a, b [][]float64) ([][]float64, error) {
	if !validation.IsEqualDimensions2D(a, b) {
		return [][]float64{}, errors.New("matrices must have identical dimensions when calculating the hadamard product")
	}

	res := Create2D(len(a), len(a[0]))
	for i := range a {
		for j := range a[i] {
			res[i][j] = a[i][j] * b[i][j]
		}
	}

	return res, nil
}

func Add1D(a, b []float64) ([]float64, error) {
	if !validation.IsEqualDimensions1D(a, b) {
		return []float64{}, errors.New("cannot add 2 arrays of different dimensions")
	}

	var res []float64

	for i := range a {
		res = append(res, a[i]+b[i])
	}

	return res, nil
}

type mapper func(float64) float64

func Map1D(m []float64, fn mapper) []float64 {
	var res []float64

	for _, v := range m {
		res = append(res, fn(v))
	}

	return res
}
