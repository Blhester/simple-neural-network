package layer

import "strconv"

type BaseLayer interface {
	GetLayer() *BasicLayer
	Forward() error
	Backward(outputGradient [][]float64, learningRate float64) error
}

type BasicLayer struct {
	Inputs  []float64
	Weights [][]float64
	Bias    []float64
	Outputs []float64
}

func (l *BasicLayer) GetLayer() *BasicLayer {
	return l
}

func (l *BasicLayer) Forward() error {
	panic("Not implemented")
}

func (l *BasicLayer) Backward(outputGradient [][]float64, learningRate float64) error {
	panic("Not implemented")
}

func (l *BasicLayer) VerifyLayerSetup() {
	if len(l.Inputs) == 0 {
		panic("Inputs length is 0")
	}
	if len(l.Weights) == 0 {
		panic("Weights length is 0")
	}
	if len(l.Bias) == 0 {
		panic("Bias length is 0")
	}
	if len(l.Outputs) == 0 {
		panic("Outputs length is 0")
	}

	if len(l.Inputs) != len(l.Weights[0]) {
		panic("Inputs length " + strconv.Itoa(len(l.Inputs)) + " must match the length of the Weights " + strconv.Itoa(len(l.Weights[0])))
	}

	if len(l.Weights) != len(l.Bias) {
		panic("Weights length " + strconv.Itoa(len(l.Weights)) + " must match the length of the Bias " + strconv.Itoa(len(l.Bias)))
	}

}
