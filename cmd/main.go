package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	products "github.com/lucy-zzz/bgw7-supermarket-exercise/internal/handler"
)

func main() {
	router := chi.NewRouter()
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
	})

	router.Get("/products", products.Products)
	router.Get("/products/{id}", products.GetProductById)
	router.Get("/products/search", products.SearchProducts)
	// router.Post("/checkout", products.Checkout)

	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
