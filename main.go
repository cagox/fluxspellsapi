package main

import (
	"fmt"
	"github.com/cagox/fluxspellsapi/api"
	"github.com/cagox/fluxspellsapi/app"
	"github.com/cagox/fluxspellsapi/app/database"
	"log"
	"net/http"
)

func main() {
	app.StartLogging()
	defer app.StopLogging()
	database.OpenDatabase()
	defer database.CloseDatabase()
	api.SetupRoutes()

	fmt.Println("Hello World!")
	fmt.Println("Welcome to " + app.Config.SiteName)

	log.Fatal(http.ListenAndServe("localhost:8080", app.Config.Router))

}
