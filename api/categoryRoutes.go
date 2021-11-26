package api

import (
	"github.com/cagox/fluxspellsapi/app/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func listCategoriesHeader(w http.ResponseWriter, r *http.Request) {
	categories := models.GetCategoryHeaderList()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(categories)
}

func listCategories(w http.ResponseWriter, r *http.Request) {
	categories := models.GetCategoryListAsJSON()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(categories)
}

func viewCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idVal, err := strconv.Atoi(vars["category_id"])
	if err != nil {
		panic(err) //TODO: Make this more useful.
	}
	category := models.GetCategoryAsJSON(idVal)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(category)
}

func categorySpells(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idVal, err := strconv.Atoi(vars["category_id"])
	if err != nil {
		panic(err) //TODO: Make this more useful.
	}
	spells := models.GetSpellsByCategoryAsJSON(idVal)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(spells)

}
