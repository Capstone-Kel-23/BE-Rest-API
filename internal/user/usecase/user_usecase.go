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

func (u *userUsecase) GetListAllUsers() (interface{}, error) {
	var users []interface{}
	allUsers, err := u.userRepository.FindAll()
	if err != nil {
		return users, err
	}
	for _, user := range *allUsers {
		users = append(users, map[string]interface{}{
			"name":       user.Fullname,
			"username":   user.Username,
			"id":         user.ID,
			"role":       user.Roles[0].Name,
			"email":      user.Email,
			"verified":   user.Verified,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
		})
	}
	return users, nil
}
