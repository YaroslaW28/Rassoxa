package main

import "fmt"

func hasZeroInEachRowAndCol(matrix [][]int) bool {
	rows := len(matrix)
	if rows == 0 {
		return false
	}
	cols := len(matrix[0])
	for i := 0; i < rows; i++ {
		found := false
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	for j := 0; j < cols; j++ {
		found := false
		for i := 0; i < rows; i++ {
			if matrix[i][j] == 0 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func main() {
	var rows, cols int
	fmt.Scan(&rows, &cols)
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
		for j := range matrix[i] {
			fmt.Scan(&matrix[i][j])
		}
	}
	fmt.Println(hasZeroInEachRowAndCol(matrix))
}

func test() {
	m := [][]int{{0, 1, 2}, {3, 0, 5}, {6, 7, 0}}
	fmt.Println(hasZeroInEachRowAndCol(m)) // true
	m2 := [][]int{{1, 2}, {3, 4}}
	fmt.Println(hasZeroInEachRowAndCol(m2)) // false
}
