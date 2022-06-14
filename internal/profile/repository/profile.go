package repository

import (
	"github.com/Capstone-Kel-23/BE-Rest-API/domain"
	"gorm.io/gorm"
)

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepoitory(db *gorm.DB) domain.ProfileRepository {
	return &profileRepository{
		db: db,
	}
}

func (p *profileRepository) Save(profile *domain.Profile) (*domain.Profile, error) {
	err := p.db.Save(&profile).Error
	return profile, err
}

func (p *profileRepository) UpdateByUserID(userid string, profile *domain.Profile) (*domain.Profile, error) {
	err := p.db.Model(&profile).Where("user_id = ? ", userid).Updates(&profile).Error
	return profile, err
}

func (p *profileRepository) FindByUserID(userid string) (profile *domain.Profile, err error) {
	err = p.db.Where("user_id = ?", userid).Find(&profile).Error
	return profile, err
}
