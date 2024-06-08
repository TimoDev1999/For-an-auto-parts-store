package services

import (
	"my-go-project/internal/models"
	"my-go-project/pkg/database"
)

func GetAllCustomers() ([]models.Customer, error) {
	db := database.GetDB()
	rows, err := db.Query("SELECT id, name, contact_info FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var customer models.Customer
		if err := rows.Scan(&customer.ID, &customer.Name, &customer.ContactInfo); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}
