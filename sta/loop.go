package main

import (
	"fmt"
)

func numbers_range() {
	for l := 1; l < 100; l++ {
		sum := l + l
		fmt.Printf("Run of loop  %d ->", sum)
	}
}

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i * i)
	}
	numbers_range()

	aSlice := []int{-1, 2, 1, -1, 2, -2}
	for i, v := range aSlice {
		fmt.Println("index ", i, " value ", v)
	}
}
