package app

import(
	"database/sql"
	"net/http"

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
	e := echo.New()

	registerMethod(User)
	db := dbConn()

	userHandler := controller.User{}
	userHandler.RegisterHandler(db, e)
	// e.GET("/", func (c echo.Context) error {
	// 	return c.String(http.StatusOK, "hell")
	// })

	// e.GET("/test", getData)
	e.GET("/user", GetAllUser)
	e.Logger.Fatal(e.Start(":1322"))
}
