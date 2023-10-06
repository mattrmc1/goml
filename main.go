package main

import (
	"errors"
	"fmt"
	formulas "goml/math"
	"goml/math/matrix"
	"goml/validation"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

type Config struct {
	layers        []int
	learning_rate float64
}

const (
	DEFAULT_LEARNING_RATE = 0.1
)

var rate float64
var layers []int
var weights [][][]float64
var biases [][]float64

func initialize(input int, output int, config Config) {
	layers = nil
	layers = append(layers, input)
	layers = append(layers, config.layers...)
	layers = append(layers, output)

	rate = config.learning_rate

	weights = make([][][]float64, len(layers)-1)
	biases = make([][]float64, len(layers)-1)

	for i := 1; i < len(layers); i++ {

		weights[i-1] = make([][]float64, layers[i])
		for j := range weights[i-1] {
			weights[i-1][j] = make([]float64, layers[i-1])
			for k := range weights[i-1][j] {
				weights[i-1][j][k] = rand.Float64()
			}
		}

		biases[i-1] = make([]float64, layers[i])
		for j := range biases[i-1] {
			biases[i-1][j] = rand.Float64()
		}
	}
}

func feedforward(input []float64) ([]float64, error) {
	if len(layers) == 0 {
		return []float64{}, errors.New("network not initialized correctly")
	}

	if len(input) != layers[0] {
		return []float64{}, fmt.Errorf("invalid input - expected size: %v", layers[0])
	}

	if validation.Validate1D(input, func(v float64) bool {
		return v >= 0 && v <= 1
	}) {
		return []float64{}, errors.New("all input values must be a float between 0 and 1")
	}

	var a []float64
	copy(a, input)

	// a = sigmoid(dot(w, a) + b)
	for i := range weights {
		w := weights[i]
		b := biases[i]

		p, err := matrix.Dot(w, a)
		if err != nil {
			return []float64{}, err
		}

		zl, err := matrix.Add1D(p, b)
		if err != nil {
			return []float64{}, err
		}

		a = matrix.Map1D(zl, formulas.Sigmoid)
	}

	return a, nil
}

func backpropagate(output, y []float64) ([][][]float64, [][]float64, error) {
	// validate output dimensions equal y dimensions

	// w := calculate weight deltas
	//		-> dCdA -> either cost func or

	// b := calculate biases deltas

	return [][][]float64{}, [][]float64{}, nil
}

func train() {
	// initialize based on training data dimensions
	// foreach training data point {
	//		var output := feedforward
	//		var cost := backpropagate (note the weights and biases get changed in this function)
	//
	// }
}

func main() {
	a := mat.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})

	// Create a matrix formatting value with a prefix and calculating each column
	// width individually...
	fa := mat.Formatted(a, mat.Prefix("    "), mat.Squeeze())

	// and then print with and without zero value elements.
	fmt.Printf("with all values:\na = %v\n\n", fa)
	fmt.Printf("with only non-zero values:\na = % v\n\n", fa)
}
