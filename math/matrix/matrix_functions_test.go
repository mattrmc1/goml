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

func TestDotCase1(t *testing.T) {
	w := [][]float64{
		{1, 3, -5},
	}
	a := []float64{4, -2, -1}
	expected := []float64{3}

	actual, err := Dot(w, a)

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

	actual, err := Dot(w, a)

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
	_, err := Dot(w, a)

	if err == nil {
		t.Fatalf("expected error")
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
