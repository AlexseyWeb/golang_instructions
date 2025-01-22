package main

import (
	"fmt"
	"math/rand"
	"os"
	"path"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}

func search(key string) *Entry {
	for i, v := range data {
		if v.Surname == key {
			return &data[i]
		}
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func getString() {
	//TODO
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func populate(n int, s []Entry) {
	for i := 0; i < n; i++ {
		//name := getString(4)
		//surname := getString(5)
		//n := strconv.Itoa(random(100, 199))
		//data = append(data, Entry{name, surname, n})
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	data = append(data, Entry{"Alexsey", "Gusakov", "89401002222"})
	data = append(data, Entry{"Sergey", "Frost", "89803332323"})
	data = append(data, Entry{"Misha", "Crown", "89102223843"})

	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		result := search(arguments[2])
		if result == nil {
			fmt.Println("Entry not found:", arguments[2])
			return
		}
		fmt.Println(*result)
	case "list":
		list()
	default:
		fmt.Println("Not a valid option")
	}
}
