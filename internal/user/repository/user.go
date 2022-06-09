package repository

import (
	"github.com/Capstone-Kel-23/BE-Rest-API/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) FindByID(id string) (user *domain.User, err error) {
	err = u.db.Preload("Roles").Where("id = ?", id).Find(&user).Error
	return user, err
}

func (u *userRepository) FindByEmail(email string) (user *domain.User, err error) {
	err = u.db.Preload("Roles").Where("email = ?", email).Find(&user).Error
	return user, err
}

func (u *userRepository) Save(user *domain.User) (*domain.User, error) {
	err := u.db.Create(&user).Error
	return user, err
}

func (u *userRepository) UpdateVerifiedByEmail(email string, status bool) error {
	err := u.db.Model(&domain.User{}).Where("email = ?", email).Update("verified", status).Error
	return err
}

func (u *userRepository) FindAll() (users *domain.Users, err error) {
	err = u.db.Preload("Roles").Find(&users).Error
	return users, err
}

func (u *userRepository) FindWithProfile(id string) (user *domain.User, err error) {
	err = u.db.Preload("Roles").Preload("Profile").Where("id = ?", id).Find(&user).Error
	return user, err
}
