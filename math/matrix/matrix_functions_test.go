package matrix

import (
	"testing"
)

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

func isEqual2D(a, b [][]float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}

	}
	return true
}

func TestDotCreateWeights(t *testing.T) {
	right := []float64{
		1,
		2,
		3,
		4,
	}
	left := []float64{
		1, 2, 3,
	}

	expected := [][]float64{
		{1, 2, 3},
		{2, 4, 6},
		{3, 6, 9},
		{4, 8, 12},
	}
	actual := DotToCreateWeights(right, left)

	if !isEqual2D(actual, expected) {
		t.Fatalf("expected %v but got %v", expected, actual)
	}

}

func TestDotCase1(t *testing.T) {
	w := [][]float64{
		{1, 3, -5},
	}
	a := []float64{4, -2, -1}
	expected := []float64{3}

	actual, err := DotWeightsAndActivations(w, a)

	if err != nil {
		t.Fatalf("error received but expected success")
	}

	if !isEqual(actual, expected) {
		t.Fatalf("expected: %v | actual: %v", expected, actual)
	}
}

func TestDotCase2(t *testing.T) {
	w := [][]float64{
		{1, 3, -5},
		{0, 0, 0},
		{1, 2, 3},
		{-1, -2, -3},
	}
	a := []float64{4, -2, -1}
	expected := []float64{3, 0, -3, 3}

	actual, err := DotWeightsAndActivations(w, a)

	if err != nil {
		t.Fatalf("error received but expected success")
	}

	if !isEqual(actual, expected) {
		t.Fatalf("expected: %v | actual: %v", expected, actual)
	}
}

func TestDotFail(t *testing.T) {
	w := [][]float64{
		{1, 3, -5},
		{0, 0, 0},
		{1, 2, 3},
		{-1, -2, -3},
	}
	a := []float64{4, -2, -1, 1}
	_, err := DotWeightsAndActivations(w, a)

	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestHadamardPass(t *testing.T) {
	a := [][]float64{
		{1, 3, -5},
		{0, 0, 0},
		{1, 2, 3},
		{-1, -2, -3},
	}
	b := [][]float64{
		{2, 3, 4},
		{0, 2, 0},
		{1, 9, 3},
		{-1, -6, -1},
	}
	expected := [][]float64{
		{2, 9, -20},
		{0, 0, 0},
		{1, 18, 9},
		{1, 12, 3},
	}

	actual, err := Hadamard2D(a, b)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(actual) != len(expected) {
		t.Fatalf("expected %v \n actual %v", expected, actual)
	}

	for i := range actual {
		if len(actual[i]) != len(expected[i]) {
			t.Fatalf("expected %v \n actual %v", expected, actual)
		}
		for j := range actual[i] {
			if actual[i][j] != expected[i][j] {
				t.Fatalf("expected %v \n actual %v", expected, actual)
			}
		}
	}
}

func TestHadamardFail(t *testing.T) {
	a1 := [][]float64{
		{1, 3, -5},
		{0, 0, 0},
		{1, 2, 3},
		{-1, -2, -3},
	}
	b1 := [][]float64{
		{2, 3, 4},
		{0, 2, 0},
		{1, 9, 3},
	}

	_, err := Hadamard2D(a1, b1)

	if err == nil {
		t.Fatalf("test1 expected error")
	}

	a2 := [][]float64{
		{1, 3, -5},
		{0, 0, 0},
		{1, 2, 3},
		{-1, -2, -3},
	}
	b2 := [][]float64{
		{2, 3, 4},
		{0, 2, 0},
		{1, 3},
		{-1, -6, -1},
	}

	_, err = Hadamard2D(a2, b2)

	if err == nil {
		t.Fatalf("test2 expected error")
	}
}

func TestAdd1DPass(t *testing.T) {
	a := []float64{1, 1, 1}
	b := []float64{9, -1, 0}
	expected := []float64{10, 0, 1}
	actual, err := Add1D(a, b)

	if err != nil {
		t.Fatalf("error received but expected success")
	}

	if !isEqual(actual, expected) {
		t.Fatalf("expected: %v | actual: %v", expected, actual)
	}
}

func TestAdd1DFail(t *testing.T) {
	a := []float64{1, 1, 1}
	b := []float64{9, -1, 0, 1}
	_, err := Add1D(a, b)

	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestMap1D(t *testing.T) {
	m := []float64{1, 2, 3}
	expected := []float64{1, 4, 9}
	actual := Map1D(m, func(v float64) float64 {
		return v * v
	})

	if !isEqual(actual, expected) {
		t.Fatalf("expected: %v | actual: %v", expected, actual)
	}
}
