package layer

type DenseLayer struct {
	Inputs  []float64
	Weights [][]float64
	Bias    []float64
	Outputs []float64
}

func (l *DenseLayer) GetLayer() *BasicLayer {
	return (*BasicLayer)(l)
}

func (l *DenseLayer) Forward() error {
	panic("Not implemented")
}

func (l *DenseLayer) Backward(outputGradient [][]float64, learningRate float64) error {
	panic("Not implemented")
}
