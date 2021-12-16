package api

import (
	"encoding/json"
	"fmt"
	"github.com/cagox/fluxspellsapi/app/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func listSpells(w http.ResponseWriter, r *http.Request) {
	spells := models.GetSpellSummaryListAsJSON()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(spells)
}

func viewSpell(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idVal, err := strconv.Atoi(vars["spell_id"])
	if err != nil {
		panic(err) //TODO: Make this more useful.
	}
	spell := models.GetSpellAsJSON(idVal)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(spell)
}

func addSpellHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Called.")
	var bodyPost models.SpellPost
	err := json.NewDecoder(r.Body).Decode(&bodyPost)
	//fmt.Println("err: ", err)
	//fmt.Println("bodyPost: ", bodyPost)
	//TODO: Check for Authentication here, or wrap handler in middleware.

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if err == nil {
		spell, err := models.InsertSpell(&bodyPost.BodySpell)
		if err != nil {
			fmt.Println("err on insert: ", err) //TODO: Make this more useful
		}
		w.Write(spell.ToJSON())
		return
	}

	message := `{"spell_id": 6, "message": "It didn't work."}`
	w.Write([]byte(message))

}
