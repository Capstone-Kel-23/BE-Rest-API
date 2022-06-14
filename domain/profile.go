package domain

import uuid "github.com/satori/go.uuid"

type Profile struct {
	ID          uuid.UUID `json:"id" gorm:"PrimaryKey"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:varchar;size:191"`
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

type ProfileRepository interface {
	FindByUserID(userid string) (*Profile, error)
	Save(profile *Profile) (*Profile, error)
	UpdateByUserID(userid string, profile *Profile) (*Profile, error)
}

type ProfileUsecase interface {
	UpdateProfileByUserID(userid string, profile *Profile) (*Profile, error)
}
