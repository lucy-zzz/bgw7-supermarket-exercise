package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	products "github.com/lucy-zzz/bgw7-supermarket-exercise/internal/handler"
)

func main() {
	router := chi.NewRouter()
	router.Get(("/ping"), func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(`{message: "pong"}`)
	})

	router.Get("/products", products.Products)
	router.Post("/checkout", products.Checkout)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
