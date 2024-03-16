package utils

import (
	"reflect"
	"testing"
)

func TestMatrixDotProduct(t *testing.T) {
	t.Log("Testing MatrixDotProduct")
	matrix1 := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}
	matrix2 := [][]float64{
		{7, 8},
		{9, 10},
		{11, 12},
	}
	result, _ := MatrixDotProduct(matrix1, matrix2)
	expected := [][]float64{
		{58, 64},
		{139, 154},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
