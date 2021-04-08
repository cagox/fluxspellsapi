package routes

import (
	"github.com/cagox/fluxspellsapi/app"
	"github.com/cagox/fluxspellsapi/app/models"
	"net/http"
)

func spellsHandler(w http.ResponseWriter, r *http.Request) {
	spells := models.SpellsToJSON()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(spells)
}


func init() {
	app.Config.Router.HandleFunc("/api/spells/", spellsHandler)
	app.Config.Router.HandleFunc("/api/spells", spellsHandler)
}