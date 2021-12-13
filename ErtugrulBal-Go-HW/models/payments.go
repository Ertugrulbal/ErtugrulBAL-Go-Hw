package models

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// *Payments struct (Model)
type Payment struct {
	ID           string  `json:"id"`
	CustomerName string  `json:"customername"`
	Basket       *Basket `json:"basket"`
}

// Init baskets var as a slice Basket struct
var Payments []Payment

// *Payment Methods ***

// Get all payments
func GetPayments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Payments)
}

// Get single payment
func GetPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through baskets and find one with the id from the params
	for _, item := range Payments {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Payment{})
}

// Add new payment
func CreatePayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var payment Payment
	_ = json.NewDecoder(r.Body).Decode(&payment)
	payment.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	Payments = append(Payments, payment)
	json.NewEncoder(w).Encode(payment)
}

// Update Payment
func UpdatePayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Payments {
		if item.ID == params["id"] {
			Payments = append(Payments[:index], Payments[index+1:]...)
			var payment Payment
			_ = json.NewDecoder(r.Body).Decode(&payment)
			payment.ID = params["id"]
			Payments = append(Payments, payment)
			json.NewEncoder(w).Encode(payment)
			return
		}
	}
}

// Delete Payment
func DeletePayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Payments {
		if item.ID == params["id"] {
			Payments = append(Payments[:index], Payments[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Payments)
}
