package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	if len(os.Args) != 2 {
		fmt.Println("Usage: dates parse_string")
		return
	}
	dateString := os.Args[1]
	d, err := time.Parse("02 January 2025", dateString)
	if err == nil {
		fmt.Println("Full: ", d)
		fmt.Println("Time: ", d.Day(), d.Month(), d.Year())
	}

	d, err = time.Parse("02 January 2025 15:04", dateString)
	if err == nil {
		fmt.Println("Full ", d)
		fmt.Println("Date ", d.Day(), d.Month(), d.Year())
		fmt.Println("Time: ", d.Hour(), d.Minute())

	}
	fmt.Println(start)
}
