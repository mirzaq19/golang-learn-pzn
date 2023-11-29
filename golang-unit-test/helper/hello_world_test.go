package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Mirzaq")
	assert.Equal(t, "Hello Mirzaq", result, "Result must be 'Hello Mirzaq'")

	fmt.Println("TestHelloWorld with assert done!")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Mirzaq")
	require.Equal(t, "Hello Mirzaq", result, "Result must be 'Hello Mirzaq'")

	fmt.Println("TestHelloWorld with require done!")
}

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
