package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


type App struct {
	Router *mux.Router
}


type Document struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Description string `json:"description"`
}


// bouchon
var documents []Document 

func respondWithError(w http.ResponseWriter, code int, message string) {
	response, _ := json.Marshal(map[string]string{"error": message})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func getDocumentById(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	i, err := strconv.Atoi(params["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid document ID")
		return
  }



	for _, document := range documents {
		if document.ID == i {
			json.NewEncoder(w).Encode(&document)
		}
	}
}


func deleteDocumentById(w http.ResponseWriter, r *http.Request)  {
  params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid document ID")
		return
}

	for i, document := range documents {
		if document.ID == id {
			documents = append(documents[:i],documents[i+1:]... )
		}
	}
}

func  addDocument(w http.ResponseWriter, r *http.Request)  {
	var document Document

	decoder := json.NewDecoder(r.Body)

	

	if  decoder.Decode(&document) != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}


	if document.ID != 0 {
		respondWithError(w, http.StatusBadRequest, "A new document must not have an id")
		return
	}

	document.ID = len(documents)+1


	documents = append(documents, document)
  json.NewEncoder(w).Encode(&document)
}

func getDocuments(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(documents)
}


func (a *App) Initalize() {
	documents = append(documents,Document{ID:1,Name:"document1",Description:"Test 1"},Document{ID:2,Name:"document2",Description:"Test 2"})
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/document/{id}",getDocumentById).Methods("GET")
	a.Router.HandleFunc("/document",getDocuments).Methods("GET")
	a.Router.HandleFunc("/document",addDocument).Methods("POST")
	a.Router.HandleFunc("/document/{id}",deleteDocumentById).Methods("DELETE")
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}


func main() {
	a := App{}
	a.Initalize()
	a.Run()
}
