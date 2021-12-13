package main

import (
	"log"
	"net/http"

	. "github.com/ertugrulbal/models"
	"github.com/gorilla/mux"
)

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database Books
	Books = append(Books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	Books = append(Books, Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})
	// Route handles & endpoints For Books
	r.HandleFunc("/books", GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", GetBook).Methods("GET")
	r.HandleFunc("/books", CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")

	// Hardcoded data - @todo: add database Authors
	Authors = append(Authors, Author{ID: "1", Firstname: "John", Lastname: "Doe"})
	Authors = append(Authors, Author{ID: "2", Firstname: "Steve", Lastname: "Smith"})
	// Route handles & endpoints For Authors
	r.HandleFunc("/authors", GetAuthors).Methods("GET")
	r.HandleFunc("/authors/{id}", GetAuthor).Methods("GET")
	r.HandleFunc("/authors", CreateAuthor).Methods("POST")
	r.HandleFunc("/authors/{id}", UpdateAuthor).Methods("PUT")
	r.HandleFunc("/autors/{id}", DeleteAuthor).Methods("DELETE")

	// Hardcoded data - @todo: add database Companies
	Companies = append(Companies, Company{ID: "1", Name: "D&R", Address: "ANKARA, Çankaya"})
	Companies = append(Companies, Company{ID: "2", Name: "DOST", Address: "İSTANBUL, Kadıköy"})
	// Route handles & endpoints For Companies
	r.HandleFunc("/companies", GetCompanies).Methods("GET")
	r.HandleFunc("/companies/{id}", GetCompany).Methods("GET")
	r.HandleFunc("/companies", CreateCompany).Methods("POST")
	r.HandleFunc("/companies/{id}", UpdateCompany).Methods("PUT")
	r.HandleFunc("/companies/{id}", DeleteCompany).Methods("DELETE")

	// Hardcoded data - @todo: add database Category
	Categories = append(Categories, Category{ID: "1", Name: "Autobiography"})
	Categories = append(Categories, Category{ID: "2", Name: "Business"})
	// Route handles & endpoints For Categories
	r.HandleFunc("/categories", GetCategories).Methods("GET")
	r.HandleFunc("/categories/{id}", GetCategory).Methods("GET")
	r.HandleFunc("/categories", CreateCategory).Methods("POST")
	r.HandleFunc("/categories/{id}", UpdateCategory).Methods("PUT")
	r.HandleFunc("/categories/{id}", DeleteCategory).Methods("DELETE")

	// Hardcoded data - @todo: add database Basket
	Baskets = append(Baskets, Basket{ID: "1", CustomerName: "John Smith", Product: "Issız Adam", TotalPrice: "100"})
	Baskets = append(Baskets, Basket{ID: "2", CustomerName: "Marry Curie", Product: "Yalnızız", TotalPrice: "120"})
	// Route handles & endpoints For Baskets
	r.HandleFunc("/baskets", GetBaskets).Methods("GET")
	r.HandleFunc("/baskets/{id}", GetBasket).Methods("GET")
	r.HandleFunc("/baskets", CreateBasket).Methods("POST")
	r.HandleFunc("/baskets/{id}", UpdateBasket).Methods("PUT")
	r.HandleFunc("/baskets/{id}", DeleteBasket).Methods("DELETE")

	// Hardcoded data - @todo: add database Payment
	Payments = append(Payments, Payment{ID: "1", CustomerName: "Ertugrul BAL", Basket: &Basket{ID: "1", CustomerName: "John Smith", Product: "Issız Adam", TotalPrice: "100"}})
	Payments = append(Payments, Payment{ID: "2", CustomerName: "Necmi Demir", Basket: &Basket{ID: "2", CustomerName: "Marry Curie", Product: "Yalnızız", TotalPrice: "120"}})
	// Route handles & endpoints For Payments
	r.HandleFunc("/payments", GetPayments).Methods("GET")
	r.HandleFunc("/payments/{id}", GetPayment).Methods("GET")
	r.HandleFunc("/payments", CreatePayment).Methods("POST")
	r.HandleFunc("/payments/{id}", UpdatePayment).Methods("PUT")
	r.HandleFunc("/payments/{id}", DeletePayment).Methods("DELETE")

	// Hardcoded data - @todo: add database Stores
	Stores = append(Stores, Store{ID: "1", Name: "Istanbul Branch", Address: "Istanbul, Ulus"})
	Stores = append(Stores, Store{ID: "2", Name: "Stockholm Branch", Address: "Stockholm, Street"})
	// Route handles & endpoints For Stores
	r.HandleFunc("/stores", GetStores).Methods("GET")
	r.HandleFunc("/stores/{id}", GetStore).Methods("GET")
	r.HandleFunc("/stores", CreateStore).Methods("POST")
	r.HandleFunc("/stores/{id}", UpdateStore).Methods("PUT")
	r.HandleFunc("/stores/{id}", DeleteStore).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
