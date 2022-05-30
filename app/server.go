package app

import (
	"fmt"
	"github.com/Capstone-Kel-23/BE-Rest-API/app/routers"
	"github.com/labstack/echo/v4/middleware"
	"os"

	"github.com/Capstone-Kel-23/BE-Rest-API/app/config"
	docs "github.com/Capstone-Kel-23/BE-Rest-API/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Capstone application Documentation
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

	routers.SetupRouter(e, db)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	address := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	if err := e.Start(address); err != nil {
		log.Info("shutting down the server")
	}
}
