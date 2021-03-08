package models

import (
	"encoding/gob"
)

type School struct {
	ID          int    `json:"school_id"`
	Name        string `json:"school_name"`
	Description string `json:"school_description"`
	Summary     string `json:"summary"`
	//Spells     Many-to-Many will be mapped as a function later.
}

func (school School) Spells() {
	//This is a stub. It will return a list of spells in this school.
}

func init() {
	gob.Register(School{})
}
