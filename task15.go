package main

import "fmt"

func bubbleSort(arr []float64) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]float64, n)
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	bubbleSort(arr)
	fmt.Println(arr)
}

func test() {
	a := []float64{3.5, 1.2, 4.1, 1.5, 9.2, 2.6}
	bubbleSort(a)
	fmt.Println(a) // [1.2 1.5 2.6 3.5 4.1 9.2]
}
