package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

// Item represents a product in the cart
type Item struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Cart represents a user's shopping cart
type Cart struct {
	Items []Item `json:"items"`
}

var (
	carts = make(map[string]Cart) // In-memory store for carts
	mu    sync.Mutex
)

func addToCart(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userID"]
	var item Item
	json.NewDecoder(r.Body).Decode(&item)

	mu.Lock()
	cart := carts[userID]
	cart.Items = append(cart.Items, item)
	carts[userID] = cart
	mu.Unlock()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cart)
}

func getCart(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userID"]

	mu.Lock()
	cart, exists := carts[userID]
	mu.Unlock()

	if !exists {
		http.Error(w, "Cart not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/cart/{userID}", getCart).Methods("GET")
	r.HandleFunc("/cart/{userID}/add", addToCart).Methods("POST")

	log.Println("Cart service running on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", r))
}
