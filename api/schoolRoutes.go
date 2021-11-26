package api

import (
	"github.com/cagox/fluxspellsapi/app/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func listSchoolsHeader(w http.ResponseWriter, r *http.Request) {
	schools := models.GetSchoolHeaderList()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(schools)
}

func listSchools(w http.ResponseWriter, r *http.Request) {
	schools := models.GetSchoolListAsJSON()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(schools)
}

func viewSchool(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idVal, err := strconv.Atoi(vars["school_id"])
	if err != nil {
		panic(err) //TODO: Make this more useful.
	}
	school := models.GetSchoolAsJSON(idVal)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(school)
}

func schoolSpells(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idVal, err := strconv.Atoi(vars["school_id"])
	if err != nil {
		panic(err) //TODO: Make this more useful.
	}
	spells := models.GetSpellsBySchoolAsJSON(idVal)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(spells)

}
