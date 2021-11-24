package api

import (
	//"net/http"
	"github.com/cagox/fluxspellsapi/app"
)

func SetupRoutes() {
	app.Config.Router.HandleFunc("/", indexHandler) //This one won't technically be needed, but should be included to avoid anomalies.

	//Authentication
	//app.Config.Router.HandleFunc("/api/login", userLogin).Methods("POST")
	//app.Config.Router.HandleFunc("/api/logout", userLogout).Methods.("GET")

	//Create things (The C in Crud)
	//app.Config.Router.HandleFunc("/api/school", createSchool).Methods("POST")
	//app.Config.Router.HandleFunc("/api/category", createCategory).Methods("POST")
	//app.Config.Router.HandleFunc("/api/spell", createSpell).Methods("POST")

	//View things (the R in CRUD)
	//app.Config.Router.HandleFunc("/api/schools", listSchools).Methods("GET")
	app.Config.Router.HandleFunc("/api/schools/header", listSchoolsHeader).Methods("GET")
	//app.Config.Router.HandleFunc("/api/categories", listCategories).Methods("GET")
	//app.Config.Router.Handlefunct("/api/spells", listSpells).Methods("GET)
	//app.Config.Router.HandleFunc("/api/schools/{id}", viewSchool).Methods("GET")
	//app.Config.Router.HandleFunc("/api/categories/{id}", viewCategory).Methods("GET")
	//app.Config.Router.HandleFunc("/api/spells/{id}", viewSpell).Methods("GET")

	//Edit things (The U in CRUD)
	//app.Config.Router.HandleFunc("/api/spells/{id}/edit", editSpell).Methods("GET")
	//app.Config.Router.HandleFunc("/api/spells/{id}, updateSpell).Methods("PATCH")
	//app.Config.Router.HandleFunc("/api/schools/{id}/edit", editSchool).Methods("GET")
	//app.Config.Router.HandleFunc("/api/schools/{id}, updateSchool).Methods("PATCH")
	//app.Config.Router.HandleFunc("/api/categories/{id}/edit", editCategory).Methods("GET")
	//app.Config.Router.HandleFunc("/api/categories/{id}, updateCategory).Methods("PATCH")

	//Delete things (The D in CRUD)
	//app.Config.Router.HandleFunc("/api/spells/{id}/delete", confirmSpellDelete).Methods("GET")
	//app.Config.Router.HandleFunc("/api/spells/{id}, deleteSpell).Methods("DELETE")
	//app.Config.Router.HandleFunc("/api/schools/{id}/delete", confirmSchoolDelete).Methods("GET")
	//app.Config.Router.HandleFunc("/api/schools/{id}, deleteSchool).Methods("DELETE")
	//app.Config.Router.HandleFunc("/api/categories/{id}/delete", confirmCategoryDelete).Methods("GET")
	//app.Config.Router.HandleFunc("/api/categories/{id}, deleteCategory).Methods("DELETE")
}
