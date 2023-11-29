package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldMirzaq(t *testing.T) {
	result := HelloWorld("Mirzaq")
	if result != "Hello Mirzaq" {
		// error
		t.Error("Result must be 'Hello Mirzaq'")
	}

	fmt.Println("TestHelloWorldMirzaq Done!")
}

func TestHelloWorldBudi(t *testing.T) {
	result := HelloWorld("Budi")
	if result != "Hello Budi" {
		// error
		t.Fatal("Result must be 'Hello Budi'")
	}

	fmt.Println("TestHelloWorldBudi Done!")
}
