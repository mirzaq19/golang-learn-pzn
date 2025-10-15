package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	jsonString := `{"FirstName":"M. Auliya","MiddleName":"Mirzaq","LastName":"Romdloni","Age":30,"Married":true}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	json.Unmarshal(jsonBytes, customer)

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
}
