package domain

import uuid "github.com/satori/go.uuid"

type Item struct {
	ID          uuid.UUID `json:"id" gorm:"PrimaryKey"`
	InvoiceID   uuid.UUID `json:"invoice_id" gorm:"type:varchar;size:191"`
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	Description string    `json:"description"`
	Tax         float64   `json:"tax"`
	Price       float64   `json:"price"`
	TotalPrice  float64   `json:"total_price"`
}
