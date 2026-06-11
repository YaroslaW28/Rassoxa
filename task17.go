package main

import "fmt"

func lastIndex(s string, c byte) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			return i
		}
	}
	return -1
}

func main() {
	var s string
	var c string
	fmt.Scan(&s, &c)
	fmt.Println(lastIndex(s, c[0]))
}

func test() {
	fmt.Println(lastIndex("hello world", 'l')) // 9
	fmt.Println(lastIndex("hello", 'z'))       // -1
}
