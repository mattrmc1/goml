package matrix

import (
	"errors"
	"goml/validation"
)

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
