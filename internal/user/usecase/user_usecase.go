package usecase

import "github.com/Capstone-Kel-23/BE-Rest-API/domain"

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(ur domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepository: ur,
	}
}

func (u *userUsecase) GetDetailUserByEmail(email string) (*domain.User, error) {
	user, err := u.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
