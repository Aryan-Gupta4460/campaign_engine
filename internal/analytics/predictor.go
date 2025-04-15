package analytics

import "math/rand"

func PredictCPC() float64 {
	return 0.5 + rand.Float64()
}

func PredictCVR() float64 {
	return 0.01 + rand.Float64()/100
}
