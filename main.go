package main

import (
	"GoMe/db"
	"GoMe/model"
	"io"
	"net/http"
	"os"
	"strings"

	_ "GoMe/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server CRUD App.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host gome.swagger.io
// @BasePath /api
func main() {
	e := echo.New()

	db.GetData()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", helloWorld)
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.GET("/users/all", displayUsers)
	e.GET("/show", show)
	e.Logger.Fatal(e.Start(":1323"))
}

var a []model.User

// e.get("/users/all", displayUsers)
func displayUsers(c echo.Context) error {
	var result strings.Builder

	for _, part := range a {
		result.WriteString(part.Name + " " + part.Email + "\n")
	}
	return c.String(http.StatusOK, result.String())
}

// e.get("/", helloWorld)
func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// e.POST("/users", createUser)
func createUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	a = append(a, *u)

	return c.JSON(http.StatusCreated, u)
	// or
	// return c.XML(http.StatusCreated, u)
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	for x := 0; x < len(a); x++ {
		if id == a[x].Name {
			return c.String(http.StatusOK, id)
		}
	}
	return c.String(http.StatusNotFound, "User not found")
}

//e.GET("/show", show)
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// e.POST("/users", save)
func saveUser(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

func saveAvatar(c echo.Context) error {
	// Get name
	name := c.FormValue("name")
	// Get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}