package models

import (
	"encoding/gob"
)

type Type struct {
	ID          int    `json:"type_id"`
	Name        string `json:"type_name"`
	Description string `json:"type_description"`
	Summary     string `json:"summary"`
	//Spells     Many-to-Many will be mapped as a function later.
}

func (thisType Type) Spells() {
	//This is a stub. It will return a list of spells in this school.
}

func init() {
	gob.Register(Type{})
}
