package neuralnetwork

import (
	"goml/math/formulas"
	"goml/math/matrix"
)

// z(L) -> w(l) * a(l-1) + b(l)
func z(l int) ([]float64, error) {
	p, err := matrix.DotWeightsAndActivations(weights[l], activations[l])
	if err != nil {
		return []float64{}, err
	}
	return matrix.Add1D(p, biases[l])
}

// dC/dA = ∑ w(l+1) • s'(z(l+1)) • dC/dA(l+1) OR deltacost(output, y)
func dCdA(l int, y []float64) ([]float64, error) {
	if l == len(weights)-1 {
		return formulas.DeltaCost(activations[l+1], y)
	}

	zL1, err := z(l + 1)
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

// dC/dW and dC/dB for a given layer
func calculateDeltas(l int, y []float64) ([][]float64, []float64, error) {

	// s'(z(l))
	zl, err := z(l)
	szl := matrix.Map1D(zl, formulas.DeltaSigmoid)
	if err != nil {
		return [][]float64{}, []float64{}, err
	}

	// dC/dA(l)
	dcda, err := dCdA(l, y)
	if err != nil {
		return [][]float64{}, []float64{}, err
	}

	// dC/dB -> s'(z(l)) • dC/dA(l)
	dcdb, err := matrix.Hadamard1D(dcda, szl)
	if err != nil {
		return [][]float64{}, []float64{}, err
	}

	// dC/dW -> a(l-1) • s'(z(l)) • dC/dA(l)
	dcdw := matrix.DotToCreateWeights(dcdb, activations[l])

	return dcdw, dcdb, nil
}
