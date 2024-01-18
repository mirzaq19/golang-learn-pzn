package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (page MyPage) SayHello(name string) string {
	return "Hello " + name + ", My name is " + page.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	templateString := `{{.SayHello "Budi"}}`
	t := template.Must(template.New("FUNCTION").Parse(templateString))

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Mirzaq",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	record := httptest.NewRecorder()

	TemplateFunction(record, request)

	body, _ := io.ReadAll(record.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	templateString := `{{len .Name}}`
	t := template.Must(template.New("FUNCTION").Parse(templateString))

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Mirzaq",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	record := httptest.NewRecorder()

	TemplateFunctionGlobal(record, request)

	body, _ := io.ReadAll(record.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionCreateGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]any{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{upper .Name}}`))

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Mirzaq",
	})
}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	record := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(record, request)

	body, _ := io.ReadAll(record.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionPipelines(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]any{
		"SayHello": func(value string) string {
			return "Hello, " + value
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{SayHello .Name|upper}}`))

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Mirzaq",
	})
}

func TestTemplateFunctionPipelines(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	record := httptest.NewRecorder()

	TemplateFunctionPipelines(record, request)

	body, _ := io.ReadAll(record.Result().Body)
	fmt.Println(string(body))
}
