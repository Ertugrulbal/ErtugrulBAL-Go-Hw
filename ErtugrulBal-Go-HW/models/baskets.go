package models

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// *Basket struct (Model)
type Basket struct {
	ID           string `json:"id"`
	CustomerName string `json:"customername"`
	Product      string `json:"product"`
	TotalPrice   string `json:"totalprice"`
}

// Init baskets var as a slice Basket struct
var Baskets []Basket

// *Basket Methods ***

// Get all baskets
func GetBaskets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Baskets)
}

// Get single basket
func GetBasket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through baskets and find one with the id from the params
	for _, item := range Baskets {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Basket{})
}

// Add new book
func CreateBasket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var basket Basket
	_ = json.NewDecoder(r.Body).Decode(&basket)
	basket.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	Baskets = append(Baskets, basket)
	json.NewEncoder(w).Encode(basket)
}

// Update book
func UpdateBasket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Baskets {
		if item.ID == params["id"] {
			Baskets = append(Baskets[:index], Baskets[index+1:]...)
			var basket Basket
			_ = json.NewDecoder(r.Body).Decode(&basket)
			basket.ID = params["id"]
			Baskets = append(Baskets, basket)
			json.NewEncoder(w).Encode(basket)
			return
		}
	}
}

// Delete Basket
func DeleteBasket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Baskets {
		if item.ID == params["id"] {
			Baskets = append(Baskets[:index], Baskets[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Baskets)
}
