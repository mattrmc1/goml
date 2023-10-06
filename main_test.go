package main

import (
	"testing"
)

func cleanup() {
	rate = 0
	layers = nil
	weights = nil
	biases = nil
}

func TestInitialize(t *testing.T) {
	defer cleanup()

	dimensions := []int{3, 8, 5, 12, 2}
	config := Config{[]int{8, 5, 12}, 0.12}
	initialize(3, 2, config)

	// init learning rate
	if rate != config.learning_rate {
		t.Fatalf("Learning rate was not initialized correctly")
	}

	// init layer dimensions
	for i, layer := range layers {
		if dimensions[i] != layer {
			t.Fatalf("Layer dimensions were not initialized correctly: %v", layers)
		}
	}

	// init weights
	if len(weights) != 4 {
		// one per activation excluding the input layer
		t.Fatalf("Bad weight dimensions at the first level")
	}
	if len(weights[0]) != 8 ||
		len(weights[1]) != 5 ||
		len(weights[2]) != 12 ||
		len(weights[3]) != 2 {
		// size of weight[i] should equal activation[i+1]
		t.Fatalf("Bad weight dimensions at the second level")
	}
	if len(weights[0][0]) != 3 ||
		len(weights[1][0]) != 8 ||
		len(weights[2][0]) != 5 ||
		len(weights[3][0]) != 12 {
		// size of weight[i][j] should equal activation[i]
		t.Fatalf("Bad weight dimensions at the third level")
	}
	// ensure len w[i][j] match per i (i.e. they're valid matrices)
	for i := range weights {
		size := len(weights[i][0])
		for j := range weights[i] {
			if len(weights[i][j]) != size {
				t.Fatalf("Weights matrix is not 'square'")
			}
		}
	}

	// init biases
	if len(biases) != 4 {
		t.Fatalf("Bad bias dimensions at the first level")
	}
	if len(biases[0]) != 8 ||
		len(biases[1]) != 5 ||
		len(biases[2]) != 12 ||
		len(biases[3]) != 2 {
		t.Fatalf("Bad bias dimensions at the second level")
	}
}

func TestFeedForward(t *testing.T) {
	defer cleanup()

	input_size := 10
	output_size := 4
	initialize(input_size, output_size, Config{[]int{8, 5}, 0.1})

	input := []float64{0.01, .1, .2, .3, .4, .5, .6, .7, .8, .9}
	res, err := feedforward(input)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(res) != output_size {
		t.Fatalf("expected output size: %v but got %v", output_size, len(res))
	}

	for _, v := range res {
		if v < 0 || v > 1 {
			t.Fatalf("received output values out of range: %v", res)
		}
	}

}
