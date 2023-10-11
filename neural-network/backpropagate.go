package neuralnetwork

func backpropagate(tInput, tOutput []float64) ([][][]float64, [][]float64, error) {
	// Validate initialization
	//		-> learning rate should be defined
	//		-> layers[0] dimensions should match tInput
	//		-> layers[len-1] dimensions should match tOutput
	//		-> weights and biases should be initialized with dummy data

	// Validate output dimensions match expected dimensions

	feedforward(tInput)

	var deltaWeights [][][]float64
	var deltaBaises [][]float64

	for l := len(weights) - 1; l >= 0; l-- {
		dW, dB, err := calculateDeltas(l, tOutput)
		if err != nil {
			panic(err)
		}

		deltaWeights = append([][][]float64{dW}, deltaWeights...)
		deltaBaises = append([][]float64{dB}, deltaBaises...)
	}

	return deltaWeights, deltaBaises, nil
}
