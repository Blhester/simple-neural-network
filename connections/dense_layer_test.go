package connections

import (
	"reflect"
	"testing"
)

func TestLayer(t *testing.T) {
	t.Log("Testing DenseLayers")
	layer1 := DenseLayers{}
	layer1.Init(4, 5, 0)
	inputs := [][]float64{
		{1, 2, 3, 2.5},
		{2.0, 5.0, -1.0, 2.0},
		{-1.5, 2.7, 3.3, -0.8},
	}
	output, _ := layer1.Forward(inputs)
	t.Log("Layer1", output)

	layer2 := DenseLayers{}
	layer2.Init(5, 2, 0)
	output, _ = layer2.Forward(output)
	t.Log("Layer2", output)
}

func BenchmarkLayer(b *testing.B) {
	b.Logf("Running ( %d ) times", b.N)
	layer1 := DenseLayers{}
	layer1.Init(4, 5, 0)
	inputs := [][]float64{
		{1, 2, 3, 2.5},
		{2.0, 5.0, -1.0, 2.0},
		{-1.5, 2.7, 3.3, -0.8},
	}
	for i := 0; i < b.N; i++ {
		layer1.Forward(inputs)
	}
}

func TestLayerPass_Forward(t *testing.T) {
	inputs := [][]float64{
		{0, 1, 1, 1},
	}

	layer := DenseLayers{}
	layer.Init(4, 3, 0)
	output, _ := layer.Forward(inputs)
	layerSecond := DenseLayers{}
	layerSecond.Init(3, 2, 0)
	output, _ = layerSecond.Forward(output)
	layerLast := DenseLayers{}
	layerLast.Init(2, 1, 0)
	output, _ = layerLast.Forward(output)
	t.Run("LayerPass_Forward", func(t *testing.T) {
		if !reflect.DeepEqual(output, [][]float64{{0, 1}}) {
			t.Errorf("Expected %v, got %v", [][]float64{{0, 1}}, output)
		}
	})
}
