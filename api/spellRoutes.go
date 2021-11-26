package api

import (
	"github.com/cagox/fluxspellsapi/app/models"
	"net/http"
)

func listSpells(w http.ResponseWriter, r *http.Request) {
	spells := models.GetSpellSummaryListAsJSON()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(spells)
}
