package models

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Struct
type Store struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

var Stores []Store

//  Stores' methods

// Get all Stores
func GetStores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Stores)
}

// Get single Stores
func GetStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range Stores {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Store{})
}

// Add new Stores
func CreateStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var store Store
	_ = json.NewDecoder(r.Body).Decode(&store)
	store.ID = strconv.Itoa(rand.Intn(10000)) // Mock ID - not safe
	Stores = append(Stores, store)
	json.NewEncoder(w).Encode(store)
}

// Update Stores
func UpdateStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Stores {
		if item.ID == params["id"] {
			Stores = append(Stores[:index], Stores[index+1:]...)
			var store Store
			_ = json.NewDecoder(r.Body).Decode(&store)
			store.ID = params["id"]
			Stores = append(Stores, store)
			json.NewEncoder(w).Encode(store)
			return
		}
	}
}

// Delete Stores
func DeleteStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Stores {
		if item.ID == params["id"] {
			Stores = append(Stores[:index], Stores[index+1:]...)
			break

		}
	}
	json.NewEncoder(w).Encode(Stores)
}
