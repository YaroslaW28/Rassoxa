package main

import (
	"fmt"
)

func findLuckyTickets() []int {
	var result []int
	for i := 0; i <= 999999; i++ {
		firstHalf := i / 1000
		secondHalf := i % 1000

		sum1 := (firstHalf / 100) + ((firstHalf / 10) % 10) + (firstHalf % 10)
		sum2 := (secondHalf / 100) + ((secondHalf / 10) % 10) + (secondHalf % 10)

		if sum1 == sum2 {
			result = append(result, i)
		}
	}
	return result
}

func test3() {
	res := findLuckyTickets()
	if len(res) == 55252 {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed")
	}
}

func rasssss() {
	test()
	tickets := findLuckyTickets()
	fmt.Printf("Total lucky tickets: %d\n", len(tickets))
	fmt.Printf("First 5 lucky tickets: %v\n", tickets[:5])
}
