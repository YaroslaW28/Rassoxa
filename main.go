package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rows := 100
	cols := 200
	matrix := make([][]float64, rows)

	for i := range matrix {
		matrix[i] = make([]float64, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = rand.Float64()*2 - 1
		}
	}
	fmt.Printf("Пример значения: %f\n", matrix[0][0])
}
