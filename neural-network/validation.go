package neuralnetwork

// TODO:
func IsInitialized() bool {
	if rate <= 0 || rate >= 1 {
		return false
	}

	if len(layers) < 3 {
		return false
	}

	// rate must be float between 0 and 1
	// layers must be > len 3
	// activation dimensions should match layers
	// bias dimensions should match activations (excluding input layer)
	// weight dimensions should match activations
	//  len(w[i]) -> len(a[i+1])
	//	len(w[i][j]) -> len(a[i])

	return true

}

func isValidInput(input []float64) bool {
	return len(input) == len(activations[0])
}

func isValidOutput(output []float64) bool {
	return len(output) == len(activations[len(activations)-1])
}
