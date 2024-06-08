package handlers

import (
	"encoding/json"
	"my-go-project/internal/services"
	"net/http"
)

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := services.GetAllCustomers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(customers)
}
