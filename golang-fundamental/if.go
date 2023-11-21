package main

import "fmt"

func main() {
	var name = "mirzaq"

	if name == "mirzaq" {
		fmt.Println("Hello mirzaq")
	} else if name == "budi" {
		fmt.Println("Hello budi")
	} else {
		fmt.Println("Hi, boleh kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
