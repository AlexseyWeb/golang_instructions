package main

import "fmt"

func main() {
	b := make([]byte, 12)
	fmt.Println("Byte slice -> ", b)

	b = []byte("Byte slice:¼ ")
	fmt.Println(b)
	fmt.Println(string(b))
}
