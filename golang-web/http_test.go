package golang_web

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	t.Run("Check body request not nil", func(t *testing.T) {
		assert.NotEqual(t, nil, bodyString)
	})
	t.Run("Check body request is hello world", func(t *testing.T) {
		assert.Equal(t, "Hello World", bodyString, "They should be equal")
	})

	fmt.Println(bodyString)
}
