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
		fmt.Fprintf(writer, "Hello %s, My Name is %s", name, name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8009/hello?name=Yanto", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request) {
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8009/hello?first_name=Yanto&last_name=Iswanto", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func MultipleParameterValues(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]
	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestMultipleParameterValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8009/hello?name=Muhammad&name=Arshaka&name=Deansyah", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
