package handler

import (
	"GoMe/model"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
)

var a []model.User

func GetUsers(c echo.Context) error {
	// Create one entity to and rewrite every time
	var a []model.User
	u := new(model.User)

	// Execute the query
	rows, err := model.DB.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Fetch rows
	for rows.Next() {
		// Output area
		// Adding to array
		err = rows.Scan(&u.UserId, &u.Name, &u.Email, &u.CourseId)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		a = append(a, *u)
		fmt.Println("UserId:", u.UserId, " Name:", u.Name, " Email:", u.Email.String, " CourseId:", u.CourseId.String)
	}

	return c.JSON(http.StatusOK, a)
}

// e.get("/users/all", displayUsers)
func DisplayUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, a)
}

// e.POST("/users", createUser)
func CreateUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	a = append(a, *u)

	stmtIns, err := model.DB.Prepare("INSERT INTO users(name) VALUES(?)") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	_, err = stmtIns.Exec(u.Name) // Insert tuples (i, i^2)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	return c.JSON(http.StatusCreated, u)
}

// e.GET("/users/:id", getUser)
func GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	for x := 0; x < len(a); x++ {
		if id == a[x].Name {
			return c.String(http.StatusOK, id)
		}
	}
	return c.String(http.StatusNotFound, "User not found")
}

func SaveAvatar(c echo.Context) error {
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
