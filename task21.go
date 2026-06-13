package main

import "fmt"

// F(n+2) = 2*F(n+1) - F(n), F1=2, F0=1
func F(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 2
	}
	return 2*F(n-1) - F(n-2)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(F(n))
}

func test() {
	for i := 0; i <= 6; i++ {
		fmt.Printf("F(%d) = %d\n", i, F(i))
	}
	// 1, 2, 3, 4, 5, 6, 7
}
