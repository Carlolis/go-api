package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
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

func TestGetNonExistentDocument(t *testing.T) {

	req, _ := http.NewRequest("GET", "/document/11111", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func TestValidGet(t *testing.T) {

	req, _ := http.NewRequest("GET", "/document/0", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

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

func TestGetAll(t *testing.T) {

	req, _ := http.NewRequest("GET", "/document", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

}

func TestPostDocument(t *testing.T) {

	data := Document{Name: "document42", Description: "Test 42"}

	jsonValue, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "/document", bytes.NewBuffer(jsonValue))
	req.Header.Add("Content-Type", "application/json")
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

}

func TestPostInvalidPayload(t *testing.T) {

	data := "\"Name\": \"Document42\", \"commentaire\": \"Test 42\""

	req, _ := http.NewRequest("POST", "/document", strings.NewReader(data))
	req.Body.Close()
	req.Header.Add("Content-Type", "text/plain")
	response := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)
}

func TestPostDocumentWithoutName(t *testing.T) {

	data := Document{Description: "Test 42"}

	jsonValue, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "/document", bytes.NewBuffer(jsonValue))
	req.Header.Add("Content-Type", "application/json")
	response := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)
}

func TestPostDocumentWithId(t *testing.T) {

	data := Document{Id: 1, Name: "document14", Description: "Test 14"}

	jsonValue, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "/document", bytes.NewBuffer(jsonValue))
	req.Header.Add("Content-Type", "application/json")
	response := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)
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
