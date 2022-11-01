package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type product struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description`
	Price int `json:"price"`
}

var products []product

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func init Products(){
	initProducts(){
		//products = []products
		bs, err := os.ReadFile("products.json")
		if err != nil {
			log.Fatal(err)
		}
		if err := json.Unmarshal(bs, &products); err != nil{
			log.fatal(err)
		}
}

func main() {
	initProducts()

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/products", getProductsHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
