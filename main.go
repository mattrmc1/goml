package main

import (
	"fmt"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

type TrainingData struct {
	tInput  []float64
	tOutput []float64
}

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

func Initialize(input int, output int, config Config) {
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
				weights[i-1][j][k] = rand.Float64() * .5
			}
		}

		biases[i-1] = make([]float64, layers[i])
		for j := range biases[i-1] {
			biases[i-1][j] = float64(0)
		}
	}
}

func Train() {
	// initialize based on training data dimensions

	tmp := []TrainingData{
		{
			[]float64{1, 0, 0, 0},
			[]float64{1, 1},
		},
		{
			[]float64{0, 0, 1, 0},
			[]float64{1, 1},
		},
		{
			[]float64{0, 1, 0, 0},
			[]float64{0, 0},
		},
		{
			[]float64{0, 0, 0, 1},
			[]float64{0, 0},
		},
	}

	// tmp init
	Initialize(4, 2, Config{[]int{5, 6}, DEFAULT_LEARNING_RATE})

	tmpMaxIter := 10000

	fmt.Printf("\n before %v \n", weights)

	for i := 0; i < tmpMaxIter; i++ {
		for _, td := range tmp {

			deltaWeights, _, err := Backpropagate(td.tInput, td.tOutput)
			if err != nil {
				fmt.Printf("\n err %v", err)
				return
			}
			for i := range weights {
				for j := range weights[i] {
					for k := range weights[i][j] {
						weights[i][j][k] = weights[i][j][k] - deltaWeights[i][j][k]*rate
					}
				}
			}
			// fmt.Printf("\n dW %v \n", deltaWeights)
			// fmt.Printf("\n dB %v \n", deltaBiases)
		}
	}

	fmt.Printf("\n after %v \n", weights)

	for _, td := range tmp {
		output, _ := Feedforward(td.tInput)
		fmt.Printf("\n\n Expected: %v \n Actual: %v \n\n", output, td.tOutput)
	}
}

func Run() {
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
