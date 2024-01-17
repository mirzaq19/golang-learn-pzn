package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml", "./templates/layout.gohtml", "./templates/footer.gohtml",
	))
	t.ExecuteTemplate(writer, "layout", map[string]any{
		"Title":       "Template Layout",
		"TitleHeader": "Ini Header Layout",
		"Name":        "Mirzaq",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	record := httptest.NewRecorder()

	TemplateLayout(record, request)

	body, _ := io.ReadAll(record.Result().Body)
	fmt.Println(string(body))
}
