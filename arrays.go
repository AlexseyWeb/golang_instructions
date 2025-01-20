package main

import (
	"fmt"
)

func main() {
	aslice := []float64{}
	fmt.Println(aslice, len(aslice), cap(aslice))

	//added elements
	aslice = append(aslice, 1234.56)
	aslice = append(aslice, -34.0)
	aslice = append(aslice, -12.0)
	aslice = append(aslice, -4.0)
	fmt.Println(aslice, "with length ", len(aslice))

	t := make([]int, 4)
	t[0] = -1
	t[1] = -2
	t[2] = -3
	t[3] = -4

	t = append(t, -5)
	fmt.Println(t)
	fmt.Println("First three elements -> ", aslice[:3])

	two2D := [][]int{{1, 2, 3}, {4, 5, 6}}
	for _, i := range two2D {
		for _, k := range i {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}

	make2D := make([][]int, 2)
	fmt.Println(make2D)
	make2D[0] = []int{1, 2, 3, 4, 5, 6}
	make2D[1] = []int{0, 1, 0, 1, 0, 1}
	fmt.Println(make2D)
}
