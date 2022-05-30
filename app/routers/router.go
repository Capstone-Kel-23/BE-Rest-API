package routers

import (
	"github.com/Capstone-Kel-23/BE-Rest-API/app/routers/auth"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRouter(e *echo.Echo, db *gorm.DB) {
	auth.AuthRouter(e, db)
}
