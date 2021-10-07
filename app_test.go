package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a.Initalize()
	code := m.Run()
	os.Exit(code)
}

func TestGetNotValidId(t *testing.T) {

	req, _ := http.NewRequest("GET", "/document/a", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["error"] != "Invalid document ID" {
		t.Errorf("Expected an 'Invalid document ID' error. Got %s", response.Body.String())
	}
}

func TestDeleteNotValidId(t *testing.T) {

	req, _ := http.NewRequest("DELETE", "/document/a", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["error"] != "Invalid document ID" {
		t.Errorf("Expected an 'Invalid document ID' error. Got %s", response.Body.String())
	}
}

func TestGetNonExistentDocument(t *testing.T) {

	req, _ := http.NewRequest("GET", "/document/110", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "" {
		t.Errorf("Expected empty response. Got '%s'", body)
	}
}

func TestGetAll(t *testing.T) {

	req, _ := http.NewRequest("GET", "/document", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

}

func TestPostDocument(t *testing.T) {

	data := map[string]string{"name": "document14", "description": "Test 14"}

	jsonValue, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "/document", bytes.NewBuffer(jsonValue))
	req.Header.Add("Content-Type", "application/json")
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

}

// TODO créer un payload avec l'id comme un entier et non une string
func TestPostDocumentWithId(t *testing.T) {

	// Ne construit pas le bon json, l'ID doit un nombre.
	data := map[string]string{"ID": "10", "name": "document14", "description": "Test 14"}

	jsonValue, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "/document", bytes.NewBuffer(jsonValue))
	req.Header.Add("Content-Type", "application/json")
	response := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["error"] != "A new document must not have an id" {
		t.Errorf("Expected an 'A new document must not have an id' error. Got %s", response.Body.String())
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
