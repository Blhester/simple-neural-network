package connections

import (
	"simple-neural-network/activation"
	. "simple-neural-network/layer"
	"simple-neural-network/utils"
)

type DenseLayers struct {
	// The layer's neurons
	Seed                  uint64
	PreviousPassErrorRate []float64
	LearningRate          float64
	Inputs                [][]float64
}

type CostError struct {
	Delta                         float64
	DifferenceOfExpectedAndActual []float64
	AccuracyPercentage            float64
	DidPredictCorrectly           bool
}

func (d *DenseLayers) ForwardPass(hiddenNeuronCount, outputNeuronCount int, weightRange WeightRange) ([][]float64, error) {
	hiddenLayer := NewLayer(&d.Inputs, activation.ReLU)
	hiddenLayer.Init(hiddenNeuronCount, 1, weightRange)
	hiddenLayerOutput, err := runLayer(hiddenLayer)
	if err != nil {
		return nil, err
	}

	outputLayer := NewLayer(&hiddenLayerOutput, activation.Sigmoid)
	outputLayer.Init(outputNeuronCount, 1, weightRange)
	outputLayerOutput, err := runLayer(outputLayer)
	outputLayerOutput = d.BackwardsPropagate(outputLayerOutput)
	if err != nil {
		return nil, err
	}

	return outputLayerOutput, nil
}

func (d *DenseLayers) BackwardsPropagate(outputs [][]float64) [][]float64 {
	if d.PreviousPassErrorRate != nil {
		currentErrorPercent := utils.CreateMatrixFromVectors(outputs[0], d.PreviousPassErrorRate)
		return utils.MultiplyMatrixByScalar(currentErrorPercent, -d.LearningRate)
	}
	return outputs
}

func CalculateCostError(expectedOutput, actualOutput []float64) CostError {
	delta := 1 / 10 *
		utils.SumArray(
			utils.SquareArray(
				utils.SubtractArrays(actualOutput, expectedOutput)))
	differenceOfExpectedAndActual := utils.SubtractArrays(actualOutput, expectedOutput)
	accuracyPercentage := utils.SumArray(differenceOfExpectedAndActual) / float64(len(differenceOfExpectedAndActual))
	didPredictCorrectly := utils.MaxArrayPosition(actualOutput) == utils.MaxArrayPosition(expectedOutput)

	return CostError{
		delta,
		differenceOfExpectedAndActual,
		accuracyPercentage,
		didPredictCorrectly,
	}
}

func runLayer(layer *Layer) ([][]float64, error) {
	output, err := layer.MultiplyAndAddBias()
	if err != nil {
		return nil, err
	}
	return output, nil
}
