package main

import (
	"embed"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()
	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	t.Run("Hello", func(t *testing.T) {
		request := httptest.NewRequest("GET", "http://localhost:3000/files/hello.txt", nil)
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, request)
		response := recorder.Result()
		body, _ := io.ReadAll(response.Body)

		assert.Equal(t, "Hello http router", string(body))
	})

	t.Run("Goodbye", func(t *testing.T) {
		request := httptest.NewRequest("GET", "http://localhost:3000/files/goodbye.txt", nil)
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, request)
		response := recorder.Result()
		body, _ := io.ReadAll(response.Body)

		assert.Equal(t, "Good bye http router", string(body))
	})
}
