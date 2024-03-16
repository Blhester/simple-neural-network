package main

import (
	"os"
	. "simple-neural-network/connections"
	. "simple-neural-network/file_readers/images"
	. "simple-neural-network/layer"
	"simple-neural-network/utils"
	"strconv"
	"testing"
	"time"
)

func TestPlayground(t *testing.T) {
	totalTimeStart := time.Now()
	totalNumberOfImages := 0
	dirsInImages, err := os.ReadDir("tmp/images/")
	if err != nil {
		t.Errorf("No images found: %v", err)
	}
	accuracyRate := make([]float64, 0)
	for _, dir := range dirsInImages {
		expectedNum, err := strconv.Atoi(dir.Name())
		if err != nil {
			t.Errorf("Error: %v", err)
			return
		}
		fileDir := "tmp/images/" + dir.Name() + "/"
		imageBytesArray, err := GetAllImagesAsBytes(fileDir)
		if err != nil {
			t.Errorf("Error: %v", err)
			return
		}
		numberOfImages := len(imageBytesArray)

		scoreMap := make(map[int]int)

		startTime := time.Now()
		errorRate := make([]float64, 10)
		for i := range imageBytesArray {
			inputs := [][]float64{imageBytesArray[i]}

			denseLayers := DenseLayers{
				Inputs:                inputs,
				PreviousPassErrorRate: errorRate,
				LearningRate:          0.01,
			}
			results, err := denseLayers.ForwardPass(20, 10, NewRange(-0.5, 0.5))
			if err != nil {
				t.Errorf("Error: %v", err)
				return
			}

			costError := CalculateCostError(getExpectedOutput(expectedNum), results[0])
			errorRate = costError.DifferenceOfExpectedAndActual
			if costError.DidPredictCorrectly {
				scoreMap[expectedNum]++
			} else {
				scoreMap[utils.MaxArrayPosition(results[0])]++
			}
		}
		accuracyPercent := float64(scoreMap[expectedNum]) / float64(numberOfImages)
		accuracyRate = append(accuracyRate, accuracyPercent)

		elapsedTime := time.Since(startTime)
		t.Logf("Score Map: %v\n", scoreMap)
		t.Logf("Accuracy percent: %v\n", accuracyPercent)
		t.Logf("Processed %v images in - Elapsed Time: %v\n", numberOfImages, elapsedTime)
		t.Logf("Final Error Rate Weights: %v\n", errorRate)
		totalNumberOfImages += numberOfImages
	}
	totalTimeElapsed := time.Since(totalTimeStart)
	t.Logf("Total Number of Images Processed: %v\n", totalNumberOfImages)
	t.Logf("Total Elapsed Time: %v\n", totalTimeElapsed)

	t.Logf("Success Rate: %v\n", accuracyRate)

	totalAccuracyPercent := utils.SumArray(accuracyRate) / float64(len(accuracyRate))
	t.Logf("Total Accuracy Rate: %v\n", totalAccuracyPercent)
	if totalAccuracyPercent < 0.9 {
		t.Errorf("Failed to reach 90 percent accuracy")
	}
}

func getExpectedOutput(num int) []float64 {
	expectedOutput := make([]float64, 10)
	expectedOutput[num] = 1
	return expectedOutput
}
