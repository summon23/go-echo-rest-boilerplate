package controller

import (
	"database/sql"
	"net/http"
	"encoding/json"

	"github.com/labstack/echo"
)

var conn *sql.DB

type UserRecord struct {
	id int
	name string
	email string
}

type User struct {
	db *sql.DB
	echo *echo.Echo
}

func getList(c echo.Context) error {
	if conn == nil {
		panic("db not connect")
	}
	resDB, err := conn.Query("SELECT id, name, email from users")
	if err != nil {
		panic(err.Error())
	}

	userList := UserRecord{}
	for resDB.Next() {
		var id int
		var name, email string
		err = resDB.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}
		userList.id = id
		userList.name = name
		userList.email = email
	}

	r, err := json.Marshal(userList)
	return c.JSON(http.StatusOK, r)
}

func getOne(c echo.Context) error {
	if conn == nil {
		panic("db not connect")
	}
	_, err := conn.Query("SELECT * from users")
	if err != nil {
		panic(err.Error())
	}

	return c.String(http.StatusOK, "returns sstring")
}

// RegisterHandler map controller router
func (r *User) RegisterHandler(DB *sql.DB, route *echo.Echo) {
	r.db = DB
	r.echo = route
	
	route.GET("/ee", func (c echo.Context) error {
		return c.String(http.StatusOK, "hell")
	})
	route.GET("/user", getList)
	route.GET("/user/:id", getOne)
	route.POST("/user", func (c echo.Context) error {
		return c.String(http.StatusOK, "POST")
	})
}