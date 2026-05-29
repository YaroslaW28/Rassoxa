package main

import (
	"fmt"
)

func rasss() {
	machineZero := 1.0
	for machineZero/2.0 > 0.0 {
		machineZero /= 2.0
	}
	machineEpsilon := 1.0
	for 1.0+machineEpsilon/2.0 > 1.0 {
		machineEpsilon /= 2.0
	}
	fmt.Printf("Машинный ноль для float64:    %e\n", machineZero)
	fmt.Printf("Машинный эпсилон для float64: %e\n", machineEpsilon)
}
