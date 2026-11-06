package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"payreminder/rest/middleware"
	"strconv"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func Start() {

	// mngr := middleware.NewManager()

	mux := http.NewServeMux()

	mux.Handle(
		"GET /products",
		middleware.CORSMiddleware(middleware.Logger(http.HandlerFunc(Product))),
	)

	mux.Handle(
		"GET /products/{id}",
		middleware.CORSMiddleware(middleware.Logger(http.HandlerFunc(ProductId))),
	)

	mux.Handle(
		"POST /createproduct",
		middleware.CORSMiddleware(middleware.Logger(http.HandlerFunc(Createproduct))),
	)
	// Swagger UI
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
	fmt.Println("Server running on : 8080")
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println("Error starting server", err)
		return
	}
}

func Createproduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, "Give me POST request", http.StatusMethodNotAllowed)
		return
	}
	var newproduct Productinfo

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newproduct)

	if err != nil {
		http.Error(w, "Please give me valid json", 400)
		return
	}

	newproduct.ID = len(productlist) + 1
	productlist = append(productlist, newproduct)

	encoder := json.NewEncoder(w)
	encoder.Encode(newproduct)

}

func Product(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		http.Error(w, "Give me GET request", http.StatusMethodNotAllowed)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(productlist)
}

func ProductId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	pId := r.PathValue("id")

	id, err := strconv.Atoi(pId)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for _, product := range productlist {
		if product.ID == id {
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

var productlist []Productinfo

type Productinfo struct {
	ID          int
	Name        string
	Price       float64
	Description string
}

func init() {

	pro1 := Productinfo{
		ID:          1,
		Name:        "Laptop",
		Price:       75000,
		Description: "Gaming Laptop",
	}

	pro2 := Productinfo{
		ID:          2,
		Name:        "Smartphone",
		Price:       35000,
		Description: "Android Smartphone",
	}

	pro3 := Productinfo{
		ID:          3,
		Name:        "Keyboard",
		Price:       2500,
		Description: "Mechanical Keyboard",
	}

	pro4 := Productinfo{
		ID:          4,
		Name:        "Mouse",
		Price:       1500,
		Description: "Wireless Mouse",
	}

	pro5 := Productinfo{
		ID:          5,
		Name:        "Monitor",
		Price:       22000,
		Description: "24-inch Full HD Monitor",
	}

	productlist = append(productlist, pro1, pro2, pro3, pro4, pro5)
}
