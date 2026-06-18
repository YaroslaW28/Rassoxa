package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 24
)

func clear() {
	fmt.Print("\033[H\033[2J")
}

func drawPixel(x, y int) {
	fmt.Printf("\033[%d;%dH*", y+1, x+1)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	x, y := width/2, height/2
	steps := 500
	for i := 0; i < steps; i++ {
		clear()
		drawPixel(x, y)
		fmt.Printf("\033[%d;1H", height+1)
		time.Sleep(50 * time.Millisecond)
		dx := (rand.Intn(3) - 1) * 2
		dy := (rand.Intn(3) - 1) * 2
		x += dx
		y += dy
		if x < 0 {
			x = 0
		}
		if x >= width {
			x = width - 1
		}
		if y < 0 {
			y = 0
		}
		if y >= height {
			y = height - 1
		}
	}
}
