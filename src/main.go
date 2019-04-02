package main

import(
	"github.com/summon23/go-echo-rest-boilerplate/src/app"
)

func main() {
	mainApp := app.App{}
	app.startServer()
}