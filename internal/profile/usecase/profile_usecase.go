package usecase

import (
	"github.com/Capstone-Kel-23/BE-Rest-API/domain"
	uuid "github.com/satori/go.uuid"
)

type profileUsecase struct {
	profileRepository domain.ProfileRepository
}

func NewProfileUsecase(pr domain.ProfileRepository) domain.ProfileUsecase {
	return &profileUsecase{
		profileRepository: pr,
	}
}

func (p *profileUsecase) UpdateProfileByUserID(userid string, profile *domain.Profile) (*domain.Profile, error) {
	profileExisting, _ := p.profileRepository.FindByUserID(userid)
	if profileExisting.ID == uuid.FromStringOrNil("") {
		newProfile := &domain.Profile{
			ID:     uuid.NewV4(),
			UserID: uuid.FromStringOrNil(userid),
			Email:  profile.Email,
		}
		p.profileRepository.Save(newProfile)
	}
	_, err := p.profileRepository.UpdateByUserID(userid, profile)
	if err != nil {
		return nil, err
	}
	profileRes, _ := p.profileRepository.FindByUserID(userid)
	return profileRes, nil
}
