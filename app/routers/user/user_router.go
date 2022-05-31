package user

import (
	"github.com/Capstone-Kel-23/BE-Rest-API/internal/user/delivery/http"
	mid "github.com/Capstone-Kel-23/BE-Rest-API/internal/user/delivery/http/middleware"
	"github.com/Capstone-Kel-23/BE-Rest-API/internal/user/repository"
	"github.com/Capstone-Kel-23/BE-Rest-API/internal/user/usecase"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRouter(e *echo.Echo, db *gorm.DB) {
	authMiddleware := mid.NewGoMiddleware().AuthMiddleware()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := http.NewUserController(userUsecase)

	// User Endpoints
	e.GET("/api/v1/users", userController.GetListAllUsers, authMiddleware)
}
