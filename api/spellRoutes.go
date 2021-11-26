package api

import (
	"github.com/cagox/fluxspellsapi/app/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func listSpells(w http.ResponseWriter, r *http.Request) {
	spells := models.GetSpellSummaryListAsJSON()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(spells)
}

func viewSpell(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idVal, err := strconv.Atoi(vars["spell_id"])
	if err != nil {
		panic(err) //TODO: Make this more useful.
	}
	spell := models.GetSpellAsJSON(idVal)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(spell)
}
