package main

import "fmt"

func main() {
	name := "Mirzaq"

	switch name {
	case "Mirzaq":
		fmt.Println("Hello mirzaq")
	case "Budi":
		fmt.Println("Hello Budi")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Hi, boleh kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
