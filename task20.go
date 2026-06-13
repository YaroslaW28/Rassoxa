package main

import (
	"fmt"
	"math"
)

func findZero(f func(float64) float64, a, b, eps float64) float64 {
	if math.Abs(b-a) < eps {
		return (a + b) / 2
	}
	mid := (a + b) / 2
	if f(mid) == 0 {
		return mid
	}
	if f(a)*f(mid) < 0 {
		return findZero(f, a, mid, eps)
	}
	return findZero(f, mid, b, eps)
}

func main() {
	f := func(x float64) float64 { return x*x*x - x - 2 }
	fmt.Printf("%.8f\n", findZero(f, 1.0, 2.0, 1e-9))
}

func test() {
	f := func(x float64) float64 { return x*x - 2 }
	fmt.Printf("sqrt(2) ≈ %.8f\n", findZero(f, 1.0, 2.0, 1e-9)) // ~1.41421356
}
