package handler

import (
	"GoMe/model"
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
)

var a []model.User

func GetUsers(c echo.Context) error {
	// Execute the query
	rows, err := model.DB.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Returning array of objects
	// Create one object to and rewrite every time
	var a []model.User
	u := new(model.User)

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// Backend logs area
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Print(columns[i], ":", value, "\t")
		}
		fmt.Println()

		// Output area
		// Adding to array
		err = rows.Scan(&u.UserId, &u.Name, &u.Email)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		a = append(a, *u)
	}

	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
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
