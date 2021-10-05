package main

import (
	"encoding/json"
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
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
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

	if err := decoder.Decode(&document); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
 }


	decoder.Decode(&document)
	documents = append(documents, document)
  json.NewEncoder(w).Encode(&document)
}

func getDocuments(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(documents)
}


func (a *App) Run() {
	documents = append(documents,Document{ID:1,Name:"document1",Description:"Test 1"},Document{ID:2,Name:"document2",Description:"Test 2"})
	r := mux.NewRouter()

	r.HandleFunc("/document/{id:[0-9]+}",getDocumentById).Methods("GET")
	r.HandleFunc("/document",getDocuments).Methods("GET")
	r.HandleFunc("/document",addDocument).Methods("POST")
	r.HandleFunc("/document/{id:[0-9]+}",deleteDocumentById).Methods("DELETE")
	

	http.ListenAndServe(":8010",r)
}

func main() {
	a := App{}
	a.Run()
}
