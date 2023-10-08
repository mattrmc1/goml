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

func Dot(w [][]float64, a []float64) ([]float64, error) {
	var res []float64

	for i, v := range w {
		if !validation.IsEqualDimensions1D(v, a) {
			return []float64{}, errors.New("weight row size must match activation column size")
		}

		sum := float64(0)
		for j := range v {
			sum += w[i][j] * a[j]
		}
		res = append(res, sum)
	}

	return res, nil
}

func Hadamard(a, b [][]float64) ([][]float64, error) {
	if !validation.IsEqualDimensions2D(a, b) {
		return [][]float64{}, errors.New("both matrices must be identical dimensions when calculating the hadamard product")
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
