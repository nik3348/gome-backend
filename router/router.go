package router

import (
	"GoMe/handler"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {

	api := e.Group("/api/users")
	{
		api.POST("/", handler.CreateUser)
		api.POST("/avatar", handler.SaveAvatar)
		api.GET("/:id", handler.GetUser)
		api.GET("/all", handler.DisplayUsers)
	}

}
