package routes

import (
	"github.com/cagox/fluxspellsapi/app"
	"github.com/cagox/fluxspellsapi/app/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func schoolsHandler(w http.ResponseWriter, r *http.Request) {
	schools := models.SchoolsToJSON()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(schools)

}

func schoolHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_val, err := strconv.Atoi(vars["school_id"])
	if err != nil {
		panic(err)
	}

	school := models.GetSchoolByID(id_val)
	
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(school.ToJSON())
}

func schoolSpellsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_val, err := strconv.Atoi(vars["school_id"])
	if err != nil {
		panic(err)
	}

	school := models.GetSchoolByID(id_val)
	
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(school.SpellsToJSON())
}



func init() {
	app.Config.Router.HandleFunc("/api/schools/", schoolsHandler)
	app.Config.Router.HandleFunc("/api/schools", schoolsHandler)
	app.Config.Router.HandleFunc("/api/schools/{school_id:[0-9]+}", schoolHandler)
	app.Config.Router.HandleFunc("/api/schools/{school_id:[0-9]+}/spells", schoolSpellsHandler)
	app.Config.Router.HandleFunc("/api/schools/{school_id:[0-9]+}/spells/", schoolSpellsHandler)
}


