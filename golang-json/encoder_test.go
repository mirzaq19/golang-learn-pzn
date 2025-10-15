package golangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("sample_output.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "M. Auliya",
		MiddleName: "Mirzaq",
		LastName:   "Romdloni",
		Age:        30,
		Married:    true,
		Hobbies:    []string{"Gaming", "Coding", "Sleeping"},
	}

	_ = encoder.Encode(customer)

	fmt.Println(customer)
}
