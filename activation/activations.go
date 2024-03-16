package activation

import (
	"math"
)

type FunctionType func(float64) float64
type ActivationFunction struct {
	Fn FunctionType
}

func Sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func SigmoidPrime(x float64) float64 {
	return Sigmoid(x) * (1 - Sigmoid(x))
}

func ReLU(x float64) float64 {
	if x > 0 {
		return x
	}
	return 0
}

func ReLUPrime(x float64) float64 {
	if x > 0 {
		return 1
	}
	return 0
}

func SoftMax(logits []float64) []float64 {
	maxLogit := logits[0]
	for _, logit := range logits {
		if logit > maxLogit {
			maxLogit = logit
		}
	}
	sumExp := 0.0
	for _, logit := range logits {
		sumExp += math.Exp(logit - maxLogit) // Subtracting maxLogit for numerical stability
	}

	probabilities := make([]float64, len(logits))
	for i, logit := range logits {
		probabilities[i] = math.Exp(logit-maxLogit) / sumExp
	}

	return probabilities
}

func None(x float64) float64 {
	return x
}
