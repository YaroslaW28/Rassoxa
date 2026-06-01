package main

import (
	"fmt"
	"math"
)

func targetFunction(x float64) float64 {
	return -(x * x) + 4*x + 5
}

func findMax(f func(float64) float64, a, b float64) float64 {
	step := 0.001
	maxVal := f(a)
	for x := a; x <= b; x += step {
		val := f(x)
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func test4() {
	res := findMax(targetFunction, 0.0, 4.0)
	if math.Abs(res-9.0) < 0.01 {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed")
	}
}

func rassssssss() {
	test()
	maxVal := findMax(targetFunction, -1.0, 5.0)
	fmt.Printf("Maximum value: %.4f\n", maxVal)
}
