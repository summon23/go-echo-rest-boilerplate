package controller

import (
	"database/sql"
	"net/http"
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

func getData(c echo.Context) error {
	// db := dbConn()
	if conn == nil {
		panic("db not connect")
	}
	_, err := conn.Query("SELECT * from users")
	if err != nil {
		panic(err.Error())
	}

	return c.String(http.StatusOK, "returns sstring")
}

func (r *User) RegisterHandler(DB *sql.DB, route *echo.Echo) {
	r.db = DB
	r.echo = route
	
	route.GET("/", func (c echo.Context) error {
		return c.String(http.StatusOK, "hell")
	})
}