package utils

import (
	"reflect"
	"testing"
)

func Test_Unit_MatrixDotProduct(t *testing.T) {
	twoByThreeMatrix := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}
	threeByTwoMatrix := [][]float64{
		{7, 8},
		{9, 10},
		{11, 12},
	}

	resultCaseOne, _ := MatrixDotProduct(twoByThreeMatrix, threeByTwoMatrix)
	expectedCaseOne := [][]float64{
		{58, 64},
		{139, 154},
	}
	if !reflect.DeepEqual(resultCaseOne, expectedCaseOne) {
		t.Errorf("Expected %v, got %v", expectedCaseOne, resultCaseOne)
	}

	resultCaseTwo, _ := MatrixDotProduct(threeByTwoMatrix, twoByThreeMatrix)
	expectedCaseTwo := [][]float64{
		{39, 54, 69},
		{49, 68, 87},
		{59, 82, 105},
	}
	if !reflect.DeepEqual(resultCaseTwo, expectedCaseTwo) {
		t.Errorf("Expected %v, got %v", expectedCaseTwo, resultCaseTwo)
	}
}

func Test_Unit_MatrixDotProduct_Error(t *testing.T) {
	t.Log("Testing MatrixDotProduct Error")
	matrix1 := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}
	matrix2 := [][]float64{
		{7, 8},
		{9, 10},
	}
	_, err := MatrixDotProduct(matrix1, matrix2)
	expectedErrorMessage := "the number of columns in matrixA must match the number of rows in matrixB: 3 != 2"
	if err == nil && err.Error() != expectedErrorMessage {
		t.Errorf("Expected error, got nil")
	}
}

func Test_Unit_MultiplyMatrixByScalar(t *testing.T) {
	t.Log("Testing MultiplyMatrixByScalar")
	matrix := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}
	scalar := 2
	result := MultiplyMatrixByScalar(matrix, float64(scalar))
	expected := [][]float64{
		{2, 4, 6},
		{8, 10, 12},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func Test_Unit_CreateMatrixFromVectors(t *testing.T) {
	t.Log("Testing CreateMatrixFromVectors")
	vectorA := []float64{1, 2, 3}
	vectorB := []float64{4, 5, 6}
	result := CreateMatrixFromVectors(vectorA, vectorB)
	expected := [][]float64{
		{4, 5, 6},
		{8, 10, 12},
		{12, 15, 18},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func Test_Unit_Flatten(t *testing.T) {
	t.Log("Testing Flatten")
	matrix := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}
	result := Flatten(matrix)
	expected := []float64{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func Test_Unit_RandomUniform(t *testing.T) {
	t.Log("Testing RandomUniform")
	result := RandomUniform(-0.5, 0.5, 2, 3)
	if len(result) != 2 || len(result[0]) != 3 {
		t.Errorf("Expected 2x3 matrix, got %dx%d", len(result), len(result[0]))
	}
}

func Test_Unit_SumArray(t *testing.T) {
	t.Log("Testing SumArray")
	result := SumArray([]float64{1, 2, 3, 4, 5})
	expected := 15.0
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func Test_Unit_SubtractArrays(t *testing.T) {
	t.Log("Testing SubtractArrays")
	result, _ := SubtractArrays([]float64{1, 2, 3, 4, 5}, []float64{1, 2, 3, 4, 5})
	expected := []float64{0, 0, 0, 0, 0}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func Test_Unit_SubtractArrays_Error(t *testing.T) {
	t.Log("Testing SubtractArrays Error")
	_, err := SubtractArrays([]float64{1, 2, 3}, []float64{1, 2, 3, 4, 5})
	expectedErrorMessage := "arrays must be the same length: 3 != 5"
	if err == nil && err.Error() != expectedErrorMessage {
		t.Errorf("Expected error, got nil")
	}
}

func Test_Unit_SumArrays(t *testing.T) {
	t.Log("Testing SumArrays")
	result, _ := SumArrays([]float64{1, 2, 3, 4, 5}, []float64{1, 2, 3, 4, 5})
	expected := []float64{2, 4, 6, 8, 10}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func Test_Unit_SumArrays_Error(t *testing.T) {
	t.Log("Testing SumArrays Error")
	_, err := SumArrays([]float64{1, 2, 3}, []float64{1, 2, 3, 4, 5})
	expectedErrorMessage := "arrays must be the same length: 3 != 5"
	if err == nil && err.Error() != expectedErrorMessage {
		t.Errorf("Expected error, got nil")
	}
}

func Test_Unit_SquareArray(t *testing.T) {
	t.Log("Testing SquareArray")
	result := SquareArray([]float64{1, 2, 3, 4, 5})
	expected := []float64{1, 4, 9, 16, 25}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func Test_Unit_MaxArrayPosition(t *testing.T) {
	t.Log("Testing MaxArrayPosition")
	position, maxVal := MaxArrayPosition([]float64{1, 2, 3, 4, 5})
	expectedPos := 4
	expectedMaxVal := 5.0
	if position != expectedPos || maxVal != expectedMaxVal {
		t.Errorf("Expected Position %v, got %v and Value %v but got %v", expectedPos, position, expectedMaxVal, maxVal)
	}
}
