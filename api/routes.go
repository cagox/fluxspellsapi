package api

import (
	//"net/http"
	"github.com/cagox/fluxspellsapi/app"
)

func SetupRoutes() {
	app.Config.Router.HandleFunc("/", indexHandler) //This one won't technically be needed, but should be included to avoid anomalies.

	//Authentication
	app.Config.Router.HandleFunc("/login", userLogin).Methods("POST")
	//app.Config.Router.HandleFunc("/logout", userLogout).Methods.("GET")

	//Create things (The C in Crud)
	//app.Config.Router.HandleFunc("/school", addSchoolHandler).Methods("POST")
	//app.Config.Router.HandleFunc("/category", addCategoryHandler).Methods("POST")
	app.Config.Router.HandleFunc("/spell", IfAuthenticated(addSpellHandler)).Methods("POST", "OPTIONS")

	//View things (the R in CRUD)
	app.Config.Router.HandleFunc("/schools", listSchools).Methods("GET")
	app.Config.Router.HandleFunc("/schools/header", listSchoolsHeader).Methods("GET")
	app.Config.Router.HandleFunc("/categories", listCategories).Methods("GET")
	app.Config.Router.HandleFunc("/categories/header", listCategoriesHeader).Methods("GET")
	app.Config.Router.HandleFunc("/spells", listSpells).Methods("GET")
	app.Config.Router.HandleFunc("/schools/{school_id}", viewSchool).Methods("GET")
	app.Config.Router.HandleFunc("/schools/{school_id}/spells", schoolSpells).Methods("GET")
	app.Config.Router.HandleFunc("/categories/{category_id}", viewCategory).Methods("GET")
	app.Config.Router.HandleFunc("/categories/{category_id}/spells", categorySpells).Methods("GET")
	app.Config.Router.HandleFunc("/spells/{spell_id}", viewSpell).Methods("GET")
	app.Config.Router.HandleFunc("/abilityscores", listScores).Methods("GET")
	app.Config.Router.HandleFunc("/abilityscores/{ability_score_id}", viewScore).Methods("GET")

	//Edit things (The U in CRUD)
	//app.Config.Router.HandleFunc("/spells/{id}/edit", editSpell).Methods("GET")
	//app.Config.Router.HandleFunc("/spells/{id}, updateSpell).Methods("PATCH")
	//app.Config.Router.HandleFunc("/schools/{id}/edit", editSchool).Methods("GET")
	//app.Config.Router.HandleFunc("/schools/{id}, updateSchool).Methods("PATCH")
	//app.Config.Router.HandleFunc("/categories/{id}/edit", editCategory).Methods("GET")
	//app.Config.Router.HandleFunc("/categories/{id}, updateCategory).Methods("PATCH")

	//Delete things (The D in CRUD)
	//app.Config.Router.HandleFunc("/spells/{id}/delete", confirmSpellDelete).Methods("GET")
	//app.Config.Router.HandleFunc("/spells/{id}, deleteSpell).Methods("DELETE")
	//app.Config.Router.HandleFunc("/schools/{id}/delete", confirmSchoolDelete).Methods("GET")
	//app.Config.Router.HandleFunc("/schools/{id}, deleteSchool).Methods("DELETE")
	//app.Config.Router.HandleFunc("/categories/{id}/delete", confirmCategoryDelete).Methods("GET")
	//app.Config.Router.HandleFunc("/categories/{id}, deleteCategory).Methods("DELETE")
}
