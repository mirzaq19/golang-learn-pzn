package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Apple Mac Book Pro","price":12000000,"image_url":"https://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	var result map[string]any
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(int32(result["price"].(float64)))
	fmt.Println(result["image_url"])

	bytes, _ := json.Marshal(result)
	fmt.Println(string(bytes))
}

func TestMapEncode(t *testing.T) {
	product := map[string]any{
		"id":    "P0001",
		"name":  "Apple Mac Book Pro",
		"price": 12000000,
	}

	bytes, _ := json.Marshal(product)

	fmt.Println(string(bytes))
}
