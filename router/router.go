package router

import (
	"GoMe/handler"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {

	api := e.Group("/api/users")
	{
		api.GET("", handler.GetUsers)
		api.GET("/:id", handler.GetUserById)
		api.POST("", handler.PostUser)
		api.POST("/avatar", handler.SaveAvatar)
	}
}
