package api

import (
	"github.com/cagox/fluxspellsapi/app/models"
	"net/http"
)

func listSchoolsHeader(w http.ResponseWriter, r *http.Request) {
	schools := models.GetSchoolHeaderList()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(schools)
}
