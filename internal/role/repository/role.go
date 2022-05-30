package repository

import (
	"github.com/Capstone-Kel-23/BE-Rest-API/domain"
	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) domain.RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) FindByName(name string) (role *domain.Role, err error) {
	err = r.db.Where("name = ?", name).Find(&role).Error
	return role, err
}
