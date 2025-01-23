package main

import (
	"container/list"
	"fmt"
)

func outputList(array list.List) {
	for element := array.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}

func main() {
	var initList list.List

	initList.PushBack(12)
	initList.PushBack(33)
	initList.PushBack(44)

	fmt.Println("Len of list -> ", initList.Len())
	fmt.Println("Remove element of list -> ", initList.Remove(initList.Front()))
	outputList(initList)
}
