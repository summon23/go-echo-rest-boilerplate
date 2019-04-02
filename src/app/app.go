package app

import(
	"net/http"

	"github.com/labstack/echo"
)

func startServer() {
	e := echo.New()
	e.GET("/", func (c echo.Context) error {
		return c.String(http.StatusOK, "hell")
	})
	e.Logger.Fatal(e.Start(":1312"))
}