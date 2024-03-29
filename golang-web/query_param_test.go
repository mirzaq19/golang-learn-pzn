package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello, %s!", name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=Mirzaq", nil)
	record := httptest.NewRecorder()

	SayHello(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request) {
	firstname := request.URL.Query().Get("firstname")
	lastname := request.URL.Query().Get("lastname")
	fmt.Fprintf(writer, "Hello, %s %s", firstname, lastname)
}
func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?firstname=M.%20Auliya&lastname=Mirzaq%20Romdloni", nil)
	record := httptest.NewRecorder()

	MultipleQueryParameter(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleParameterValues(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]
	fmt.Fprintf(writer, strings.Join(names, ", "))
}

func TestMultipleParameterValues(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=M.%20Auliya&name=Mirzaq%20Romdloni", nil)
	record := httptest.NewRecorder()

	MultipleParameterValues(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
