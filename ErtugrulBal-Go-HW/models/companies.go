package models

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Struct
type Company struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

var Companies []Company

//  Companies' methods

// Get all Companies
func GetCompanies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Companies)
}

// Get single Company
func GetCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range Companies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Company{})
}

// Add new Company
func CreateCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var company Company
	_ = json.NewDecoder(r.Body).Decode(&company)
	company.ID = strconv.Itoa(rand.Intn(10000)) // Mock ID - not safe
	Companies = append(Companies, company)
	json.NewEncoder(w).Encode(company)
}

// Update Company
func UpdateCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index1, item := range Companies {
		if item.ID == params["id"] {
			Companies = append(Companies[:index1], Companies[index1+1:]...)
			var company Company
			_ = json.NewDecoder(r.Body).Decode(&company)
			company.ID = params["id"]
			Companies = append(Companies, company)
			json.NewEncoder(w).Encode(company)
			return
		}
	}
}

// Delete Company
func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Companies {
		if item.ID == params["id"] {
			Companies = append(Companies[:index], Companies[index+1:]...)
			break

		}
	}
	json.NewEncoder(w).Encode(Companies)
}
