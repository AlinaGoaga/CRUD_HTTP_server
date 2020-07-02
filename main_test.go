package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
)


func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(returnAllBooks)
	handler.ServeHTTP(rr, req)

    return rr
}

func TestReturnAllBooks(t *testing.T) {
	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := executeRequest(req)
	checkResponseCode(t, http.StatusOK, res.Code)

	// Check the response body is what we expect.
	expected := `[{"id":"1","title":"Book1","author":"Author1"},{"id":"2","title":"Book2","author":"Author2"}]`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}

func TestReturnSingleBook(t *testing.T) {
	req, err := http.NewRequest("GET", "/books/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := executeRequest(req)
	checkResponseCode(t, http.StatusOK, res.Code)

	expected := `{"id":"1","title":"Book1","author":"Author1"}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}