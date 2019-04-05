package controller

import (
	"database/sql"
	"net/http"
	"log"
	// "encoding/json"

	"github.com/labstack/echo"
)

var conn *sql.DB

type UserRecord struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type User struct {
	db *sql.DB
	echo *echo.Echo
}

func getList(c echo.Context) error {
	// var arrResponse []UserRecord
	var temp UserRecord

	if conn == nil {
		panic("db not connect")
	}
	resDB, err := conn.Query("SELECT id, name, email from users")
	if err != nil {
		panic(err.Error())
	}

	log.Print(resDB)
	defer resDB.Close()

	for resDB.Next() {
		log.Print("next")
		err = resDB.Scan(&temp.Id, &temp.Name, &temp.Email)
		log.Println(temp.Id, temp.Name)
		if err != nil {
			panic(err.Error())
		}
		// arrResponse = append(arrResponse, temp)
	}
	r := UserRecord{
		Id: 1,
		Name: "andry",
		Email: "aaa"}
	// r, err := json.Marshal(temp)
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
	conn = DB
	route.GET("/ee", func (c echo.Context) error {
		return c.String(http.StatusOK, "hell")
	})
	route.GET("/user", getList)
	route.GET("/user/:id", getOne)
	route.POST("/user", func (c echo.Context) error {
		return c.String(http.StatusOK, "POST")
	})
}