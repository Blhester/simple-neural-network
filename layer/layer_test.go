package layer

import (
	"math/rand"
	"testing"
	"time"
)

const MATRIX_DIMENSION = 100

func BenchmarkNeuronOutput(b *testing.B) {
	b.Logf("Running ( %d ) times", b.N)

	// Initialize a neuron with random inputs and weights
	inputs := make([][]float64, MATRIX_DIMENSION)
	weights := make([][]float64, MATRIX_DIMENSION)
	bias := make([]float64, MATRIX_DIMENSION)
	for i := range inputs {
		inputs[i] = make([]float64, MATRIX_DIMENSION)
		weights[i] = make([]float64, MATRIX_DIMENSION)
		for j := 0; j < len(weights); j++ {
			inputs[i][j] = rand.Float64()
			weights[i][j] = rand.Float64()
		}
	}
	nr := Layer{
		Inputs:  &inputs,
		Weights: &weights,
		Bias:    &bias,
	}

	startTime := time.Now()
	avgTimeForRuns := 0.0
	for i := 0; i < b.N; i++ {
		nr.MultiplyAndAddBias()
		avgTimeForRuns += float64(time.Since(startTime).Nanoseconds())
	}
	avgTimeForRuns /= float64(b.N)
	totalTime := time.Since(startTime)
	b.Logf("Total time: %v | Average time: %v", totalTime, avgTimeForRuns)
}
