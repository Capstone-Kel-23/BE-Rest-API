package http

import (
	"github.com/Capstone-Kel-23/BE-Rest-API/domain"
	"github.com/Capstone-Kel-23/BE-Rest-API/web/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController interface {
	GetListAllUsers(c echo.Context) error
}

type userController struct {
	userUsecase domain.UserUsecase
}

func NewUserController(uu domain.UserUsecase) UserController {
	return &userController{
		userUsecase: uu,
	}
}

// GetListAllUsers godoc
// @Summary Get all list users
// @Description Get all list users
// @Tags User
// @accept json
// @Produce json
// @Router /users [get]
// @Success 200 {object} response.JSONSuccessResult{data=interface{}}
// @Failure 400 {object} response.JSONBadRequestResult{}
// @Security JWT
func (u *userController) GetListAllUsers(c echo.Context) error {
	users, err := u.userUsecase.GetListAllUsers()
	if err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, true, "success get list users", users)
}
