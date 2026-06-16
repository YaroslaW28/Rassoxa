package main

import (
	"bytes"
	"fmt"
	"os"
)

func filesEqual(f1, f2 string) (bool, error) {
	b1, err := os.ReadFile(f1)
	if err != nil {
		return false, err
	}
	b2, err := os.ReadFile(f2)
	if err != nil {
		return false, err
	}
	return bytes.Equal(b1, b2), nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: program file1 file2")
		return
	}
	eq, err := filesEqual(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(eq)
}

func test() {
	eq, err := filesEqual("a.txt", "b.txt")
	fmt.Println(eq, err)
}
