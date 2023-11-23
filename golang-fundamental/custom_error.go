package main

import "fmt"

type validationError struct {
	Message string
}
type notFoundError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

func (v *notFoundError) Error() string {
	return v.Message
}

func saveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation error"}
	}

	if id != "Mirzaq" {
		return &notFoundError{"Not found error"}
	}

	return nil
}

func halo(name string) any {
	return "Halo " + name
}

func main() {
	err := saveData("budi", nil)
	if err != nil {
		// terjadi error
		//if validationErr, ok := err.(*validationError); ok {
		//	fmt.Println("validation error", validationErr.Error())
		//} else if notFoundErr, ok := err.(*notFoundError); ok {
		//	fmt.Println("Not Found error", notFoundErr.Error())
		//} else {
		//	fmt.Println("Unknown Error")
		//}

		switch err.(type) {
		case *validationError:
			fmt.Println("Validation error:", err.Error())
		case *notFoundError:
			fmt.Println("Not found error", err.Error())
		default:
			fmt.Println("Unknown error")
		}
	} else {
		fmt.Println("Sukses")
	}
}
