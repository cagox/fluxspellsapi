package api

import (
	//"net/http"
	"github.com/cagox/fluxspellsapi/app"
)

func SetupRoutes() {
	app.Config.Router.HandleFunc("/", indexHandler) //This one won't technically be needed, but should be included to avoid anomalies.
	/*
		//Authentication
		app.Config.Router.HandleFunc("/login", userLogin).Methods("POST")
		app.Config.Router.HandleFunc("/logout", userLogout).Methods.("GET")

		//Create things (The C in Crud)
		app.Config.Router.HandleFunc("/school", createSchool).Methods("POST")
		app.Config.Router.HandleFunc("/category", createCategory).Methods("POST")
		app.Config.Router.HandleFunc("/spell", createSpell).Methods("POST")


		//View things (the R in CRUD)
		app.Config.Router.HandleFunc("/schools", listSchools).Methods("GET")
		app.Config.Router.HandleFunc("/categories", listCategories).Methods("GET")
		app.Config.Router.Handlefunct("/spells", listSpells).Methods("GET)
		app.Config.Router.HandleFunc("/schools/{id}", viewSchool).Methods("GET")
		app.Config.Router.HandleFunc("/categories/{id}", viewCategory).Methods("GET")
		app.Config.Router.HandleFunc("/spells/{id}", viewSpell).Methods("GET")

		//Edit things (The U in CRUD)
		app.Config.Router.HandleFunc("/spells/{id}/edit", editSpell).Methods("GET")
		app.Config.Router.HandleFunc("/spells/{id}, updateSpell).Methods("PATCH")
		app.Config.Router.HandleFunc("/schools/{id}/edit", editSchool).Methods("GET")
		app.Config.Router.HandleFunc("/schools/{id}, updateSchool).Methods("PATCH")
		app.Config.Router.HandleFunc("/categories/{id}/edit", editCategory).Methods("GET")
		app.Config.Router.HandleFunc("/categories/{id}, updateCategory).Methods("PATCH")

		//Delete things (The D in CRUD)
		app.Config.Router.HandleFunc("/spells/{id}/delete", confirmSpellDelete).Methods("GET")
		app.Config.Router.HandleFunc("/spells/{id}, deleteSpell).Methods("DELETE")
		app.Config.Router.HandleFunc("/schools/{id}/delete", confirmSchoolDelete).Methods("GET")
		app.Config.Router.HandleFunc("/schools/{id}, deleteSchool).Methods("DELETE")
		app.Config.Router.HandleFunc("/categories/{id}/delete", confirmCategoryDelete).Methods("GET")
		app.Config.Router.HandleFunc("/categories/{id}, deleteCategory).Methods("DELETE")
	*/
}
