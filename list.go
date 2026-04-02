package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Ben Arief")
	data.PushBack("Windah Basudara")
	data.PushBack("Windut")

	head := data.Front()
	next := head.Next()
	next2 := head.Next().Next()
	fmt.Println(head.Value) // Ben Arief
	fmt.Println(next.Value) // Windah Basudara
	fmt.Println(next2.Value) // Windut

	//perulangan
	for e := data.Front(); e != nil; e =e.Next() {
		fmt.Println(e.Value)
	}
}