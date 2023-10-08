package main

import (
	"testing"
)

func TestFeedForward(t *testing.T) {
	defer Cleanup()

	input_size := 10
	output_size := 4
	initialize(input_size, output_size, Config{[]int{8, 5}, 0.1})

	input := []float64{0.01, .1, .2, .3, .4, .5, .6, .7, .8, .9}
	res, err := Feedforward(input)

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

	for i := range activations {
		for j, a := range activations[i] {
			if a < 0 || a > 1 {
				t.Fatalf("activation nodes should always be a float between 0 and 1")
			}
			if i == 0 && input[j] != a {
				t.Fatalf("first activation layer should match the input")
			}
			if i == len(layers)-1 && res[j] != a {
				t.Fatalf("last activation layer should match the output result")
			}

		}
	}
}
