package main

import "fmt"

func allDigitsDifferent(n int) bool {
	if n < 0 {
		n = -n
	}
	digits := [10]bool{}
	for n > 0 {
		d := n % 10
		if digits[d] {
			return false
		}
		digits[d] = true
		n /= 10
	}
	return true
}

func findNumbers(N int) []int {
	var result []int
	for i := 1; i <= N; i++ {
		if allDigitsDifferent(i) {
			result = append(result, i)
		}
	}
	return result
}

func zadanie11() {
	var N int
	fmt.Scan(&N)
	fmt.Println(findNumbers(N))
}

func testtt() {
	fmt.Println(allDigitsDifferent(123)) // true
	fmt.Println(allDigitsDifferent(112)) // false
	fmt.Println(findNumbers(20))
}
