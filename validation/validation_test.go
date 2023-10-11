package validation

import (
	"testing"
)

func TestIsEqualDimensions1D(t *testing.T) {
	if !IsEqualDimensions1D([]float64{1, 1, 1}, []float64{4, 3, 6}) {
		t.Fatalf("expected to pass but failed (1D)")
	}
	if IsEqualDimensions1D([]float64{1, 1}, []float64{4, 3, 6}) {
		t.Fatalf("expected to fail but passed (1D)")
	}
}

func TestIsEqualDimensions2D(t *testing.T) {
	shouldReturnTrue := IsEqualDimensions2D(
		[][]float64{
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
		},
		[][]float64{
			{0, 1, 1, 30},
			{1, 20, 1, 1},
			{1, 1, 10, 1},
		},
	)

	if !shouldReturnTrue {
		t.Fail()
	}

	shouldReturnFalse1 := IsEqualDimensions2D(
		[][]float64{
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
		},
		[][]float64{
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
		},
	)

	if shouldReturnFalse1 {
		t.Fail()
	}

	shouldReturnFalse2 := IsEqualDimensions2D(
		[][]float64{
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
		},
		[][]float64{
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1},
		},
	)

	if shouldReturnFalse2 {
		t.Fail()
	}
}

func TestIsEqualDimensions3D(t *testing.T) {
	shouldReturnTrue := IsEqualDimensions3D(
		[][][]float64{
			{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
		},
		[][][]float64{
			{
				{0, 1, 1, 30},
				{1, 20, 1, 1},
			},
			{
				{1, 20, 1, 1},
				{1, 1, 10, 1},
			},
		},
	)

	if !shouldReturnTrue {
		t.Fail()
	}

	shouldReturnFalse1 := IsEqualDimensions3D(
		[][][]float64{
			{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
		},
		[][][]float64{
			{
				{0, 1, 1, 30},
				{1, 20, 1, 1},
			},
		},
	)

	if shouldReturnFalse1 {
		t.Fail()
	}

	shouldReturnFalse2 := IsEqualDimensions3D(
		[][][]float64{
			{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
		},
		[][][]float64{
			{
				{1, 20, 1, 1},
			},
			{
				{1, 1, 10, 1},
			},
		},
	)

	if shouldReturnFalse2 {
		t.Fail()
	}

	shouldReturnFalse3 := IsEqualDimensions3D(
		[][][]float64{
			{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
		},
		[][][]float64{
			{
				{0, 1, 30},
				{1, 20, 1},
			},
			{
				{1, 20, 1},
				{1, 10, 1},
			},
		},
	)

	if shouldReturnFalse3 {
		t.Fail()
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
