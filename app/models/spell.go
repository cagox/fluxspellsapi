package models

import (
	"encoding/gob"
)

type Spell struct {
	ID            int    `json:"spell_id"`
	Name          string `json:"spellname"`
	Description   string `json:"spelldescription"`
	Summary       string `json:"summary"`
	Cost          string `json:"cost"`
	Difficulty    string `json:"difficulty"`
	Prerequisites string `json:"prerequisites"`
	Components    string `json:"components"`
	Range         string `json:"spellrange"`
	//Types        Many-to-Many will be mapped as a function later.
	//Schools      Many-to-Many will be mapped as a function later.
}

func Types() {
	//this is a stub. It will eventually return a list of types/categories associated with the spell.
}

func Schools() {
	//This is a stub. It will eventually return a list of schools that the spell belongs in.
}

func init() {
	gob.Register(Spell{})
}
