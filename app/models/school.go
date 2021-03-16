package models

import (
	"encoding/gob"
)

type School struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Summary     string `json:"summary"`
	//Spells     Many-to-Many will be mapped as a function later.
}

func (school School) Spells() {
	//This is a stub. It will return a list of spells in this school.
}

func init() {
	gob.Register(School{})
}
