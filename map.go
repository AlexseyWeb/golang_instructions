package main

import (
	"fmt"
)

func printMap(person map[string]string) {
	for k, v := range person {
		fmt.Println("Key -> ", k, " Value -> ", v)
	}
	fmt.Println("--------------------------------")
}

func main() {
	person := make(map[string]string)

	person["name"] = "Alexsey"
	person["age"] = "33"
	person["sallary"] = "30_000"

	printMap(person)

	html_tag := make(map[string]string)
	html_tag["h1"] = "Header"
	html_tag["h2"] = "Header 2"
	html_tag["body"] = "Body"

	printMap(html_tag)
}
