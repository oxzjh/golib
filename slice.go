package golib

import (
	"math"
	"math/rand"
)

func Sum(arr []float64) float64 {
	val := 0.0
	for _, v := range arr {
		val += v
	}
	return val
}

func Weight(arr []float64) int {
	val := Sum(arr) * rand.Float64()
	t := 0.0
	for i, v := range arr {
		t += v
		if val < t {
			return i
		}
	}
	return 0
}

func Average(arr []float64) float64 {
	return Sum(arr) / float64(len(arr))
}

func Softmax(arr []float64) []float64 {
	result := make([]float64, len(arr))
	maxVal := arr[0]
	for _, v := range arr {
		if v > maxVal {
			maxVal = v
		}
	}
	var sum float64
	for i, v := range arr {
		result[i] = math.Exp(v - maxVal)
		sum += result[i]
	}
	for i, v := range result {
		result[i] = v / sum
	}
	return result
}
