package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	apiRoot := "/api"
	//  /api/....
	http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
		message := API{"API Home Root"}
		output, err := json.Marshal(message)
		checkError(err)
		fmt.Fprintf(w, string(output))
	})

	// .../api/about
	http.HandleFunc(apiRoot+"/about", func(w http.ResponseWriter, r *http.Request) {
		output, err := json.Marshal("This Application was created at Go-Bootcamp")
		checkError(err)
		fmt.Fprintf(w, string(output))
	})

	// .../api/users
	http.HandleFunc(apiRoot+"/users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			{ID: 1, FirstName: "Zeynep", LastName: "Kesici", Age: 32},
			{ID: 2, FirstName: "Murtaza", LastName: "Biçici", Age: 56},
			{ID: 5, FirstName: "Adem", LastName: "Hazreti", Age: 74},
		}
		output, err := json.Marshal(users)
		checkError(err)
		fmt.Fprintf(w, string(output))
	})

	// .../api/Products
	http.HandleFunc(apiRoot+"/products", func(w http.ResponseWriter, r *http.Request) {
		products := []Products{
			{ID: 1, Name: "Kaban", Color: "Siyah", Size: "XL"},
			{ID: 2, Name: "Sweatshirt", Color: "Beyaz", Size: "L"},
			{ID: 3, Name: "Şal", Color: "Mor", Size: "M"},
		}
		output, err := json.Marshal(products)
		checkError(err)
		fmt.Fprintf(w, string(output))
	})

	// .../api/Categories
	http.HandleFunc(apiRoot+"/categories", func(w http.ResponseWriter, r *http.Request) {
		categories := []Categories{
			{Name: "Underwear"},
			{Name: "Top"},
			{Name: "Accessory"},
		}
		output, err := json.Marshal(categories)
		checkError(err)
		fmt.Fprintf(w, string(output))
	})

	//    .../api/Customer
	http.HandleFunc(apiRoot+"/customers", func(w http.ResponseWriter, r *http.Request) {
		customers := []Customers{
			{ID: 1, Name: "Ertugrul BAL", Phone: 5558896545, Address: "Ankara"},
			{ID: 2, Name: "Sema Oguz", Phone: 533212457, Address: "İstanbul"},
			{ID: 3, Name: "Derya Bingöl", Phone: 5526898798, Address: "İzmir"},
		}
		// message := users
		output, err := json.Marshal(customers)
		checkError(err)
		fmt.Fprintf(w, string(output))
	})
	//     .../api/Orders
	http.HandleFunc(apiRoot+"/orders", func(w http.ResponseWriter, r *http.Request) {
		orders := []Orders{
			{ID: 1, ProductName: "Kaban", Quantity: 3, Price: 100},
			{ID: 2, ProductName: "Kazak", Quantity: 5, Price: 50},
			{ID: 3, ProductName: "Sweat", Quantity: 1, Price: 60},
		}
		output, err := json.Marshal(orders)
		checkError(err)
		fmt.Fprintf(w, string(output))
	})

	http.ListenAndServe(":8080", nil)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error : ", err.Error())
		os.Exit(1)
	}
}

// Struct of Users

type User struct {
	ID        int    `json:id`
	FirstName string `json:firstname`
	LastName  string `json:lastname`
	Age       int    `json:age`
}

// Struct of Products

type Products struct {
	ID    int    `json:id`
	Name  string `json:name`
	Color string `json:color`
	Size  string `json:size`
}

// Struct of Customers

type Customers struct {
	ID      int    `json:id`
	Name    string `json:name`
	Phone   uint64 `json:phone`
	Address string `json:address`
}

//	Struct of Orders

type Orders struct {
	ID          int     `json:id`
	ProductName string  `json:productname`
	Quantity    int     `json:quantity`
	Price       float64 `json:price`
}

//	Struct of Categories

type Categories struct {
	Name string `json:name`
}

//	Struct of API
type API struct {
	Message string `json:message`
}
