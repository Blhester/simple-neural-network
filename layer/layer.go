package layer

import (
	"simple-neural-network/activation"
	"simple-neural-network/utils"
)

type Layer struct {
	// The neuron's input
	Inputs *[][]float64
	// The neuron's weights
	Weights *[][]float64
	// The neuron's bias
	Bias       *[]float64
	Activation activation.FunctionType
}

func NewLayer(inputs *[][]float64, activation activation.FunctionType) *Layer {
	return &Layer{
		Inputs:     inputs,
		Activation: activation,
		Weights:    &[][]float64{},
		Bias:       &[]float64{},
	}
}
func (l *Layer) Init(numberOfNeurons int, bias float64, weightRange WeightRange) *Layer {
	if l.Inputs == nil {
		panic("Inputs is nil")
	}

	inputsElementsSize := len((*l.Inputs)[0])
	if inputsElementsSize == 0 {
		panic("Inputs length is 0")
	}

	*l.Weights = utils.RandomUniform(weightRange.Start, weightRange.End, inputsElementsSize, numberOfNeurons)

	*l.Bias = make([]float64, inputsElementsSize)
	for i := range *l.Bias {
		(*l.Bias)[i] = bias
	}
	return l
}

// Matrix multiplication and Bias addition
func (l *Layer) MultiplyAndAddBias() ([][]float64, error) {
	inputs := *l.Inputs
	weights := *l.Weights
	bias := *l.Bias
	if l.Activation == nil {
		l.Activation = activation.None
	}

	outputs, err := utils.MatrixDotProduct(inputs, weights)
	if err != nil {
		return nil, err
	}

	for i := range outputs {
		for j := range outputs[i] {
			outputs[i][j] += bias[j]
			activationResults := l.Activation(outputs[i][j])
			outputs[i][j] = activationResults
		}
	}

	return outputs, nil
}
