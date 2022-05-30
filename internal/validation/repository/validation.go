package repository

import (
	"github.com/Capstone-Kel-23/BE-Rest-API/domain"
	"gorm.io/gorm"
)

type validationRepository struct {
	db *gorm.DB
}

func NewValidationRepository(db *gorm.DB) domain.ValidationRepository {
	return &validationRepository{
		db: db,
	}
}

func (v *validationRepository) FindByEmailAndType(email string, t string) (validation *domain.Validation, err error) {
	err = v.db.Where("user_email = ? AND type = ?", email, t).Find(&validation).Error
	return validation, err
}

func (v *validationRepository) Save(validation *domain.Validation) (*domain.Validation, error) {
	err := v.db.Create(&validation).Error
	return validation, err
}

func (v *validationRepository) Delete(validation *domain.Validation) error {
	err := v.db.Where("user_email = ? AND type = ?", validation.UserEmail, validation.Type).Delete(&validation).Error
	return err
}

func (v *validationRepository) FindByCode(code string) (validation *domain.Validation, err error) {
	err = v.db.Where("code = ?", code).Find(&validation).Error
	return validation, err
}
