package main

import (
	"fmt"
	"github.com/cagox/fluxspellsapi/app"
	"github.com/cagox/fluxspellsapi/app/database"
	_ "github.com/cagox/fluxspellsapi/app/routes"
	"log"
	"net/http"
)

func main() {
	app.StartLogging()
	defer app.StopLogging()
	database.OpenDatabase()
	defer database.CloseDatabase()

	fmt.Println("Hello World!")
	fmt.Println("Welcome to " + app.Config.SiteName)

	log.Fatal(http.ListenAndServe("localhost:8080", app.Config.Router))

}
