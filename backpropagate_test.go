package main

import (
	"testing"
)

func TestDCDA(t *testing.T) {
	defer Cleanup()

	Initialize(10, 4, Config{[]int{8, 5}, 0.1})

	input := []float64{0.01, .1, .2, .3, .4, .5, .6, .7, .8, .9}
	tOutput := []float64{0.5, 0.5, 0.5, 0.5}
	Feedforward(input)

	res0, err := dCdA(0, tOutput)
	if len(res0) != 8 {
		t.Fatalf("expected dC/dA at layer %v to be %v but got %v", 0, 8, len(res0))
	}
	if err != nil {
		t.Fatalf("error %v", err)
	}

	res1, err := dCdA(1, tOutput)
	if len(res1) != 5 {
		t.Fatalf("expected dC/dA at layer %v to be %v but got %v", 1, 5, len(res1))
	}
	if err != nil {
		t.Fatalf("error %v", err)
	}

	res2, err := dCdA(2, tOutput)
	if len(res2) != 4 {
		t.Fatalf("expected dC/dA at layer %v to be %v but got %v", 2, 4, len(res2))
	}
	if err != nil {
		t.Fatalf("error %v", err)
	}
}

func TestBackpropagate(t *testing.T) {
	defer Cleanup()

	input_size := 10
	output_size := 4
	Initialize(input_size, output_size, Config{[]int{8, 5}, 0.1})

	input := []float64{0.01, .1, .2, .3, .4, .5, .6, .7, .8, .9}

	dw, db, err := Backpropagate(input, []float64{0.5, 0.5, 0.5, 0.5})

	if err != nil {
		t.Fatalf("error %v", err)
	}

	if len(dw) != len(weights) {
		t.Fatalf("dC/dW dimensions should match weights dimensions")
	}
	for i, v := range dw {
		if len(v) != len(weights[i]) {
			t.Fatalf("dC/dW dimensions should match weights dimensions")
		}
		for j, v2 := range v {
			if len(v2) != len(weights[i][j]) {
				t.Fatalf("dC/dW dimensions should match weights dimensions")
			}
		}
	}

	if len(db) != len(biases) {
		t.Fatalf("dC/dB dimensions should match biases dimensions")
	}
	for i, v := range db {
		if len(v) != len(biases[i]) {
			t.Fatalf("dC/dB dimensions should match biases dimensions")
		}
	}
}
