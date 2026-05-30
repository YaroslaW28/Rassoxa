package main

import (
	"fmt"
	"reflect"
)

func findPrimes(n int) []int {
	if n < 2 {
		return []int{}
	}
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for p := 2; p*p <= n; p++ {
		if isPrime[p] {
			for i := p * p; i <= n; i += p {
				isPrime[i] = false
			}
		}
	}

	var primes []int
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
func test1() {
	res1 := findPrimes(10)
	expected1 := []int{2, 3, 5, 7}
	if reflect.DeepEqual(res1, expected1) {
		fmt.Println("Тест 1 пройден (N = 10)")
	} else {
		fmt.Printf("Тест 1 ПРОВАЛЕН: ожидалось %v, получено %v\n", expected1, res1)
	}
	res2 := findPrimes(1)
	if len(res2) == 0 {
		fmt.Println("Тест 2 пройден (N = 1)")
	} else {
		fmt.Printf("Тест 2 ПРОВАЛЕН: ожидался пустой слайс, получено %v\n", res2)
	}
}

func rasssss() {
	fmt.Println("--- Запуск тестов ---")
	test()
	fmt.Println("---------------------")
	n := 50
	primes := findPrimes(n)
	fmt.Printf("Простые числа до %d: %v\n", n, primes)
}
