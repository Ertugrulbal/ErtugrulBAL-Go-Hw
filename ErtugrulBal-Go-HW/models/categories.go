package models

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// *Category struct (Model)
type Category struct {
	ID   string `json:"id"`
	Name string `json:"isbn"`
}

// Init categories var as a slice Category struct
var Categories []Category

// *Category Methods ***

// Get all categories
func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Categories)
}

// Get single category
func GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through categories and find one with the id from the params
	for _, item := range Categories {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Category{})
}

// Add new category
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var category Category
	_ = json.NewDecoder(r.Body).Decode(&category)
	category.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	Categories = append(Categories, category)
	json.NewEncoder(w).Encode(category)
}

// Update Category
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Categories {
		if item.ID == params["id"] {
			Categories = append(Categories[:index], Categories[index+1:]...)
			var category Category
			_ = json.NewDecoder(r.Body).Decode(&category)
			category.ID = params["id"]
			Categories = append(Categories, category)
			json.NewEncoder(w).Encode(category)
			return
		}
	}
}

// Delete Category
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Categories {
		if item.ID == params["id"] {
			Categories = append(Categories[:index], Categories[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Categories)
}
