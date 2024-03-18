package connections

import (
	"simple-neural-network/activation"
	. "simple-neural-network/layer"
	"simple-neural-network/utils"
)

type DenseLayers struct {
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
	if err != nil {
		return nil, err
	}
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
	differenceOfExpectedAndActual, err := utils.SubtractArrays(actualOutput, expectedOutput)
	if err != nil {
		panic(err)
	}
	squaredDifferenceOfExpectedAndActual := utils.SquareArray(differenceOfExpectedAndActual)
	summedSquaredDifferenceOfExpectedAndActual := utils.SumArray(squaredDifferenceOfExpectedAndActual)
	delta := 1 / 10 * summedSquaredDifferenceOfExpectedAndActual

	accuracyPercentage := utils.SumArray(differenceOfExpectedAndActual) / float64(len(differenceOfExpectedAndActual))
	maxActualValuePosition, _ := utils.MaxArrayPosition(actualOutput)
	maxExpectedValuePosition, _ := utils.MaxArrayPosition(expectedOutput)
	didPredictCorrectly := maxActualValuePosition == maxExpectedValuePosition

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
