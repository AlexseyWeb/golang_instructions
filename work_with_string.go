package main

import (
	"fmt"
	"strings"
)

func infinite_numbers(numbers ...int) {
	for _, number := range numbers {
		fmt.Println(number)
	}
}

func main() {
	names := "This+Is+Very+Long+A+String!"
	var numb = []int{7, 8, 9, 10, 11, 12}
	fmt.Print("|")
	for _, name := range strings.Split(names, "+") {
		fmt.Println("%s|", name)
	}
	infinite_numbers(1, 2, 3, 4, 5, 6)
	infinite_numbers(numb...)

}
