package routers

import (
	"github.com/Capstone-Kel-23/BE-Rest-API/app/routers/auth"
	"github.com/Capstone-Kel-23/BE-Rest-API/app/routers/invoice"
	"github.com/Capstone-Kel-23/BE-Rest-API/app/routers/profile"
	"github.com/Capstone-Kel-23/BE-Rest-API/app/routers/user"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRouter(e *echo.Echo, db *gorm.DB) {
	auth.AuthRouter(e, db)
	user.UserRouter(e, db)
	profile.ProfileRouter(e, db)
	invoice.InvoiceRouter(e, db)
}
