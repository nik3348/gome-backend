package main

import (
	"GoMe/model"
	"GoMe/router"
	"context"
	"os"
	"os/signal"
	"time"

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
	e.Use(middleware.Secure())
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	_, err := model.DBConn()
	model.InitData()
	if err != nil {
		panic(err.Error())
	}

	router.Init(e)

	// Start server
	go func() {
		if err := e.Start(":1323"); err != nil {
			e.Logger.Fatal(err)
			e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

// TODO: Basic Auth/JWT
// TODO: Http/2
