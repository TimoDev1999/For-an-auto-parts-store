package models

type Customer struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ContactInfo string `json:"contact_info"`
}
