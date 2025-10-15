package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "M. Auliya",
		MiddleName: "Mirzaq",
		LastName:   "Romdloni",
		Age:        30,
		Married:    true,
		Hobbies:    []string{"Gaming", "Coding", "Sleeping"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"M. Auliya","MiddleName":"Mirzaq","LastName":"Romdloni","Age":30,"Married":true,"Hobbies":["Gaming","Coding","Sleeping"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	json.Unmarshal(jsonBytes, customer)

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Mirzaq",
		Addresses: []Address{
			{
				Street:     "Jalan Belum Ada",
				Country:    "Indonesia",
				PostalCode: "999999",
			},
			{
				Street:     "Jalan Lagi Dibangun",
				Country:    "Indonesia",
				PostalCode: "888888",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Mirzaq","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"999999"},{"Street":"Jalan Lagi Dibangun","Country":"Indonesia","PostalCode":"888888"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	json.Unmarshal(jsonBytes, customer)

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"999999"},{"Street":"Jalan Lagi Dibangun","Country":"Indonesia","PostalCode":"888888"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	json.Unmarshal(jsonBytes, addresses)

	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jalan Belum Ada",
			Country:    "Indonesia",
			PostalCode: "999999",
		},
		{
			Street:     "Jalan Lagi Dibangun",
			Country:    "Indonesia",
			PostalCode: "888888",
		},
	}

	bytes, _ := json.Marshal(addresses)

	fmt.Println(string(bytes))
}
