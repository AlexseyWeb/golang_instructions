package main

import (
	"fmt"
)

func main() {
	aString := []byte("A little string")
	for _, v := range aString {
		fmt.Println("String ", string(v))
	}
	my_rune := '&'
	fmt.Printf("Runa -> %x = %x", my_rune, my_rune)

}
