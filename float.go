package golib

import "math"

func Minf(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func Maxf(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func Clipf(value, lower, upper float32) float32 {
	if value < lower {
		return lower
	}
	if value > upper {
		return upper
	}
	return value
}

func Distancef(x0, y0, x1, y1 float32) float32 {
	dx := x1 - x0
	dy := y1 - y0
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}
