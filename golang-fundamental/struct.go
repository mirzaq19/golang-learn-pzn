package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var mirzaq Customer
	mirzaq.Name = "M. Auliya Mirzaq Romdloni"
	mirzaq.Address = "Mojokerto"
	mirzaq.Age = 23

	fmt.Println(mirzaq)
	fmt.Println(mirzaq.Name)
	fmt.Println(mirzaq.Address)
	fmt.Println(mirzaq.Age)

	budi := Customer{
		Name:    "Budi",
		Age:     34,
		Address: "Surabaya",
	}

	fmt.Println(budi)
	fmt.Println(budi.Name)
	fmt.Println(budi.Address)
	fmt.Println(budi.Age)

	joko := Customer{"Joko", "Jakarta", 23}
	fmt.Println(joko)
}
