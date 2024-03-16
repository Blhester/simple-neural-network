package layer

type WeightRange struct {
	Start float64
	End   float64
}

func NewRange(start, end float64) WeightRange {
	return WeightRange{start, end}
}
