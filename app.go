package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

type Document struct {
	Id          int    `json:"Id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var documents = map[int]Document{}

func respondWithError(w http.ResponseWriter, code int, message string) {
	response, _ := json.Marshal(map[string]string{"error": message})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func getDocumentById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid document ID")
		return
	}

	document, isPresent := documents[id]

	if !isPresent {
		respondWithError(w, http.StatusNotFound, "No document with id : "+fmt.Sprint(id))
		return
	}
	json.NewEncoder(w).Encode(document)
}

func deleteDocumentById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid document ID")
		return
	}

	delete(documents, id)
}

func addDocument(w http.ResponseWriter, r *http.Request) {
	var document Document

	decoder := json.NewDecoder(r.Body)

	if decoder.Decode(&document) != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	id := len(documents) + 1

	documents[id] = document
	json.NewEncoder(w).Encode(&document)
}

func getDocuments(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(documents)
}

func (a *App) Initalize() {
	documents[0] = Document{Id: 0, Name: "document1", Description: "Test 1"}
	documents[1] = Document{Id: 1, Name: "document2", Description: "Test 2"}
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/document/{id}", getDocumentById).Methods("GET")
	a.Router.HandleFunc("/document", getDocuments).Methods("GET")
	a.Router.HandleFunc("/document", addDocument).Methods("POST")
	a.Router.HandleFunc("/document/{id}", deleteDocumentById).Methods("DELETE")
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}

func main() {
	a := App{}
	a.Initalize()
	a.Run()
}
