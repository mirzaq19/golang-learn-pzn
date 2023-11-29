package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Mirzaq")
	if result != "Hello Mirzaq" {
		// error
		panic("Result is not Hello Mirzaq")
	}
}

func TestHelloWorldBudi(t *testing.T) {
	result := HelloWorld("Budi")
	if result != "Hello Budi" {
		// error
		panic("Result is not Hello Budi")
	}
}
