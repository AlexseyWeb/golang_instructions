package main

import (
	"fmt"
)

func xm(a string) {
	fmt.Println("Hello is another module ", a)
}

func xm_circle(a string) {
	for i := range a {
		fmt.Println(i)
	}
}
