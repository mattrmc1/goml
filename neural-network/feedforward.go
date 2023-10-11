package neuralnetwork

import (
	"errors"
	"fmt"
	"goml/math/formulas"
	"goml/math/matrix"
	"goml/validation"
)

// a(l) -> squish(w(l) * a(l-1) + b(l))
func feedforward(input []float64) ([]float64, error) {
	if len(layers) == 0 {
		return []float64{}, errors.New("network not initialized correctly")
	}

	if len(input) != layers[0] {
		return []float64{}, fmt.Errorf("invalid input - expected size: %v", layers[0])
	}

	if !validation.Validate1D(input, func(v float64) bool {
		return v >= 0 && v <= 1
	}) {
		return []float64{}, errors.New("all input values must be a float between 0 and 1")
	}

	a := make([]float64, len(input))
	copy(a, input)

	activations[0] = make([]float64, len(input))
	copy(activations[0], input)

	for i := range weights {
		zl, err := z(i)
		if err != nil {
			return []float64{}, err
		}

		a = matrix.Map1D(zl, formulas.Sigmoid)
		activations[i+1] = a
	}

	return a, nil
}
