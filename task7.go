package main

import (
	"fmt"
	"reflect"
)

func findPalindromes(n int) []int {
	var result []int
	for i := 1; i <= n; i++ {
		temp := i
		reversed := 0
		for temp > 0 {
			reversed = reversed*10 + temp%10
			temp /= 10
		}
		if i == reversed {
			result = append(result, i)
		}
	}
	return result
}

func test2() {
	res := findPalindromes(20)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11}
	if reflect.DeepEqual(res, expected) {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed")
	}
}

func rasssssss() {
	test()
	fmt.Println(findPalindromes(150))
}
