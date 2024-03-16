package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func MultiplyMatrixByScalar(matrix [][]float64, scalar float64) [][]float64 {
	result := make([][]float64, len(matrix))
	for i := range matrix {
		result[i] = make([]float64, len(matrix[i]))
		for j := range matrix[i] {
			result[i][j] = matrix[i][j] * scalar
		}
	}
	return result
}

func CreateMatrixFromVectors(vectorA []float64, vectorB []float64) [][]float64 {
	result := make([][]float64, len(vectorA))
	for i := range vectorA {
		result[i] = make([]float64, len(vectorB))
		for j := range vectorB {
			result[i][j] = vectorA[i] * vectorB[j]
		}
	}
	return result
}

func MatrixDotProduct(matrixA [][]float64, matrixB [][]float64) ([][]float64, error) {
	if len(matrixA[0]) != len(matrixB) {
		return nil, fmt.Errorf("the number of columns in matrixA must match the number of rows in matrixB: %d != %d", len(matrixA[0]), len(matrixB))
	}

	result := make([][]float64, len(matrixA))
	for i := 0; i < len(matrixA); i++ {
		result[i] = make([]float64, len(matrixB[0]))
		for j := 0; j < len(matrixB[0]); j++ {
			for k := 0; k < len(matrixA[0]); k++ {
				result[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}

	return result, nil
}

func DotProduct(vectorA []float64, vectorB []float64) (float64, error) {
	if len(vectorA) != len(vectorB) {
		panic("Vectors must have the same length")
	}
	var result float64
	for i := 0; i < len(vectorA); i++ {
		result += vectorA[i] * vectorB[i]
	}
	return result, nil
}

func ReduceOutputByHalf(outputs [][]float64) [][]float64 {
	reducedOutputs := make([][]float64, len(outputs))
	for i, output := range outputs {
		for j := 0; j < len(output); j += 2 { // Skip by 2 to reduce by half, adjust as necessary
			reducedOutputs[i] = append(reducedOutputs[i], output[j])
		}
	}
	return reducedOutputs
}

func RemoveZeroSumArrays(outputs [][]float64) [][]float64 {
	var filteredOutputs [][]float64
	for _, output := range outputs {
		sum := 0.0
		for _, value := range output {
			sum += value
		}
		if sum != 0 {
			filteredOutputs = append(filteredOutputs, output)
		}
	}
	return filteredOutputs
}

func Flatten(matrix [][]float64) []float64 {
	var result []float64
	for _, row := range matrix {
		result = append(result, row...)
	}
	return result
}

// RandomUniform generates a 2D slice of size (rows x cols) filled with random floats between startVal and endVal.
func RandomUniform(startVal, endVal float64, rows, cols int) [][]float64 {
	// Seed the random number generator to ensure different results on each run
	rand.Seed(time.Now().UnixNano())

	// Initialize the 2D slice
	result := make([][]float64, rows)
	for i := range result {
		result[i] = make([]float64, cols)
		for j := range result[i] {
			// Calculate the random value
			result[i][j] = startVal + rand.Float64()*(endVal-startVal)
		}
	}
	return result
}

func SubtractArrays(arrayA, arrayB []float64) []float64 {
	result := make([]float64, len(arrayA))
	for i := range arrayA {
		result[i] = arrayA[i] - arrayB[i]
	}
	return result
}

func SquareArray(array []float64) []float64 {
	result := make([]float64, len(array))
	for i := range array {
		result[i] = array[i] * array[i]
	}
	return result
}

func SumArray(array []float64) float64 {
	var result float64
	for i := range array {
		result += array[i]
	}
	return result
}

func MaxArrayValue(array []float64) float64 {
	var result float64
	for i := range array {
		if array[i] > result {
			result = array[i]
		}
	}
	return result
}

func MaxArrayPosition(array []float64) int {
	var result float64
	var position int
	for i := range array {
		if array[i] > result {
			result = array[i]
			position = i
		}
	}
	return position
}
