package vector

import (
	"math"
)

func Normalize(amount, coef float32) float32 {
	return clamp(amount / coef)
}

func clamp(amount float32) float32 {
	if amount > 1.0 {
		return 1.0
	}
	if amount < 0.0 {
		return 0.0
	}
	return amount
}

// EuclidianDistance
//
//	Calculate the euclidian distance for a vector in the format
//		sqrt([[q1 - r2]² + [q2- r3]² + , ... [qn - rn]²])
func EuclidianDistance(vector1, vector2 []float32) float32 {
	var sum float32 = 0.0

	for i := range vector1 {
		d := vector1[i] - vector2[i]
		sum += (d * d)
	}

	return float32(math.Sqrt(float64(sum)))
}
