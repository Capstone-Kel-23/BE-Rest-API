package http

import (
	"github.com/Capstone-Kel-23/BE-Rest-API/domain"
	"github.com/Capstone-Kel-23/BE-Rest-API/utils/mail/body_email"
	"github.com/Capstone-Kel-23/BE-Rest-API/utils/mail/send_mail"
	"github.com/Capstone-Kel-23/BE-Rest-API/web/request"
	"github.com/Capstone-Kel-23/BE-Rest-API/web/response"
	"github.com/Capstone-Kel-23/BE-Rest-API/web/validation"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"net/url"
)

type AuthController interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
	VerifyUser(c echo.Context) error
	SendVerifyUser(c echo.Context) error
}

type authController struct {
	authUsecase       domain.AuthUsecase
	validationUsecase domain.ValidationUsecase
	userUsecase       domain.UserUsecase
}

func NewAuthController(au domain.AuthUsecase, vu domain.ValidationUsecase, uu domain.UserUsecase) AuthController {
	return &authController{
		authUsecase:       au,
		validationUsecase: vu,
		userUsecase:       uu,
	}
}

// Login godoc
// @Summary Login user
// @Description Login for get JWT token
// @Tags Auth
// @param data body request.LoginRequest true "required"
// @accept json
// @Produce json
// @Router /login [post]
// @Success 200 {object} response.JSONSuccessResult{}
// @Failure 400 {object} response.JSONBadRequestResult{}
// @Failure 401 {object} response.JSONUnauthorizedResult{}
func (a *authController) Login(c echo.Context) error {
	var req request.LoginRequest

	if err := c.Bind(&req); err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}

	res, err := a.authUsecase.Login(&req)
	if err != nil {
		return response.FailResponse(c, http.StatusUnauthorized, false, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, true, "login success", res)
}

// Register godoc
// @Summary Register new user
// @Description Register for create new user
// @Tags Auth
// @param data body request.UserCreateRequest true "required"
// @accept json
// @Produce json
// @Router /register [post]
// @Success 201 {object} response.JSONSuccessResult{}
// @Failure 400 {object} response.JSONBadRequestResult{}
func (a *authController) Register(c echo.Context) error {
	var req request.UserCreateRequest
	if err := c.Bind(&req); err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}
	if val, err := validation.ValidateUserCreate(req); val == false {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}

	code, err := a.validationUsecase.CreateVerifyUser(req.Email, "verify_user")
	if err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}

	templateData := body_email.VerificationEmailBody{
		URL:      "http://103.176.79.65:8080:8080/verify/" + code + "/" + url.QueryEscape(req.Email),
		Username: req.Username,
		Subject:  "Verification Email",
	}
	err = send_mail.SendEmailVerification(req.Email, templateData)
	if err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}

	createdUser, err := a.authUsecase.Register(&req)
	if err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, true, "success create new user", map[string]interface{}{
		"id":       createdUser.ID,
		"fullname": createdUser.Fullname,
		"username": createdUser.Username,
		"email":    createdUser.Email,
	})
}

func (a *authController) VerifyUser(c echo.Context) error {
	code := c.Param("code")
	emailUser := c.Param("email")
	email, err := a.validationUsecase.VerifyUser(code, emailUser)
	if err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}
	err = a.authUsecase.UpdateVerifiedUser(email)
	if err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, true, "success verified user", map[string]interface{}{
		"email":    email,
		"verified": true,
	})
}

// SendVerifyUser godoc
// @Summary Send verify email
// @Description Send verification to email user
// @Tags Auth
// @param data body request.SendVerifyRequest true "required"
// @accept json
// @Produce json
// @Router /verify/send [post]
// @Success 200 {object} response.JSONSuccessResult{}
// @Failure 400 {object} response.JSONBadRequestResult{}
// @Failure 404 {object} response.JSONBadRequestResult{}
func (a *authController) SendVerifyUser(c echo.Context) error {
	var req request.SendVerifyRequest
	if err := c.Bind(&req); err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}

	user, _ := a.userUsecase.GetDetailUserByEmail(req.Email)
	if user.ID == uuid.FromStringOrNil("") {
		return response.FailResponse(c, http.StatusNotFound, false, "email not found")
	}

	if user.Verified {
		return response.FailResponse(c, http.StatusBadRequest, false, "failed, user is verified")
	}

	code, err := a.validationUsecase.CreateVerifyUser(user.Email, "verify_user")
	if err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}

	templateData := body_email.VerificationEmailBody{
		URL:      "http://103.176.79.65:8080/verify/" + code + "/" + url.QueryEscape(user.Email),
		Username: user.Username,
		Subject:  "Verification Email",
	}

	err = send_mail.SendEmailVerification(user.Email, templateData)
	if err != nil {
		return response.FailResponse(c, http.StatusBadRequest, false, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, true, "success send verified to email user", map[string]interface{}{
		"email": user.Email,
	})
}
