package neuralnetwork

func backpropagate(tInput, tOutput []float64) ([][][]float64, [][]float64, error) {

	feedforward(tInput)

	var deltaWeights [][][]float64
	var deltaBiases [][]float64

	for l := len(weights) - 1; l >= 0; l-- {
		dW, dB, err := calculateDeltas(l, tOutput)
		if err != nil {
			return deltaWeights, deltaBiases, err
		}

		deltaWeights = append([][][]float64{dW}, deltaWeights...)
		deltaBiases = append([][]float64{dB}, deltaBiases...)
	}

	return deltaWeights, deltaBiases, nil
}
