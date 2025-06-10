package products

import (
	"net/http"
	"os"
)

func Products(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("docs/db/products.json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Failed to read products file"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
