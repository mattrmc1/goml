package main

import (
	"fmt"
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
var activations [][]float64
var weights [][][]float64
var biases [][]float64

func initialize(input int, output int, config Config) {
	layers = nil
	layers = append(layers, input)
	layers = append(layers, config.layers...)
	layers = append(layers, output)

	rate = config.learning_rate

	activations = make([][]float64, len(layers))
	weights = make([][][]float64, len(layers)-1)
	biases = make([][]float64, len(layers)-1)

	for i := 0; i < len(layers); i++ {

		activations[i] = make([]float64, layers[i])
		if i == 0 {
			continue
		}

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

func train() {
	// initialize based on training data dimensions

	// var deltaWeights, deltaBiases

	// foreach training data point {
	//		output := feedforward(tInput)
	//		deltaWeights, deltaBiases := backpropagate(output, tOutput)
	//		--> sum deltas and keep going (apply deltas now?)
	// }

	// Squeeze the deltas --> apply deltas?
}

func run() {
	// validate initialization

	// feedforward on input without storing activation nodes
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
