package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("M.")
	data.PushBack("Auliya")
	data.PushBack("Mirzaq")
	data.PushBack("Romdloni")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
