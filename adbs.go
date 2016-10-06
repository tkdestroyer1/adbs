package adbs

import (
	calculus "github.com/TheDemx27/calculus"
	//"math"
)

func CalcUpper(lower float64, expr string, k float64) float64 {
	f := calculus.NewFunc(expr)
	for uppr := lower; f.AntiDiff(lower, uppr) != k; uppr += float64(0.000001) {
		if f.AntiDiff(lower, uppr)-k < 0.000001 {
			return uppr
		} else if uppr == uppr+float64(0.000001) {
			break
		}
	}
	return lower
}