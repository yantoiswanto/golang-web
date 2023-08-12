package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "Hello World")
		if err != nil {
			return
		}
	}
	server := http.Server{Addr: "localhost:8009", Handler: handler}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
