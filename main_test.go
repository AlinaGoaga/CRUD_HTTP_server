package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)


func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}

func executeRequest(req *http.Request, function func(http.ResponseWriter, *http.Request)) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(function)
	handler.ServeHTTP(rr, req)
	
    return rr
}

func TestReturnAllBooks(t *testing.T) {
	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := executeRequest(req,returnAllBooks)
	checkResponseCode(t, http.StatusOK, res.Code)

	var expected = `[{"id":"1","title":"Book1","author":"Author1"},{"id":"2","title":"Book2","author":"Author2"}]`
	
    fmt.Println(reflect.TypeOf(expected))

	if res.Body.String() != expected {
		var body = res.Body
		fmt.Printf(body)
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}

func TestReturnSingleBook(t *testing.T) {
	req, err := http.NewRequest("GET", "/book/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := executeRequest(req, returnSingleBook)
	checkResponseCode(t, http.StatusOK, res.Code)

	expected := `{"id":"1","title":"Book1","author":"Author1"}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}

func TestCreateNewBook(t *testing.T) {
    payload := []byte(`{"id":"3","title":"Book3","author":"Author3"}`)

	req, err := http.NewRequest("POST", "/book", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	res := executeRequest(req, createNewBook)
	checkResponseCode(t, http.StatusOK, res.Code)

	var book map[string]interface{}
    json.Unmarshal(res.Body.Bytes(), &book)

    if book["id"] != "3" {
        t.Errorf("Expected book id to be '3'. Got '%v'", book["id"])
    }
}