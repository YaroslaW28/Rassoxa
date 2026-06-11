package main

import "fmt"

func derivative(f func(float64) float64, x float64) float64 {
	h := 1e-9
	return (f(x+h) - f(x-h)) / (2 * h)
}

func main() {
	f := func(x float64) float64 {
		return x*x*x - 2*x*x + x - 5
	}
	var x float64
	fmt.Scan(&x)
	fmt.Printf("f'(%.4f) = %.6f\n", x, derivative(f, x))
}

func test() {
	f := func(x float64) float64 { return x * x }
	fmt.Printf("%.6f\n", derivative(f, 3.0)) // ~6.000000
}
