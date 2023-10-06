package validation

import (
	"testing"
)

func TestIsEqualDimensions(t *testing.T) {
	if !IsEqualDimensions1D([]float64{1, 1, 1}, []float64{4, 3, 6}) {
		t.Fatalf("expected to pass but failed (1D)")
	}
	if IsEqualDimensions1D([]float64{1, 1}, []float64{4, 3, 6}) {
		t.Fatalf("expected to fail but passed (1D)")
	}

	shouldPass := [][][]float64{
		{
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
		},
		{
			{0, 1, 1, 30},
			{1, 20, 1, 1},
			{1, 1, 10, 1},
		},
	}
	shouldFail1 := [][][]float64{
		{
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
		},
		{
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
		},
	}
	shouldFail2 := [][][]float64{
		{
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
		},
		{
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
		},
	}

	if !IsEqualDimensions2D(shouldPass[0], shouldPass[1]) {
		t.Fatalf("expected to pass but failed (2D)")
	}

	if IsEqualDimensions2D(shouldFail1[0], shouldFail1[1]) {
		t.Fatalf("expected to fail but passed (2D rows)")
	}

	if IsEqualDimensions2D(shouldFail2[0], shouldFail2[1]) {
		t.Fatalf("expected to fail but passed (2D cols)")
	}
}

func predicate(v float64) bool {
	return v <= 1
}

func TestValidate1D(t *testing.T) {
	shouldPass := []float64{1, 1, 1, 1}
	shouldFail := []float64{1, 10, 1, 1}

	if !Validate1D(shouldPass, predicate) {
		t.Fatalf("expected to pass but failed")
	}

	if Validate1D(shouldFail, predicate) {
		t.Fatalf("expected to fail but passed")
	}
}

func TestValidate2D(t *testing.T) {
	shouldPass := [][]float64{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}

	shouldFail := [][]float64{
		{1, 1, 1, 1},
		{1, 10, 1, 1},
		{1, 1, 1, 1},
	}

	if !Validate2D(shouldPass, predicate) {
		t.Fatalf("expected to pass but failed")
	}

	if Validate2D(shouldFail, predicate) {
		t.Fatalf("expected to fail but passed")
	}
}
