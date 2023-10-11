package neuralnetwork

import (
	"fmt"
	"goml/math/matrix"
	"testing"
)

func setupSimpleNetwork() {
	activations = [][]float64{
		{0, 0, 1, 0},
		{1, 0, 1},
		{1, 0},
	}
	biases = [][]float64{
		{0, 0, 0},
		{0.25, 0.25},
	}
	weights = [][][]float64{
		{
			{1, 1, 1, 1},
			{2, 2, 2, 2},
			{3, 3, 3, 3},
		},
		{
			{1, 1, 1},
			{2, 2, 2},
		},
	}
}

func isEqual(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestZ(t *testing.T) {
	setupSimpleNetwork()
	defer cleanup()

	actual0, err := z(0)
	expected0 := []float64{1, 2, 3}
	if err != nil {
		t.Fatalf("error %v", err)
	}

	actual1, err := z(1)
	expected1 := []float64{2.25, 4.25}
	if err != nil {
		t.Fatalf("error %v", err)
	}

	if !isEqual(actual0, expected0) {
		fmt.Printf("expected %v but got %v", expected0, actual0)
	}

	if !isEqual(actual1, expected1) {
		fmt.Printf("expected %v but got %v", expected1, actual1)
	}
}

func TestDCDA_dimensions(t *testing.T) {
	defer cleanup()

	Initialize(10, 4, Config{[]int{8, 5}, 0.1})

	input := []float64{0.01, .1, .2, .3, .4, .5, .6, .7, .8, .9}
	tOutput := []float64{0.5, 0.5, 0.5, 0.5}
	feedforward(input)

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

func TestDCDA_output(t *testing.T) {
	setupSimpleNetwork()
	defer cleanup()

	expected := []float64{0, 0}
	actual, err := dCdA(1, []float64{1, 0})
	if err != nil {
		t.Fatalf("error %v", err)
	}
	if !isEqual(actual, expected) {
		t.Fatalf("expected %v but got %v", expected, actual)
	}
}

func TestCalculateDeltas_dimensions(t *testing.T) {
	defer cleanup()

	Initialize(10, 4, Config{[]int{8, 5}, 0.1})
	input := []float64{0.01, .1, .2, .3, .4, .5, .6, .7, .8, .9}
	y := []float64{0.5, 0.5, 0.5, 0.5}
	feedforward(input)

	for l, w := range weights {
		dw, db, err := calculateDeltas(l, y)
		if err != nil {
			t.Fatalf("error %v", err)
		}

		valid := matrix.IsEqualDimensions2D(dw, w) && matrix.IsEqualDimensions1D(db, biases[l])

		if !valid {
			t.Fail()
		}
	}
}
