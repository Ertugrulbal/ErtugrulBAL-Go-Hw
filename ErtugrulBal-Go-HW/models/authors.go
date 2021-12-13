package models

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Author struct
type Author struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var Authors []Author

//**** Methods
func GetAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Authors)
}

// Get single book
func GetAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range Authors {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Author{})
}

// Add new book
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var author Author
	_ = json.NewDecoder(r.Body).Decode(&author)
	author.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	Authors = append(Authors, author)
	json.NewEncoder(w).Encode(author)
}

// Update Author
func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Authors {
		if item.ID == params["id"] {
			Authors = append(Authors[:index], Authors[index+1:]...)
			var author Author
			_ = json.NewDecoder(r.Body).Decode(&author)
			author.ID = params["id"]
			Authors = append(Authors, author)
			json.NewEncoder(w).Encode(author)
			return
		}
	}
}

// Delete Author
func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Authors {
		if item.ID == params["id"] {
			Authors = append(Authors[:index], Authors[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Authors)
}
