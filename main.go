package main

import (
	"GoMe/model"
	"GoMe/router"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4/middleware"

	_ "GoMe/docs"

	"github.com/labstack/echo/v4"
)

// @title GoMe
// @version 1.0
// @description This is a sample Backend CRUD App using Golang.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://voidstack.xyz
// @contact.email nik3348@gmail.com

// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT

// @host localhost:1323
// @BasePath /api
func main() {
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	_, err := model.DBConn()
	if err != nil {
		panic(err.Error())
	}

	router.Init(e)
	e.Logger.Fatal(e.Start(":1323"))
}
