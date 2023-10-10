package main

import (
	"fmt"
	formulas "goml/math"
	"goml/math/matrix"
)

func dCdA(l int, y []float64) ([]float64, error) {
	if l == len(weights)-1 {
		return formulas.DeltaCost(activations[l+1], y)
	}

	zL1, err := CalculateZL(l + 1)
	if err != nil {
		return []float64{}, err
	}

	nextA, err := dCdA(l+1, y)
	if err != nil {
		return []float64{}, err
	}

	p, err := matrix.Hadamard1D(matrix.Map1D(zL1, formulas.DeltaSigmoid), nextA)
	if err != nil {
		return []float64{}, err
	}

	return matrix.DotWeightsAndActivations(matrix.Transpose(weights[l+1]), p)
}

func calculateDeltasAtLayer(l int, y []float64) ([][]float64, []float64, error) {
	// calculate z(l) -> w(l) * a(l-1) + b(l)

	// dC/dB -> deltaSig(z(l)) • dC/dA
	// dC/dW -> a(l-1) • deltaSig(z(l)) • dC/dA

	dcda, err := dCdA(l, y)
	if err != nil {
		fmt.Printf("err:%v", err.Error())
	}

	// fmt.Printf("\n\nlayer: %v", l)

	zl, _ := CalculateZL(l)
	// fmt.Printf("\n zl %v", zl)
	dzl := matrix.Map1D(zl, formulas.DeltaSigmoid)
	// fmt.Printf("\n dzl %v", dzl)

	deltaBiases, err := matrix.Hadamard1D(dcda, dzl)
	// fmt.Printf("\n deltaBiases %v", deltaBiases)
	if err != nil {
		fmt.Printf("err:%v", err.Error())
	}
	deltaWeights := matrix.DotToCreateWeights(deltaBiases, activations[l])
	// fmt.Printf("\n deltaWeights %v", deltaWeights)

	return deltaWeights, deltaBiases, nil
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

	Feedforward(tInput)

	var deltaWeights [][][]float64
	var deltaBaises [][]float64

	for l := len(weights) - 1; l >= 0; l-- {
		dW, dB, err := calculateDeltasAtLayer(l, tOutput)
		if err != nil {
			return [][][]float64{}, [][]float64{}, err
		}
		deltaWeights = append([][][]float64{dW}, deltaWeights...)
		deltaBaises = append([][]float64{dB}, deltaBaises...)
	}

	return deltaWeights, deltaBaises, nil
}
