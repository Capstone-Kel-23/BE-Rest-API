package usecase

import (
	"errors"
	"github.com/Capstone-Kel-23/BE-Rest-API/domain"
	"github.com/Capstone-Kel-23/BE-Rest-API/internal/validation/utils"
	uuid "github.com/satori/go.uuid"
	"time"
)

type validationUsecase struct {
	validationRepository domain.ValidationRepository
	userRepository       domain.UserRepository
}

func NewValidationUsecase(vr domain.ValidationRepository, ur domain.UserRepository) domain.ValidationUsecase {
	return &validationUsecase{
		validationRepository: vr,
		userRepository:       ur,
	}
}

func (v *validationUsecase) VerifyUser(code, email string) (string, error) {
	existingValidation, _ := v.validationRepository.FindByCode(code)
	if existingValidation.ID == uuid.FromStringOrNil("") {
		return "", errors.New("failed, send again verify")
	}
	if existingValidation.ExpiredAt.Before(time.Now()) {
		v.validationRepository.Delete(existingValidation)
		return "", errors.New("failed, email verification is expired")
	}
	if existingValidation.Code != code {
		return "", errors.New("failed, send again verify")
	}
	v.validationRepository.Delete(existingValidation)
	return existingValidation.UserEmail, nil
}

func (v *validationUsecase) CreateVerifyUser(email, t string) (string, error) {
	existingValidation, _ := v.validationRepository.FindByEmailAndType(email, t)
	if existingValidation.ID != uuid.FromStringOrNil("") {
		v.validationRepository.Delete(existingValidation)
	}
	newValidation := &domain.Validation{
		ID:        uuid.NewV4(),
		UserEmail: email,
		Type:      t,
		Code:      utils.GenerateRandomString(8),
		ExpiredAt: time.Now().Add(time.Hour * time.Duration(48)),
	}
	validation, err := v.validationRepository.Save(newValidation)
	if err != nil {
		return "", err
	}
	return validation.Code, nil
}
