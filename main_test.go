package main

import (
	"testing"
)

func TestInitializeDefault(t *testing.T) {
	dimensions := []int{3, 8, 5, 12, 2}
	config := Config{[]int{8, 5, 12}, 0.12}
	initialize(3, 2, config)

	// init learning rate
	if rate != 0.12 {
		t.Fatalf("Learning rate was not initialized correctly")
	}

	// init layer dimensions
	for i, layer := range layers {
		if dimensions[i] != layer {
			t.Fatalf("Layer dimensions were not initialized correctly")
		}
	}

	// init weights
	if len(weights) != 4 {
		t.Fatalf("Bad weight dimensions at the first level")
	}
	if len(weights[0]) != 8 {
		t.Fatalf("Bad weight dimensions at the second level")
	}
	if len(weights[0][0]) != 3 {
		t.Fatalf("Bad weight dimensions at the third level")
	}

	// init biases
	if len(biases) != 4 {
		t.Fatalf("Bad bias dimensions at the first level")
	}
	if len(biases[0]) != 8 {
		t.Fatalf("Bad bias dimensions at the second level")
	}

	// ensure valid matrices
	for i := range weights {
		size := len(weights[i][0])
		for j := range weights[i] {
			if len(weights[i][j]) != size {
				t.Fatalf("Weights matrix is not 'square'")
			}
		}
	}
}

func TestInitializeConfig(t *testing.T) {
	dimensions := []int{6, 3, 3, 2}
	defaultConfig := Config{getDefaultLayerSizes(), DEFAULT_LEARNING_RATE}
	initialize(6, 2, defaultConfig)

	// init learning rate
	if rate != DEFAULT_LEARNING_RATE {
		t.Fatalf("Learning rate was not initialized correctly")
	}

	// init layer dimensions
	for i, layer := range layers {
		if dimensions[i] != layer {
			t.Fatalf("Layer dimensions were not initialized correctly")
		}
	}

	// init weights
	if len(weights) != 3 {
		t.Fatalf("Bad weight dimensions at the first level")
	}
	if len(weights[0]) != 3 {
		t.Fatalf("Bad weight dimensions at the second level")
	}
	if len(weights[0][0]) != 6 {
		t.Fatalf("Bad weight dimensions at the third level")
	}

	// init biases
	if len(biases) != 3 {
		t.Fatalf("Bad bias dimensions at the first level")
	}
	if len(biases[0]) != 3 {
		t.Fatalf("Bad bias dimensions at the second level")
	}

	// ensure valid matrices
	for i := range weights {
		size := len(weights[i][0])
		for j := range weights[i] {
			if len(weights[i][j]) != size {
				t.Fatalf("Weights matrix is not 'square'")
			}
		}
	}
}
