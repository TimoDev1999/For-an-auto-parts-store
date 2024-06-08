package services

import (
	"my-go-project/internal/models"
	"my-go-project/pkg/database"
)

func GetAllParts() ([]models.Part, error) {
	db := database.GetDB()
	rows, err := db.Query("SELECT id, name, description, price, quantity FROM parts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var parts []models.Part
	for rows.Next() {
		var part models.Part
		if err := rows.Scan(&part.ID, &part.Name, &part.Description, &part.Price, &part.Quantity); err != nil {
			return nil, err
		}
		parts = append(parts, part)
	}
	return parts, nil
}
