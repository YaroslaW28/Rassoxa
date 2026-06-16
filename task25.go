package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count, scanner.Err()
}

func main() {
	var filename string
	fmt.Scan(&filename)
	n, err := countLines(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(n)
}

func test() {
	n, err := countLines("test.txt")
	fmt.Println(n, err)
}
