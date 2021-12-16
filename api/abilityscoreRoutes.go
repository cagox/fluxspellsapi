package api

import (
	"github.com/cagox/fluxspellsapi/app/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func listScores(w http.ResponseWriter, r *http.Request) {
	scores := models.AbilityScoreListAsJSON()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(scores)
}

func viewScore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idVal, err := strconv.Atoi(vars["ability_score_id"])
	if err != nil {
		panic(err) //TODO: Make this more useful.
	}
	ability := models.AbilityScoreAsJSON(int64(idVal))
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(ability)
}
