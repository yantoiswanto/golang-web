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

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	})

	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hi")
	})

	mux.HandleFunc("/images", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "images")
	})

	mux.HandleFunc("/images/thumbnails/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "thumbnails")
	})

	server := http.Server{Addr: "localhost:8009", Handler: mux}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {

		fmt.Fprintln(writer, request.Method)
		fmt.Fprintln(writer, request.RequestURI)
	}

	server := http.Server{Addr: "localhost:8009", Handler: handler}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
