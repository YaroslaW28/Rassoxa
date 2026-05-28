package main

import (
	"fmt"
	"math/rand"
)

func rass() {
	rows := 100
	cols := 200
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = rand.Intn(14) - 3
		}
	}
	defer func() {
		matrix = nil
		fmt.Println("Память успешно очищена сборщиком мусора")
	}()
	fmt.Printf("Элемент в [0][0]: %d\n", matrix[0][0])
}
