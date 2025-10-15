package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "P0001",
		Name:     "Apple Mac Book Pro",
		Price:    12000000,
		ImageURL: "https://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Apple Mac Book Pro","price":12000000,"image_url":"https://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}
	json.Unmarshal(jsonBytes, product)

	fmt.Println(product)
	fmt.Println(product.Id)
	fmt.Println(product.Name)
	fmt.Println(product.Price)
	fmt.Println(product.ImageURL)
}
