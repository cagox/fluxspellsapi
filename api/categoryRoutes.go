package api

import (
	"github.com/cagox/fluxspellsapi/app/models"
	"net/http"
)

func listCategoriesHeader(w http.ResponseWriter, r *http.Request) {
	categories := models.GetCategoryHeaderList()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(categories)
}
