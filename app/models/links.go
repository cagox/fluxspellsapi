package models

import (
	"encoding/gob"
	"github.com/cagox/fluxspellsapi/app"
)

type TypeLink struct {
	ID      int `json:"id"`
	TypeID  int `json:"type_id"`
	SpellID int `json:"spell_id"`
}

type SchoolLink struct {
	ID       int `json:"id"`
	SchoolID int `json:"school_id"`
	SpellID  int `json:"spell_id"`
}

func InsertTypeLink(spellID int, typeID int) *TypeLink {
	newLink := &TypeLink{SpellID: spellID, TypeID: typeID}

	sqlString := `INSERT INTO type_links(type_id, spell_id) VALUES(?,?)`
	res, err := app.Config.Database.Exec(sqlString, typeID, spellID)
	if err != nil {
		panic(err) //TODO: Proper error checking.
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err) //TODO: Proper error checking.
	}

	newLink.ID = int(id)
	return newLink
}

func InsertSchoolLink(spellID int, schoolID int) *SchoolLink {
	newLink := &SchoolLink{SpellID: spellID, SchoolID: schoolID}

	sqlString := `INSERT INTO school_links(school_id, spell_id) VALUES(?,?)`
	res, err := app.Config.Database.Exec(sqlString, schoolID, spellID)
	if err != nil {
		panic(err) //TODO: Proper error checking.
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err) //TODO: Proper error checking.
	}

	newLink.ID = int(id)
	return newLink
}

//TODO: Add the ability to drop links.

func init() {
	gob.Register(TypeLink{})
	gob.Register(SchoolLink{})
}
