package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task-8/model"
	"task-8/validator"
)

func main() {

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/api/products", handleProducts)

	fmt.Println("Server started at port 8000")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var product model.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	isValid, errMessage := validator.ValidateResponse(product)
	if !isValid {
		response := Response{
			StatusCode: http.StatusBadRequest,
			Message:    errMessage,
		}
		w.WriteHeader(response.StatusCode)
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			return
		}
		return
	}
	response := Response{
		StatusCode: http.StatusOK,
	}
	w.WriteHeader(response.StatusCode)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

type Response struct {
	StatusCode int
	Message    []string
}
