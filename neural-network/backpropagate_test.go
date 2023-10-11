package neuralnetwork

import (
	"goml/validation"
	"testing"
)

func TestBackpropagate(t *testing.T) {
	defer cleanup()

	input_size := 10
	output_size := 4
	Initialize(input_size, output_size, Config{[]int{8, 5}, 0.1})

	input := []float64{0.01, .1, .2, .3, .4, .5, .6, .7, .8, .9}

	dw, db, err := backpropagate(input, []float64{0.5, 0.5, 0.5, 0.5})

	if err != nil {
		t.Fatalf("error %v", err)
	}

	if !validation.IsEqualDimensions3D(dw, weights) {
		t.Fatalf("dC/dW dimensions should match weights dimensions")
	}

	if !validation.IsEqualDimensions2D(db, biases) {
		t.Fatalf("dC/dB dimensions should match biases dimensions")
	}
}
