package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}
	firstname := request.PostForm.Get("firstname")
	lastname := request.PostForm.Get("lastname")
	fmt.Fprintf(writer, "%s %s", firstname, lastname)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("firstname=M.%20Auliya&lastname=Mirzaq%20Romdloni")
	request := httptest.NewRequest(http.MethodPost, "http://localhost/", requestBody)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	record := httptest.NewRecorder()

	FormPost(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}
