package app

import (
	"fmt"
	"os"

	"github.com/Capstone-Kel-23/BE-Rest-API/app/routers"
	template_interface "github.com/Capstone-Kel-23/BE-Rest-API/templates/interface"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Capstone-Kel-23/BE-Rest-API/app/config"
	docs "github.com/Capstone-Kel-23/BE-Rest-API/docs"
	mid "github.com/Capstone-Kel-23/BE-Rest-API/internal/user/delivery/http/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Capstone Application Documentation
// @description This is a Capstone application
// @version 2.0
// @BasePath /api/v1
// @schemes http https
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization

func RunServer() {
	docs.SwaggerInfo.Host = os.Getenv("APP_HOST")

	db := config.InitDB()
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Static("/", "public")
	mid.NewGoMiddleware().LogMiddleware(e)

	e.Renderer = template_interface.NewRenderer("./templates/*.html", true)

	routers.SetupRouter(e, db)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	address := fmt.Sprintf(":%s", os.Getenv("PORT"))

	if err := e.Start(address); err != nil {
		log.Info("shutting down the server")
	}
}
