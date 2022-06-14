package domain

import (
	"github.com/Capstone-Kel-23/BE-Rest-API/web/request"
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"PrimaryKey"`
	Fullname  string    `json:"fullname" gorm:"notnull"`
	Email     string    `json:"email" gorm:"notnull"`
	Username  string    `json:"username" gorm:"notnull"`
	Password  string    `json:"password" gorm:"notnull"`
	Verified  bool      `json:"verified"`
	Roles     []Role    `json:"roles" gorm:"many2many:user_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Profile   Profile   `json:"profile,omitempty" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Invoices  []Invoice `json:"invoices,omitempty" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User

type UserRepository interface {
	FindAll() (*Users, error)
	FindByID(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	UpdateVerifiedByEmail(email string, status bool) error
	FindWithProfile(id string) (*User, error)
	Save(user *User) (*User, error)
}

type AuthUsecase interface {
	Login(req *request.LoginRequest) (interface{}, error)
	Register(user *request.UserCreateRequest) (*User, error)
	UpdateVerifiedUser(email string) error
	CheckIfUserIsAdmin(id string) (bool, error)
}

type UserUsecase interface {
	GetDetailUserByEmail(email string) (*User, error)
	GetListAllUsers() (interface{}, error)
	GetDetailUserProfile(id string) (interface{}, error)
}
