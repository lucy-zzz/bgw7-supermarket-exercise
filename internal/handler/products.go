package products

import "net/http"

func Products(w http.ResponseWriter, r *http.Request) {

}

func Checkout(w http.ResponseWriter, r *http.Request) {

}

// var p person
// decoder := json.NewDecoder(r.Body)
// decoder.Decode(&p)

// w.WriteHeader(200)
// w.Header().Set("Content-Type", "application/json")

// resp := fmt.Sprintf("Hello, %s %s", p.Name, p.LastName)

// json.NewEncoder(w).Encode(resp)
