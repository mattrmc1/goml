package matrix

type validator func(float64) bool

func IsEqualDimensions1D(a, b []float64) bool {
	return len(a) == len(b)
}

func IsEqualDimensions2D(a, b [][]float64) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
	}

	return true
}

func IsEqualDimensions3D(a, b [][][]float64) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if len(a[i][j]) != len(b[i][j]) {
				return false
			}
		}
	}

	return true
}

func Validate1D(a []float64, fn validator) bool {
	for _, v := range a {
		if !fn(v) {
			return false
		}
	}

	return true
}

func Validate2D(a [][]float64, fn validator) bool {
	for i := range a {
		for j := range a[i] {
			if !fn(a[i][j]) {
				return false
			}
		}
	}

	return true
}
