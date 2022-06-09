package profile

import (
	"github.com/Capstone-Kel-23/BE-Rest-API/internal/profile/delivery/http"
	"github.com/Capstone-Kel-23/BE-Rest-API/internal/profile/repository"
	"github.com/Capstone-Kel-23/BE-Rest-API/internal/profile/usecase"
	mid "github.com/Capstone-Kel-23/BE-Rest-API/internal/user/delivery/http/middleware"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ProfileRouter(e *echo.Echo, db *gorm.DB) {
	authMiddleware := mid.NewGoMiddleware().AuthMiddleware()
	profileRepository := repository.NewProfileRepoitory(db)
	profileUsecase := usecase.NewProfileUsecase(profileRepository)
	profileController := http.NewProfileController(profileUsecase)

	e.PUT("/api/v1/profile", profileController.UpdateProfileUser, authMiddleware)
}
