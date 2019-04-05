package app

import(
	"database/sql"
	"net/http"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"	
	"github.com/summon23/go-echo-rest-boilerplate/src/app/controller"
)

func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:@/testo")
	if err != nil {
		panic(err.Error())
	}
	return db
}

// StartServer run the server
func StartServer() {
	fmt.Println("test")
	e := echo.New()
	
	e.GET("/eeff", func (c echo.Context) error {
		return c.String(http.StatusOK, "hell")
	})

	db := dbConn()
	

	userHandler := controller.User{}
	userHandler.RegisterHandler(db, e)

	e.Logger.Fatal(e.Start(":1322"))
}
