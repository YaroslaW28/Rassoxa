package main

import "fmt"

func primeFactors(n int) []int {
	if n <= 1 {
		return nil
	}
	if n%2 == 0 {
		return append([]int{2}, primeFactors(n/2)...)
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return append([]int{i}, primeFactors(n/i)...)
		}
	}
	return []int{n}
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(primeFactors(n))
}

func test() {
	fmt.Println(primeFactors(60))  // [2 2 3 5]
	fmt.Println(primeFactors(100)) // [2 2 5 5]
	fmt.Println(primeFactors(13))  // [13]
}
