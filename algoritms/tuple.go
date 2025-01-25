package main

import (
	"fmt"
)

func createTuple(first int) (int, int) {
	return first * first, first * first * first
}

func main() {
	square, triaple := createTuple(8)

	fmt.Println("Square = ", square, "Triaple = ", triaple)
}
