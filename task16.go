package main

import "fmt"

func sortStrings(arr []string) {
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
	arr := make([]string, n)
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	sortStrings(arr)
	fmt.Println(arr)
}

func test() {
	a := []string{"banana", "apple", "cherry", "date"}
	sortStrings(a)
	fmt.Println(a) // [apple banana cherry date]
}
