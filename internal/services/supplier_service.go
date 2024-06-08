package services

import (
	"my-go-project/internal/models"
	"my-go-project/pkg/database"
)

func GetAllSuppliers() ([]models.Supplier, error) {
	db := database.GetDB()
	rows, err := db.Query("SELECT id, name, contact_name, phone_number, contact_info FROM suppliers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var suppliers []models.Supplier
	for rows.Next() {
		var supplier models.Supplier
		if err := rows.Scan(&supplier.ID, &supplier.Name, &supplier.ContactName, &supplier.PhoneNumber, &supplier.ContactInfo); err != nil {
			return nil, err
		}
		suppliers = append(suppliers, supplier)
	}
	return suppliers, nil
}
