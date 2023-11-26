package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, "with value type", structField.Type)
		fmt.Println("tag required for", structField.Name, ":", structField.Tag.Get("required"))
		fmt.Println("tag max for", structField.Name, ":", structField.Tag.Get("max"))
	}
}

func isValid(value any) (result bool) {
	t := reflect.TypeOf(value)
	result = true
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			//fmt.Println(data)
			//result = result && data != ""
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {
	readField(Sample{"Mirzaq"})
	readField(Person{"Mirzaq", "mojokerto", "mirzaq@gmail.com"})

	mirzaq := Person{"Eko", "Mojokerto", "mirzaq@gmail.com"}
	fmt.Println(isValid(mirzaq))
}
