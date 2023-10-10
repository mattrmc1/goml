package main

import (
	"errors"
	"fmt"
	formulas "goml/math"
	"goml/math/matrix"
	"goml/validation"
)

// z(l) -> w(l) * a(l-1) + b(l)
// note: activations[1] is weights[0]
// l at index 0 is "really" layer 1 bc layer 0 is the input layer
func CalculateZL(l int) ([]float64, error) {
	p, err := matrix.DotWeightsAndActivations(weights[l], activations[l])
	if err != nil {
		return []float64{}, err
	}
	return matrix.Add1D(p, biases[l])
}

func Feedforward(input []float64) ([]float64, error) {
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

	// a = sigmoid(dot(w, a) + b)
	for i := range weights {
		zl, err := CalculateZL(i)
		if err != nil {
			return []float64{}, err
		}

		a = matrix.Map1D(zl, formulas.Sigmoid)
		activations[i+1] = a
	}

	return a, nil
}
