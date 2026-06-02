package main

import (
	"fmt"
	"reflect"
)

func multiplyMatrices(a, b [][]int) [][]int {
	if len(a) == 0 || len(b) == 0 || len(a[0]) != len(b) {
		return nil
	}

	m := len(a)
	n := len(a[0])
	k := len(b[0])

	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, k)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < k; j++ {
			for l := 0; l < n; l++ {
				result[i][j] += a[i][l] * b[l][j]
			}
		}
	}
	return result
}

func test5() {
	a := [][]int{{1, 2}, {3, 4}}
	b := [][]int{{5, 6}, {7, 8}}
	expected := [][]int{{19, 22}, {43, 50}}
	res := multiplyMatrices(a, b)
	if reflect.DeepEqual(res, expected) {
		fmt.Println("Test passed")
	} else {
		fmt.Println("Test failed")
	}
}

func rasssssssss() {
	test()
	matrixA := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	matrixB := [][]int{
		{7, 8},
		{9, 1},
		{2, 3},
	}
	fmt.Println(multiplyMatrices(matrixA, matrixB))
}
