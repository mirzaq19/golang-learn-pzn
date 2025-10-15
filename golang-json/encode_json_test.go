package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data any) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	logJson("Mirzaq")
	logJson(1)
	logJson(true)
	logJson([]string{"M.", "Auliya", "Mirzaq", "Romdloni"})
}
