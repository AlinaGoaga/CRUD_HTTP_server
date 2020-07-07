package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func executeRouterRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	router := routes()
	router.ServeHTTP(rr, req)

	return rr
}

func TestReturnAllBooks(t *testing.T) {
	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := executeRouterRequest(req)
	checkResponseCode(t, http.StatusOK, res.Code)

	expected := `[{"id":"1","title":"Book1","author":"Author1"},{"id":"2","title":"Book2","author":"Author2"}]`

	res_body := strings.TrimSpace(res.Body.String())
	if res_body != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res_body, expected)
	}
}

func TestReturnSingleBook(t *testing.T) {
	req, err := http.NewRequest("GET", "/book/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := executeRouterRequest(req)
	checkResponseCode(t, http.StatusOK, res.Code)

	expected := `{"id":"1","title":"Book1","author":"Author1"}`

	res_body := strings.TrimSpace(res.Body.String())
	if res_body != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res_body, expected)
	}
}

func TestDeleteBook(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/book/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := executeRouterRequest(req)
	checkResponseCode(t, http.StatusOK, res.Code)

	// check if I can still get the book
	req, _ = http.NewRequest("GET", "/book/1", nil)
	get_single_res := executeRouterRequest(req)
	// res2 code is 200 instead of 404 - why?
	checkResponseCode(t, http.StatusNotFound, get_single_res.Code)

	// check if the book is still in the list
	expected := `[{"id":"2","title":"Book2","author":"Author2"}]`
	req, _ = http.NewRequest("GET", "/books", nil)
	get_all_res := executeRouterRequest(req)
	res_body := strings.TrimSpace(get_all_res.Body.String())
	if res_body != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res_body, expected)
	}
}

func TestCreateNewBook(t *testing.T) {
	payload := []byte(`{"id":"3","title":"Book3","author":"Author3"}`)

	req, err := http.NewRequest("POST", "/book", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	res := executeRouterRequest(req)
	checkResponseCode(t, http.StatusOK, res.Code)

	var book map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &book)

	if book["id"] != "3" {
		t.Errorf("Expected book id to be '3'. Got '%v'", book["id"])
	}
}
