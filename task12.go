package main

import "fmt"

func countEven(arr []int) int {
	count := 0
	for _, v := range arr {
		if v%2 == 0 {
			count++
		}
	}
	return count
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	fmt.Println(countEven(arr))
}

func test122() {
	fmt.Println(countEven([]int{1, 2, 3, 4, 5, 6})) // 3
	fmt.Println(countEven([]int{1, 3, 5}))          // 0
}
