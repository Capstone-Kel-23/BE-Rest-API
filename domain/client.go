package domain

import uuid "github.com/satori/go.uuid"

type Client struct {
	ID          uuid.UUID `json:"id" gorm:"PrimaryKey"`
	InvoiceID   uuid.UUID `json:"invoice_id" gorm:"type:varchar;size:191"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Address     string    `json:"address"`
	PostalCode  string    `json:"postal_code"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	Country     string    `json:"country"`
	PhoneNumber string    `json:"phone_number"`
}
