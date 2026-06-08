package main

import (
	"fmt"
	"reflect"
)

func findThreeDigitNumbers(n int) []int {
	var result []int
	for i := 100; i <= 999; i++ {
		sum := (i / 100) + ((i / 10) % 10) + (i % 10)
		if sum == n {
			result = append(result, i)
		}
	}
	return result
}

func testt() {
	res := findThreeDigitNumbers(26)
	expected := []int{899, 989, 998}
	if reflect.DeepEqual(res, expected) {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed")
	}
}

func ras10() {
	test()
	fmt.Println(findThreeDigitNumbers(5))
}
