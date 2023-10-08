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

func TestBackpropagate(t *testing.T) {
	defer cleanup()

	Backpropagate([]float64{}, []float64{})

}
