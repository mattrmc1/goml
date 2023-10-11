package neuralnetwork

import (
	"fmt"
	"math/rand"
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

// Tmp
func Train(t []TrainingData, maxIter int) {

	// validate initialization (maybe just...initialize?)
	// validate training data

	fmt.Printf("\n before %v \n", weights)

	for i := 0; i < maxIter; i++ {
		for _, td := range t {

			deltaWeights, _, err := backpropagate(td.tInput, td.tOutput)
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
		}
	}

	fmt.Printf("\n after %v \n", weights)

	for _, td := range t {
		output, _ := feedforward(td.tInput)
		fmt.Printf("\n\n Expected: %v \n Actual: %v \n\n", output, td.tOutput)
	}
}

func Run(input []float64) ([]float64, error) {
	// validate initialization
	// validate input

	return feedforward(input)
}
