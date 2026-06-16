package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
)

func writeRandomInts(filename string, n int) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	for i := 0; i < n; i++ {
		val := int32(rand.Intn(2)*2 - 1) // -1 или 1
		binary.Write(f, binary.LittleEndian, val)
	}
	return nil
}

func readAndSum(filename string, n int) (int64, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	data := make([]int32, n)
	binary.Read(f, binary.LittleEndian, data)
	var sum int64
	for _, v := range data {
		sum += int64(v)
	}
	return sum, nil
}

func main() {
	const n = 10000
	writeRandomInts("data.bin", n)
	sum, _ := readAndSum("data.bin", n)
	fmt.Println("Sum:", sum)
}
