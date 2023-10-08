package main

import "goml/math/matrix"

func z(l int) ([]float64, error) {
	p, err := matrix.Dot(weights[l], activations[l-1])
	if err != nil {
		return []float64{}, err
	}
	return matrix.Add1D(p, biases[l])
}

func dCdA(l int, tOutput []float64) []float64 {
	// dC/dA:
	// 		if last layer, compare to expected output layer
	//				-> deltaCost fn
	//		else, recursively calculate to last layer
	//				-> w(l+1)^T * deltaSig(z(l+1)) * dC/dA(l+1)

	return []float64{}
}

func calculateDeltasAtLayer(layer int) ([][]float64, []float64, error) {
	// calculate z(l) -> w(l) * a(l-1) + b(l)

	// dC/dB -> deltaSig(z(l)) • dC/dA
	// dC/dW -> a(l-1) • deltaSig(z(l)) • dC/dA

	var deltaWeights [][]float64
	var deltaBaises []float64

	return deltaWeights, deltaBaises, nil
}

// Returns all deltaWeights, deltaBaises
func Backpropagate(tInput, tOutput []float64) ([][][]float64, [][]float64, error) {
	// Validate initialization
	//		-> learning rate should be defined
	//		-> layers[0] dimensions should match tInput
	//		-> layers[len-1] dimensions should match tOutput
	//		-> weights and biases should be initialized with dummy data

	// Feedforward and store data:
	// 		-> activations (note: output is activations[len - 1])
	// 		-> weights
	// 		-> biases

	// Validate output dimensions match expected dimensions

	// loop backwards through weights:
	//		-> calculate dC/dW and dC/dB at each layer
	//		-> push to deltaWeights, deltaBiases matrices respectively
	//			note: since we're looping backwards we need to reverse. Could also unset to front instead of pushing

	var deltaWeights [][][]float64
	var deltaBaises [][]float64

	for l := len(weights) - 1; l >= 0; l-- {
		dW, dB, err := calculateDeltasAtLayer(l)
		if err != nil {
			return [][][]float64{}, [][]float64{}, err
		}
		deltaWeights = append([][][]float64{dW}, deltaWeights...)
		deltaBaises = append([][]float64{dB}, deltaBaises...)
	}

	return deltaWeights, deltaBaises, nil
}
