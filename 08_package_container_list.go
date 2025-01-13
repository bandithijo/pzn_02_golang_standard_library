package main

import (
	"container/list"
	"fmt"
)

/*
Package container/list
- Package container/list adalah implementasi struktur data double linked list di Go
- <https://pkg.go.dev/container/list>
- Kegunaan struktur dta double linked list biasanya untuk antrian
*/

func main() {
	var data *list.List = list.New()
	data.PushBack("Rizqi")
	data.PushBack("Nur")
	data.PushBack("Assyaufi")

	var head *list.Element = data.Front()
	// head.Next().Next().Next()
	fmt.Println(head.Value) // => Rizqi
	next :=  head.Next()
	fmt.Println(next.Value) // => Nur
	next = next.Next()
	fmt.Println(next.Value) // => Assyaufi

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
