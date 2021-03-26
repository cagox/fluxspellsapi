package routes

import (
	"github.com/cagox/fluxspellsapi/app"
	"github.com/cagox/fluxspellsapi/app/models"
	"net/http"
)

func schoolsHandler(w http.ResponseWriter, r *http.Request) {
	schools := models.SchoolsToJSON()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(schools)

}

func init() {
	app.Config.Router.HandleFunc("/api/schools/", schoolsHandler)
}
