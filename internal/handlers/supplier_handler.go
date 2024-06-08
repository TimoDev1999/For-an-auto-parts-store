package handlers

import (
	"encoding/json"
	"my-go-project/internal/services"
	"net/http"
)

func GetSuppliers(w http.ResponseWriter, r *http.Request) {
	suppliers, err := services.GetAllSuppliers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(suppliers)
}
