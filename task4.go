package main

import (
	"fmt"
	"math"
)

func vectorLength(v []float64) float64 {
	var sum float64
	for _, coord := range v {
		sum += coord * coord
	}
	return math.Sqrt(sum)
}

func test() {
	v1 := []float64{3.0, 4.0, 0.0}
	res1 := vectorLength(v1)
	if res1 == 5.0 {
		fmt.Println("Тест 1 пройден (3D вектор)")
	} else {
		fmt.Printf("Тест 1 ПРОВАЛЕН: ожидалось 5, получено %f\n", res1)
	}

	v2 := []float64{1.0, 1.0}
	res2 := vectorLength(v2)
	expected2 := math.Sqrt(2)
	if math.Abs(res2-expected2) < 1e-9 {
		fmt.Println("Тест 2 пройден (2D вектор)")
	} else {
		fmt.Printf("Тест 2 ПРОВАЛЕН: ожидалось %f, получено %f\n", expected2, res2)
	}
}

func rassss() {
	fmt.Println("--- Запуск тестов ---")
	test()
	fmt.Println("---------------------")
	myVector := []float64{1.0, 2.0, 2.0, 4.0, 5.0}
	length := vectorLength(myVector)
	fmt.Printf("Длина вектора %v равна: %.4f\n", myVector, length)
}
