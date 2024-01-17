package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", map[string]any{
		"Title": "Template Data Map",
		"Name":  "M. Auliya Mirzaq Romdloni",
		"Address": map[string]any{
			"Street": "Jl. xxx no.123",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	record := httptest.NewRecorder()

	TemplateDataMap(record, request)

	body, _ := io.ReadAll(record.Result().Body)
	fmt.Println(string(body))
}

type Address struct {
	Street string
}

type Page struct {
	Title   string
	Name    string
	Address Address
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "M. Auliya Mirzaq Romdloni",
		Address: Address{
			Street: "Jl. xxx no.123",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	record := httptest.NewRecorder()

	TemplateDataStruct(record, request)

	body, _ := io.ReadAll(record.Result().Body)
	fmt.Println(string(body))
}
