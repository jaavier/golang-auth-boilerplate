package main

import (
	"app/handlers"
	"app/middlewares"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(echoMiddleware.Recover())

	e.POST("/signup", handlers.Signup)
	e.POST("/login", handlers.Login)
	e.POST("/logout", handlers.Logout, middlewares.JWT)

	e.GET("/private", handlers.Private, middlewares.JWT)

	e.Logger.Fatal(e.Start(":8080"))
}
