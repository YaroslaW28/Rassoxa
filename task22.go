package main

import (
	"fmt"
	"time"
)

func measureInterval() time.Duration {
	fmt.Println("Press Enter...")
	fmt.Scanln()
	t1 := time.Now()
	fmt.Println("Press Enter again...")
	fmt.Scanln()
	t2 := time.Now()
	return t2.Sub(t1)
}

func main() {
	d := measureInterval()
	fmt.Printf("Interval: %v\n", d)
}
