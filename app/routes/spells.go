package routes

import (
	"github.com/cagox/fluxspellsapi/app"
	"github.com/cagox/fluxspellsapi/app/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func spellsHandler(w http.ResponseWriter, r *http.Request) {
	spells := models.SpellsToJSON()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(spells)
}

func spellHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id_val, err := strconv.Atoi(vars["spell_id"])
	if err != nil {
		panic(err)
	}

	spell := models.GetSpellByID(id_val)

	spellJson := spell.ToJSON()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(spellJson)
}

func init() {
	app.Config.Router.HandleFunc("/api/spells/", spellsHandler)
	app.Config.Router.HandleFunc("/api/spells", spellsHandler)
	app.Config.Router.HandleFunc("/api/spells/{spell_id:[0-9]+}", spellHandler)
	app.Config.Router.HandleFunc("/api/spells/{spell_id:[0-9]+}/", spellHandler)
}
