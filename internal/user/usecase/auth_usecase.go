package usecase

import (
	"errors"
	"github.com/Capstone-Kel-23/BE-Rest-API/domain"
	"github.com/Capstone-Kel-23/BE-Rest-API/internal/role/utils"
	"github.com/Capstone-Kel-23/BE-Rest-API/internal/user/delivery/http/helper"
	"github.com/Capstone-Kel-23/BE-Rest-API/web/request"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	userRepository domain.UserRepository
	roleRepository domain.RoleRepository
}

func NewAuthUsecase(ur domain.UserRepository, rr domain.RoleRepository) domain.AuthUsecase {
	return &authUsecase{
		userRepository: ur,
		roleRepository: rr,
	}
}

func (a *authUsecase) Login(req *request.LoginRequest) (interface{}, error) {
	user, _ := a.userRepository.FindByEmail(req.Email)
	if user.ID == uuid.FromStringOrNil("") {
		return nil, errors.New("user not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("password wrong")
	}

	jwt := helper.NewGoJWT()
	token := jwt.CreateTokenJWT(user)

	responseLogin := map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"fullname": user.Fullname,
		"email":    user.Email,
		"role":     user.Roles[0].Name,
		"verified": user.Verified,
		"token":    token,
	}

	return responseLogin, nil
}

func (a *authUsecase) Register(user *request.UserCreateRequest) (*domain.User, error) {
	var existingUser *domain.User
	existingUser, _ = a.userRepository.FindByEmail(user.Email)
	if existingUser.ID != uuid.FromStringOrNil("") {
		return nil, errors.New("user already exist")
	}

	clientRole, err := a.roleRepository.FindByName(utils.User.String())
	if err != nil {
		return nil, errors.New("role not found - " + utils.User.String())
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	newUser := domain.User{
		ID:       uuid.NewV4(),
		Fullname: user.Fullname,
		Username: user.Username,
		Email:    user.Email,
		Verified: false,
		Password: string(password),
		Roles:    []domain.Role{*clientRole},
	}

	userSave, err := a.userRepository.Save(&newUser)
	if err != nil {
		return nil, err
	}

	return userSave, nil
}

func (a *authUsecase) UpdateVerifiedUser(email string) error {
	existingUser, _ := a.userRepository.FindByEmail(email)
	if existingUser.ID == uuid.FromStringOrNil("") {
		return errors.New("email not found, please register again")
	}
	err := a.userRepository.UpdateVerifiedByEmail(existingUser.Email, true)
	if err != nil {
		return err
	}
	return nil
}

func (a *authUsecase) CheckIfUserIsAdmin(id string) (bool, error) {
	user, _ := a.userRepository.FindByID(id)
	if user.ID == uuid.FromStringOrNil("") {
		return false, errors.New("user not found")
	}
	for _, val := range user.Roles {
		if val.Name == "ROLE_ADMIN" {
			return true, nil
		}
	}

	return false, errors.New("user role not admin")
}
