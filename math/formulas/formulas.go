package formulas

import (
	"errors"
	"math"
)

// S = 1 / ( 1 + e^-x )
func Sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

// S' = S(x) * ( 1 - S(x) )
func DeltaSigmoid(x float64) float64 {
	return Sigmoid(x) * (1 - Sigmoid(x))
}

// C = (a(L) - y)^2
func Cost(a, y []float64) ([]float64, error) {
	if len(a) != len(y) {
		return []float64{}, errors.New("output array dimensions does not match expected array dimensions")
	}

	var c = make([]float64, len(a))

	for i := range a {
		c[i] = math.Pow(a[i]-y[i], 2)
	}

	return c, nil
}

// C' = 2(a(L) - y)
func DeltaCost(a, y []float64) ([]float64, error) {
	if len(a) != len(y) {
		return []float64{}, errors.New("output array dimensions does not match expected array dimensions")
	}

	var dc = make([]float64, len(a))

	for i := range a {
		dc[i] = 2 * (a[i] - y[i])
	}

	return dc, nil
}
