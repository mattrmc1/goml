package formulas

import (
	"math"
	"testing"
)

func TestSigmoid_BigPositiveNumbers(t *testing.T) {
	expected := float64(1)
	actual := sigmoid(float64(29357598345863))

	if actual != expected {
		t.Fatalf("Expected %v but got %v", expected, actual)
	}
}

func TestSigmoid_BigNegativeNumbers(t *testing.T) {
	expected := float64(0)
	actual := sigmoid(float64(-29357598345863))

	if actual != expected {
		t.Fatalf("Expected %v but got %v", expected, actual)
	}
}

func TestSigmoid_CloseToZero(t *testing.T) {
	half := float64(0.5)
	actual_gt := sigmoid(float64(0.00145))
	actual_lt := sigmoid(float64(-0.00145))

	if actual_gt <= half {
		t.Fatalf("Small numbers greater than zero should be greater than 0.5 (Value: %v)", actual_gt)
	}

	if actual_lt >= half {
		t.Fatalf("Small numbers less than zero should be less than 0.5 (Value: %v)", actual_lt)
	}

	if math.Abs(half-actual_lt) > float64(0.001) || math.Abs(half-actual_gt) > float64(0.001) {
		t.Fatalf("Imprecise result(s): %v | %v", actual_lt, actual_gt)
	}
}

func TestSigmoid_ExactlyZero(t *testing.T) {
	expected := float64(0.5)
	actual := sigmoid(float64(0))

	if actual != expected {
		t.Fatalf("Expected %v but got %v", expected, actual)
	}
}

func TestDeltaSigmoid_BigPositiveNumbers(t *testing.T) {
	expected := float64(0)
	actual := deltaSigmoid(float64(29357598345863))

	if actual != expected {
		t.Fatalf("Expected %v but got %v", expected, actual)
	}
}

func TestDeltaSigmoid_BigNegativeNumbers(t *testing.T) {
	expected := float64(0)
	actual := deltaSigmoid(float64(-29357598345863))

	if actual != expected {
		t.Fatalf("Expected %v but got %v", expected, actual)
	}
}

func TestDeltaSigmoid_CloseToZero(t *testing.T) {
	quarter := float64(0.25)
	actual_gt := deltaSigmoid(float64(-0.001974))
	actual_lt := deltaSigmoid(float64(0.001974))

	if actual_gt >= quarter || actual_lt >= quarter {
		t.Fatalf("Should be less than 0.25 (Value: %v)", actual_gt)
	}

	if math.Abs(quarter-actual_lt) > float64(0.000001) || math.Abs(quarter-actual_gt) > float64(0.000001) {
		t.Fatalf("Imprecise result(s): %v | %v", actual_lt, actual_gt)
	}
}

func TestDeltaSigmoid_ExactlyZero(t *testing.T) {
	quarter := float64(0.25)
	actual := deltaSigmoid(float64(0))

	if actual != quarter {
		t.Fatalf("Zero should result in exactly 1/4")
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

func TestCostFunction(t *testing.T) {
	a := []float64{1, 1, 1, 1, 1}
	y := []float64{2, 3, -1, -3, 1}

	expected := []float64{1, 4, 4, 16, 0}
	actual, err := cost(a, y)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !isEqual(actual, expected) {
		t.Fatalf("Expected %v but got %v", expected, actual)
	}
}

func TestCostFunctionError(t *testing.T) {
	a := []float64{1, 1, 1, 1}
	y := []float64{1, 1, 1, 1, 1}

	_, err := cost(a, y)

	if err == nil {
		t.Fatalf("Cost function should have errored because the input lengths aren't equal")
	}
}

func TestDeltaCostFunction(t *testing.T) {
	a := []float64{1, 1, 1, 1, 1}
	y := []float64{2, 3, -1, -3, 1}

	expected := []float64{-2, -4, 4, 8, 0}
	actual, err := deltaCost(a, y)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !isEqual(actual, expected) {
		t.Fatalf("Expected %v but got %v", expected, actual)
	}
}

func TestDeltaCostFunctionError(t *testing.T) {
	a := []float64{1, 1, 1, 1}
	y := []float64{1, 1, 1, 1, 1}

	_, err := deltaCost(a, y)

	if err == nil {
		t.Fatalf("Cost function should have errored because the input lengths aren't equal")
	}
}
