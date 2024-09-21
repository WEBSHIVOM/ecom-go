package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	_ "cart/docs" // Import generated docs

	"github.com/gorilla/mux"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Cart Microservice API
// @version 1.0
// @description This is the API documentation for the Cart microservice.
// @host localhost:8081
// @BasePath /

type Item struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Cart struct {
	Items []Item `json:"items"`
}

var (
	carts = make(map[string]Cart)
	mu    sync.Mutex
)

// @Summary Add an item to the cart
// @Description Adds an item to the user's cart by user ID
// @Tags cart
// @Accept json
// @Produce json
// @Param userID path string true "User ID"
// @Param item body Item true "Item"
// @Success 200 {object} Cart
// @Router /cart/{userID}/add [post]
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

// @Summary Get a user's cart
// @Description Retrieves a user's cart by user ID
// @Tags cart
// @Produce json
// @Param userID path string true "User ID"
// @Success 200 {object} Cart
// @Router /cart/{userID} [get]
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

	// Add Swagger docs route
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	r.HandleFunc("/cart/{userID}", getCart).Methods("GET")
	r.HandleFunc("/cart/{userID}/add", addToCart).Methods("POST")

	log.Println("Cart service running on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", r))
}
