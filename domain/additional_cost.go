package domain

import uuid "github.com/satori/go.uuid"

type AdditionalCost struct {
	ID        uuid.UUID `json:"id" gorm:"PrimaryKey"`
	InvoiceID uuid.UUID `json:"invoice_id" gorm:"type:varchar;size:191"`
	Type      string    `json:"type"`
	Name      string    `json:"name"`
	Total     float64   `json:"total"`
}
