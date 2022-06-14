package http

import (
	"github.com/Capstone-Kel-23/BE-Rest-API/domain"
	"github.com/Capstone-Kel-23/BE-Rest-API/web/response"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProfileController interface {
	UpdateProfileUser(c echo.Context) error
}

type profileController struct {
	profileUsecase domain.ProfileUsecase
}

func NewProfileController(pu domain.ProfileUsecase) ProfileController {
	return &profileController{
		profileUsecase: pu,
	}
}

// UpdateProfileUser godoc
// @Summary Update profile user
// @Description Update profile
// @Tags Profile
// @accept json
// @Produce json
// @Router /profile [put]
// @param data body request.ProfileUpdateRequest true "required"
// @Success 200 {object} response.JSONSuccessResult{data=interface{}}
// @Failure 400 {object} response.JSONBadRequestResult{}
// @Security JWT
func (p *profileController) UpdateProfileUser(c echo.Context) error {
	var req domain.Profile
	if err := c.Bind(&req); err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}
	jwtBearer := c.Get("user").(*jwt.Token)
	claims := jwtBearer.Claims.(jwt.MapClaims)
	userid := claims["UserID"].(string)

	profile, err := p.profileUsecase.UpdateProfileByUserID(userid, &req)
	if err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, true, "success update profile", profile)
}
