package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName, MiddleName, LastName string
	Age                             int
	Married                         bool
	Hobbies                         []string
	Addresses                       []Address
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName:  "M. Auliya",
		MiddleName: "Mirzaq",
		LastName:   "Romdloni",
		Age:        30,
		Married:    true,
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
