package main

import (
	"fmt"
	"time"
)

var months = []string{
	"январь", "февраль", "март", "апрель", "май", "июнь",
	"июль", "август", "сентябрь", "октябрь", "ноябрь", "декабрь",
}

func currentDateString() string {
	t := time.Now()
	return fmt.Sprintf("%02d/%s/%d", t.Day(), months[t.Month()-1], t.Year())
}

func main() {
	fmt.Println(currentDateString())
}

func test() {
	fmt.Println(currentDateString())
}
