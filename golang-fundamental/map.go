package main

import "fmt"

func main() {
	iniMap := map[string]string{
		"name": "Mirzaq",
	}

	fmt.Println(iniMap)

	var person map[string]string = map[string]string{
		"name":    "Mirzaq",
		"address": "Mojokerto",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["age"])

	book := make(map[string]string)
	book["nama"] = "atomic habits"
	book["author"] = "mirzaq"
	book["wrong"] = "ups"
	fmt.Println(book)
	delete(book, "wrong")
	fmt.Println(book)
}
