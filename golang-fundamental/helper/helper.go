package helper

import "fmt"

var version = "1.0.0" //Tidak bisa diakses dari luar
var Application = "golang"

func sayGoodBye(name string) string {
	return "Good bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}

func Contoh() {
	sayGoodBye("Mirzaq")
	fmt.Println(version)
}
