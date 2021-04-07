package routes

import (
	"github.com/cagox/fluxspellsapi/app"
	"github.com/cagox/fluxspellsapi/app/models"
	"github.com/gorilla/mux"
	"net/http"
)

func schoolsHandler(w http.ResponseWriter, r *http.Request) {
	schools := models.SchoolsToJSON()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(schools)

}

func schoolHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	school := School.GetSchoolByID(vars['school_id'])
	
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(school.ToJSON())
}

func schoolSpellsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	school := School.GetSchoolByID(vars['school_id'])
	
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(school.SpellsToJSON())
}



func init() {
	app.Config.Router.HandleFunc("/api/schools/", schoolsHandler)
	app.Conig.Router.HandleFunc("/api/schools/{school_id}", schoolHandler)
	app.Conig.Router.HandleFunc("/api/schools/{school_id}/spells", schoolSpellHandler)
}


