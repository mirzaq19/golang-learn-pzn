package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Mojokerto", "Jawa Timur", "Indonesia"}
	address2 := &address1 //pointer

	address2.City = "Surabaya"

	fmt.Println(address1) // ikut berubah
	fmt.Println(address2)

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
}
