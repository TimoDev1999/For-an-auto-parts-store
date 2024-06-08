package models

type Supplier struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ContactName string `json:"contact_name"`
	PhoneNumber string `json:"phone_number"`
	ContactInfo string `json:"contact_info"`
}
