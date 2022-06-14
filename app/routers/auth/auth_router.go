package auth

import (
	repository4 "github.com/Capstone-Kel-23/BE-Rest-API/internal/profile/repository"
	repository2 "github.com/Capstone-Kel-23/BE-Rest-API/internal/role/repository"
	"github.com/Capstone-Kel-23/BE-Rest-API/internal/user/delivery/http"
	"github.com/Capstone-Kel-23/BE-Rest-API/internal/user/repository"
	"github.com/Capstone-Kel-23/BE-Rest-API/internal/user/usecase"
	repository3 "github.com/Capstone-Kel-23/BE-Rest-API/internal/validation/repository"
	usecase2 "github.com/Capstone-Kel-23/BE-Rest-API/internal/validation/usecase"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AuthRouter(e *echo.Echo, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	roleRepository := repository2.NewRoleRepository(db)
	validationRepository := repository3.NewValidationRepository(db)
	profileRepository := repository4.NewProfileRepoitory(db)

	authUsecase := usecase.NewAuthUsecase(userRepository, roleRepository, profileRepository)
	validationUsecase := usecase2.NewValidationUsecase(validationRepository, userRepository)
	userUsecase := usecase.NewUserUsecase(userRepository)

	authController := http.NewAuthController(authUsecase, validationUsecase, userUsecase)

	// Auth Endpoints (User)
	e.POST("/api/v1/register", authController.Register)
	e.POST("/api/v1/login", authController.Login)

	// Verify Endpoints
	e.GET("/verify/:code/:email", authController.VerifyUser)
	e.POST("/api/v1/verify/send", authController.SendVerifyUser)
}
