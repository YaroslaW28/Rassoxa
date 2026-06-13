package main

import (
	"fmt"
	"time"
)

func compareDates(d1, d2 string) int {
	layout := "02/01/2006"
	t1, _ := time.Parse(layout, d1)
	t2, _ := time.Parse(layout, d2)
	if t1.Before(t2) {
		return -1
	} else if t1.After(t2) {
		return 1
	}
	return 0
}

func main() {
	var d1, d2 string
	fmt.Scan(&d1, &d2)
	fmt.Println(compareDates(d1, d2))
}

func test() {
	fmt.Println(compareDates("01/01/2020", "31/12/2019")) // 1
	fmt.Println(compareDates("01/01/2020", "01/01/2020")) // 0
	fmt.Println(compareDates("01/01/2019", "01/01/2020")) // -1
}
