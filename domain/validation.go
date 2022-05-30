package domain

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Validation struct {
	ID        uuid.UUID `json:"ID" gorm:"PrimaryKey"`
	UserEmail string    `json:"user_email" gorm:"notnull"`
	Type      string    `json:"type"`
	Code      string    `json:"code"`
	ExpiredAt time.Time `json:"expired_at"`
}

type ValidationRepository interface {
	FindByEmailAndType(email string, t string) (*Validation, error)
	FindByCode(code string) (*Validation, error)
	Save(validation *Validation) (*Validation, error)
	Delete(validation *Validation) error
}

type ValidationUsecase interface {
	VerifyUser(code string, email string) (string, error)
	CreateVerifyUser(email, t string) (string, error)
}
