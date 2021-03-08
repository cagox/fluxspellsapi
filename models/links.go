package models

import (
	"encoding/gob"
)

type TypeLink struct {
	TypeID  int `json:"type_id"`
	SpellID int `json:"spell_id"`
}

type SchoolLink struct {
	SchoolID int `json:"school_id"`
	SpellID  int `json:"spell_id"`
}

func init() {
	gob.Register(TypeLink{})
	gob.Register(SchoolLink{})
}
