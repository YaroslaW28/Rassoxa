package main

import "fmt"

func trim(s string) string {
	start := 0
	end := len(s) - 1
	for start <= end && s[start] == ' ' {
		start++
	}
	for end >= start && s[end] == ' ' {
		end--
	}
	return s[start : end+1]
}

func main() {
	var s string
	fmt.Scanf("%[^\n]s", &s)
	fmt.Printf("'%s'\n", trim(s))
}

func test() {
	fmt.Printf("'%s'\n", trim("  hello world  ")) // 'hello world'
	fmt.Printf("'%s'\n", trim("   "))             // ''
	fmt.Printf("'%s'\n", trim("no spaces"))       // 'no spaces'
}
