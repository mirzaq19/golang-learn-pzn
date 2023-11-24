package main

import (
	"errors"
	"fmt"
)

var (
	validationError = errors.New("validation error")
	notFoundError   = errors.New("not fount error")
)

func getById(id string) error {
	if id == "" {
		return validationError
	}

	if id != "mirzaq" {
		return notFoundError
	}

	return nil
}

func main() {
	err := getById("budi")
	if err != nil {
		if errors.Is(err, validationError) {
			fmt.Println("Validation error")
		} else if errors.Is(err, notFoundError) {
			fmt.Println("Not found error")
		} else {
			fmt.Println("Unknown error")
		}
	}
}
